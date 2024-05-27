package comment_job

import (
	"comment/comment_service/service"
	"comment/global"
	"comment/models"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func CreateCommentConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "CreateComment",
		Topic:       "create_comment",
		Partition:   0,                 // 固定分区
		MinBytes:    10e3,              // 10KB
		MaxBytes:    10e6,              // 10MB
		StartOffset: kafka.FirstOffset, // 从最早的消息开始读取
	})
	defer reader.Close()

	ctx := context.Background()

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Printf("failed to fetch message: %v", err)
			break
		}

		var cmr service.CreateMessageRequest
		err = proto.Unmarshal(msg.Value, &cmr)
		if err != nil {
			global.Log.Warn("unmarshal error")
			return
		}

		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("failed to commit message: %v", err)
		}

		//TODO 写入mysql，并刷新到redis
		commentIndexModel := models.CommentIndexModels{
			Root:      cmr.Root,
			Parent:    cmr.Parent,
			MemberID:  cmr.MemberId,
			ObjID:     cmr.ObjId,
			Count:     0,
			RootCount: 0,
			Like:      0,
			Hate:      0,
			State:     int8(cmr.State),
			ObjType:   int8(cmr.ObjType),
		}

		var cms models.CommentSubjectModels
		//开启事务修改表
		err = global.DB.Transaction(func(tx *gorm.DB) error {
			err = global.DB.Create(&commentIndexModel).Error
			if err != nil {
				global.Log.Warn("评论索引表注册失败")
				return err
			}

			commentCommentModel := models.CommentContentModels{
				CommentID:   commentIndexModel.ID,
				IP:          cmr.Ip,
				AtMemberIds: strconv.FormatInt(cmr.MemberId, 10),
				Message:     cmr.Comment,
			}
			err = global.DB.Create(&commentCommentModel).Error
			if err != nil {
				global.Log.Warn("评论内容表注册失败")
				return err
			}

			//TODO 更新subject表
			var updates map[string]interface{}
			if commentIndexModel.Root == 0 { //该comment为根评论
				updates = map[string]interface{}{
					"count":      gorm.Expr("count + ?", 1),
					"all_count":  gorm.Expr("all_count + ?", 1),
					"root_count": gorm.Expr("root_count + ?", 1),
				}
			} else {
				updates = map[string]interface{}{
					"count":     gorm.Expr("count + ?", 1),
					"all_count": gorm.Expr("all_count + ?", 1),
				}
			}
			global.DB.Model(&models.CommentSubjectModels{}).Where("obj_type = ? AND obj_id = ?", commentIndexModel.ObjType,
				commentIndexModel.ObjID).Updates(updates).Scan(&cms)

			return nil
		})
		if err != nil {
			// TODO 错误处理
		}
		global.Log.Info("message apply in mysql")

		// 更新缓存 comment_subject_cache
		cs := service.CommentSubject{
			Id:        cms.ID,
			ObjType:   int32(cms.ObjType),
			ObjId:     cms.ObjID,
			MemberId:  cms.MemberID,
			CreatedAt: timestamppb.New(cms.CreatedAt),
			UpdatedAt: timestamppb.New(cms.UpdatedAt),
			Count:     cms.Count,
			RootCount: cms.RootCount,
			AllCount:  cms.AllCount,
			State:     int32(cms.State),
		}
		csMarshal, err := proto.Marshal(&cs)
		if err != nil {
			return
		}

		ctx := context.Background()
		oidType := strconv.FormatInt(cms.ID, 10) + "_" + strconv.FormatInt(int64(cms.ObjType), 10)
		// msg.Value是序列化后的数据
		global.Redis.Set(ctx, oidType, csMarshal, 0)

		// 增量缓存 comment_index_cache
		oidTypeSort := strconv.FormatInt(cms.ID, 10) + "_" + strconv.FormatInt(int64(cms.ObjType), 10) + "sortByDESC"
		score := float64(commentIndexModel.Like + commentIndexModel.RootCount)
		global.Redis.ZAdd(ctx, oidTypeSort, []redis.Z{{Score: score, Member: commentIndexModel.ID}}...)

		// comment_content_cache
		content := service.Content{Content: cmr.Comment}
		contentMarshal, err := proto.Marshal(&content)
		if err != nil {
			return
		}

		global.Redis.Set(ctx, strconv.FormatInt(commentIndexModel.ID, 10), contentMarshal, 0)
		global.Log.Info("message apply in redis")
	}
}

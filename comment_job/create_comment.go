package comment_job

import (
	"comment/comment_service/service"
	"comment/global"
	"comment/models"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
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
			Floor:     cmr.Floor,
		}

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

			return nil
		})
		if err != nil {
			// TODO 错误处理
		}

		fmt.Println("message apply in mysql")
	}
}

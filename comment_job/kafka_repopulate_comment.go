package comment_job

import (
	"comment/comment_service/service"
	"comment/global"
	"comment/models"
	"context"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"log"
)

func RepopulateComment() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "RepopulateComment",
		Topic:       "repopulate_comment",
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

		var gcr service.GetCommentRequest
		err = proto.Unmarshal(msg.Value, &gcr)
		if err != nil {
			global.Log.Warn("unmarshal error")
			return
		}

		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("failed to commit message: %v", err)
		}

		// 回填一个对象的多个评论
		var csm models.CommentSubjectModels
		var cim []models.CommentIndexModels
		var ccm []models.CommentContentModels
		//到mysql内查询数据
		global.DB.Where("obj_id = ? AND obj_type = ?", gcr.ObjID, gcr.ObjType).Offset(int(gcr.Offset)).Take(&csm)

		global.DB.Select("id, member_id").Where("obj_id = ? and obj_type = ?", gcr.ObjID, gcr.ObjType).Offset(int(gcr.Offset)).Limit(10).Find(&cim)
		var idList []int64
		for i := 0; i < len(cim); i++ {
			idList = append(idList, cim[i].ID)
		}

		global.DB.Select("message").Where("comment_id in ?", idList).Find(&ccm)

		global.Log.Info("message get in mysql")

		// 回填redis
		// 更新缓存 comment_subject_cache
		ctx := context.Background()
		comment := make([]string, len(ccm))
		for i := 0; i < len(ccm); i++ {
			comment[i] = ccm[i].Message
		}

		addSubjectAndCommentCache(ctx, &csm, &cim,
			comment)
		//go addSubjectCache(ctx, &csm)
		//
		//// 增量缓存 comment_index_cache
		//go addCommentIndexCache(ctx, csm.ObjID, int64(csm.ObjType), &cim)
		//
		//// comment_content_cache
		//for i := 0; i < len(cim); i++ {
		//	go addCommentCommentCache(ctx, ccm[i].Message, cim[i].ID)
		//}

		global.Log.Info("message get in redis")
	}
}

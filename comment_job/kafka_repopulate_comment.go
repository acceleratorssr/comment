package comment_job

import (
	"comment/comment_service/service"
	"comment/global"
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

		var cmr service.GetCommentRequest
		err = proto.Unmarshal(msg.Value, &cmr)
		if err != nil {
			global.Log.Warn("unmarshal error")
			return
		}

		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("failed to commit message: %v", err)
		}

		// TODO 到mysql内查询数据
		global.Log.Info("message get in mysql")

		// TODO 回填redis
		// 更新缓存 comment_subject_cache

		// 增量缓存 comment_index_cache

		// comment_content_cache

		global.Log.Info("message get in redis")
	}
}

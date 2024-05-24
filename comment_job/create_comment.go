package comment_job

import (
	"comment/comment_service/service"
	"comment/global"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"log"
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

		//TODO 写入mysql，并刷新到redis

		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("failed to commit message: %v", err)
		}
		fmt.Println("message committed")
	}
}

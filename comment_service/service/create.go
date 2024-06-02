package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func producer(request []byte) {
	topic := "create_comment"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: request},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	fmt.Println("producer success")
}

//
//func initKafka() (writer *kafka.Writer) {
//	writer = &kafka.Writer{
//		Addr:         kafka.TCP("localhost:9092"),
//		Topic:        "create_comment",
//		Balancer:     &kafka.LeastBytes{},
//		WriteTimeout: 10 * time.Second,
//	}
//	return
//}
//
//func closeKafka() {
//	if writer != nil {
//		err := writer.Close()
//		if err != nil {
//			log.Fatal("failed to close writer:", err)
//		}
//	}
//}
//
//func NewProducer(request []byte) {
//	err := writer.WriteMessages(context.Background(),
//		kafka.Message{
//			Value: request,
//		},
//	)
//	if err != nil {
//		log.Println("failed to write messages:", err)
//		return
//	}
//	fmt.Println("producer success")
//}

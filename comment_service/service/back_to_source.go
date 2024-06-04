package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func backToSource(request []byte, key string) {
	topic := "repopulate_comment"
	partition := 0

	//_, err, _ := global.SF.Do(key+"kafka", func() (interface{}, error) {
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
	//return nil, nil
	//})
	//if err != nil {
	//	global.Log.Error("backToSource -> ", err)
	//}

}

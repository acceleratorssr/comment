package core

import (
	commentjob "comment/comment_job"
)

func Kafka() {
	go InitKafka()
}

func InitKafka() {
	commentjob.CreateCommentConsumer()
}

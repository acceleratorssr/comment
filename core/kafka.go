package core

import (
	commentjob "comment/comment_job"
)

func Kafka() {
	go InitKafka()
}

func InitKafka() {
	go commentjob.CreateCommentConsumer()
	go commentjob.RepopulateComment()
}

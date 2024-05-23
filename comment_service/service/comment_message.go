package service

import (
	"context"
)

type CommentMessage struct {
}

func NewCommentMessageServer() *CommentMessage {
	return &CommentMessage{}
}

// CreateCommentMessage 将评论先写入kafka，再由kafka事务写入mysql，刷新redis
func (c *CommentMessage) CreateCommentMessage(ctx context.Context, request *CreateMessageRequest) (*CreateMessageResponse, error) {
	CMR := CreateMessageResponse{
		Success: true,
	}
	return &CMR, nil
}

// GetComment 从redis查询，如果没有则从mysql查询，如果需要刷新redis则交给kafka异步处理
func (c *CommentMessage) GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error) {
	GCR := GetCommentResponse{
		ObjID: 1,
	}
	return &GCR, nil
}

func (c *CommentMessage) mustEmbedUnimplementedMessageServiceServer() {}

package service

import (
	"comment/global"
	"context"
	"google.golang.org/protobuf/proto"
	"sync"
)

type CommentMessage struct {
	wg sync.WaitGroup
}

func NewCommentMessageServer() *CommentMessage {
	return &CommentMessage{}
}

// CreateCommentMessage 将评论先写入kafka，再由kafka事务写入mysql，刷新redis
// 专注于数据处理，准备数据
func (c *CommentMessage) CreateCommentMessage(ctx context.Context, request *CreateMessageRequest) (*CreateMessageResponse, error) {
	data, err := proto.Marshal(&CreateMessageRequest{
		ObjId:    request.ObjId,
		MemberId: request.MemberId,
		State:    request.State,
		ObjType:  request.ObjType,
		Root:     request.Root,
		Parent:   request.Parent,
		Floor:    request.Floor,
		Ip:       request.Ip,
		Comment:  request.Comment,
	})
	if err != nil {
		global.Log.Warn("proto.Marshal err: %v", err)
		return nil, err
	}

	go producer(data)

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

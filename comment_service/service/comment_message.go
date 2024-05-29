package service

import (
	"comment/global"
	"context"
	"google.golang.org/protobuf/proto"
	"strconv"
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
		ObjType:  request.ObjType,
		Root:     request.Root,
		Parent:   request.Parent,
		Floor:    request.Floor,
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

// GetComment 先去redis内拿数据，如果没有则透传到mysql，并且使用kafka回源
func (c *CommentMessage) GetComment(ctx context.Context, request *GetCommentRequest) (*GetCommentResponse, error) {
	// TODO 先去redis内拿数据，如果没有则透传到mysql，并且使用kafka回源

	var gcr GetCommentResponse
	// 先查对象缓存
	oidType := strconv.FormatInt(request.ObjID, 10) + "_" + strconv.FormatInt(int64(request.ObjType), 10)
	getSub := global.Redis.Get(ctx, oidType)
	if getSub.Err() != nil {
		global.Log.Info("对象表缓存失效")
	} else {
		cs := CommentSubject{}
		err := proto.Unmarshal([]byte(getSub.Val()), &cs)
		if err != nil {
			return nil, err
		}
		AddSubToResponse(&gcr, cs.Count, cs.RootCount, cs.AllCount)
	}

	// 再查评论索引缓存
	oidTypeSort := strconv.FormatInt(request.ObjID, 10) + "_" + strconv.FormatInt(int64(request.ObjType), 10) + "sortByDESC"
	getIndex := global.Redis.ZRange(ctx, oidTypeSort, 0, -1)
	if getIndex.Err() != nil {
		global.Log.Info("评论索引表缓存失效")
	} else {
		length := int32(len(getIndex.Val()))
		// TODO 从Offset开始，一次返回10条数据
		for i := request.Offset; i < length; i++ {
			s, _ := strconv.Atoi(getIndex.Val()[i])
			AddIndexToResponse(&gcr, int64(s))
		}
	}

	// 最后查评论内容缓存
	content := Content{}
	for i := 0; i < len(gcr.Id); i++ {
		getComment := global.Redis.Get(ctx, strconv.FormatInt(gcr.Id[i], 10))
		err := proto.Unmarshal([]byte(getComment.Val()), &content)
		if err != nil {
			return nil, err
		}
		AddCommentToResponse(&gcr, content.Content)
	}

	return &gcr, nil
}

func Add(response *GetCommentResponse, memberID int64, like, hate int32) {
	response.MemberID = append(response.MemberID, memberID)
	response.Like = append(response.Like, like)
	response.Hate = append(response.Hate, hate)
}

func AddSubToResponse(response *GetCommentResponse, count, rootCount, AllCount int32) {
	response.Count = append(response.Count, count)
	response.RootCount = append(response.RootCount, rootCount)
	response.AllCount = append(response.RootCount, AllCount)
}

func AddIndexToResponse(response *GetCommentResponse, ID int64) {
	response.Id = append(response.Id, ID)
}

func AddCommentToResponse(response *GetCommentResponse, message string) {
	response.Message = append(response.Message, message)
}

func (c *CommentMessage) mustEmbedUnimplementedMessageServiceServer() {}

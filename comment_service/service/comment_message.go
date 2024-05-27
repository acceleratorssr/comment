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

// GetComment 先去redis内拿数据，如果没有则透传到mysql，并且使用kafka回源
func (c *CommentMessage) GetComment(ctx context.Context, request *GetCommentRequest) (*GetCommentResponse, error) {
	// TODO 先去redis内拿数据，如果没有则透传到mysql，并且使用kafka回源

	oidTypeSort := strconv.FormatInt(request.ObjID, 10) + "_" + strconv.FormatInt(int64(request.ObjType), 10) + "sortByDESC"
	get := global.Redis.ZRange(ctx, oidTypeSort, 0, -1)
	if get.Err() != nil {
		global.Log.Info("缓存失效")
	} else {
		// TODO 一次返回10条数据
		length := int32(len(get.Val()))
		var gcr GetCommentResponse
		for i := request.Offset; i < length; i++ {
			cs := CommentSubject{}
			err := proto.Unmarshal([]byte(get.Val()[i]), &cs)
			if err != nil {
				return nil, err
			}
			AddCommentToResponse(&gcr, 0, 0, 0, cs.Count, cs.RootCount, 0, 0, cs.State, 0, "")
		}

		return &gcr, nil
	}

	GCR := GetCommentResponse{}
	return &GCR, nil
}

func AddCommentToResponse(response *GetCommentResponse, root, parent, memberID int64, count, rootCount, like, hate, state int32, ip int64, message string) {
	response.Root = append(response.Root, root)
	response.Parent = append(response.Parent, parent)
	response.MemberID = append(response.MemberID, memberID)
	response.Count = append(response.Count, count)
	response.RootCount = append(response.RootCount, rootCount)
	response.Like = append(response.Like, like)
	response.Hate = append(response.Hate, hate)
	response.State = append(response.State, state)
	response.IP = append(response.IP, ip)
	response.Message = append(response.Message, message)
}

func (c *CommentMessage) mustEmbedUnimplementedMessageServiceServer() {}

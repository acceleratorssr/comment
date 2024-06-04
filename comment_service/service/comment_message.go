package service

import (
	"comment/global"
	"comment/models"
	"context"
	"errors"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
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
	var gcResp GetCommentResponse
	s := GetCommentRequest{
		ObjID:   request.ObjID,
		ObjType: request.ObjType,
		Offset:  request.Offset,
	}
	msg, _ := proto.Marshal(&s)
	// 先查对象缓存
	oidType := strconv.FormatInt(request.ObjID, 10) + "_" + strconv.FormatInt(int64(request.ObjType), 10)
	getSub := global.Redis.Get(ctx, oidType)
	if getSub.Err() != nil {
		global.Log.Info("对象表缓存失效")
		// 回源&返回数据
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func(gcResp *GetCommentResponse) {
			defer wg.Done()
			g, _ := backToSourceAndGetData(ctx, oidType, request.ObjID, int64(request.ObjType), request.Offset, msg)
			gcResp.Id = g.Id
			gcResp.Root = g.Root
			gcResp.Parent = g.Parent
			gcResp.MemberID = g.MemberID
			gcResp.Count = g.Count
			gcResp.RootCount = g.RootCount
			gcResp.AllCount = g.AllCount
			gcResp.Like = g.Like
			gcResp.Hate = g.Hate
			gcResp.Message = g.Message
		}(&gcResp)
		wg.Wait()
		//return backToSourceAndGetData(ctx, oidType, request.ObjID, int64(request.ObjType), request.Offset, msg)
		return &gcResp, nil
	} else {
		cs := CommentSubject{}
		err := proto.Unmarshal([]byte(getSub.Val()), &cs)
		if err != nil {
			return nil, err
		}
		AddSubToResponse(&gcResp, cs.Count, cs.RootCount, cs.AllCount)
	}

	// 再查评论索引缓存
	oidTypeSort := strconv.FormatInt(request.ObjID, 10) + "_" + strconv.FormatInt(int64(request.ObjType), 10) + "sortByDESC"
	getIndex := global.Redis.ZRange(ctx, oidTypeSort, 0, -1)
	if getIndex.Err() != nil {
		global.Log.Info("评论索引表缓存失效")
		// 回源&返回数据
		return backToSourceAndGetData(ctx, oidType, request.ObjID, int64(request.ObjType), request.Offset, msg)
	} else {
		length := int32(len(getIndex.Val()))
		// 从Offset开始，一次最多返回10条数据
		for i := request.Offset; i < length && i < request.Offset+10; i++ {
			s := getIndex.Val()[i]
			ids := strings.Split(s, "_")
			id, _ := strconv.Atoi(ids[0])
			memberID, _ := strconv.Atoi(ids[1])
			AddIndexToResponse(&gcResp, int64(id), int64(memberID))
		}
	}

	// 最后查评论内容缓存
	content := Content{}
	var comment []string
	for i := 0; i < len(gcResp.Id); i++ {
		getComment := global.Redis.Get(ctx, strconv.FormatInt(gcResp.Id[i], 10))
		if getComment.Err() != nil {
			global.Log.Info("评论内容表缓存失效")
			// 回源&返回数据
			return backToSourceAndGetData(ctx, oidType, request.ObjID, int64(request.ObjType), request.Offset, msg)
		} else {
			err := proto.Unmarshal([]byte(getComment.Val()), &content)
			if err != nil {
				return nil, err
			}
			comment = append(comment, content.Content)
		}
	}
	AddCommentToResponse(&gcResp, comment...)

	return &gcResp, nil
}

func Add(response *GetCommentResponse, like, hate int32) {
	response.Like = append(response.Like, like)
	response.Hate = append(response.Hate, hate)
}

func AddSubToResponse(response *GetCommentResponse, count, rootCount, allCount int32) {
	response.Count = append(response.Count, count)
	response.RootCount = append(response.RootCount, rootCount)
	response.AllCount = append(response.RootCount, allCount)
}

func AddIndexToResponse(response *GetCommentResponse, id, memberID int64) {
	response.Id = append(response.Id, id)
	response.MemberID = append(response.MemberID, memberID)
}

func AddCommentToResponse(response *GetCommentResponse, message ...string) {
	response.Message = append(response.Message, message...)
}

func FetchDataFromMySQL(gcResp *GetCommentResponse, obiID, objType int64, offset int32) {
	var csm models.CommentSubjectModels
	global.DB.Select("count, root_count, all_count").Where("obj_id = ? and obj_type = ?", obiID, objType).Take(&csm)
	AddSubToResponse(gcResp, csm.Count, csm.RootCount, csm.AllCount)

	var cim []models.CommentIndexModels
	global.DB.Select("id, member_id").Where("obj_id = ? and obj_type = ?", obiID, objType).Offset(int(offset)).Limit(10).Find(&cim)
	for i := 0; i < len(cim); i++ {
		AddIndexToResponse(gcResp, cim[i].ID, cim[i].MemberID)
	}

	var ccm []models.CommentContentModels
	global.DB.Select("message").Where("comment_id in ?", gcResp.Id).Find(&ccm)
	for j := 0; j < len(ccm); j++ {
		AddCommentToResponse(gcResp, ccm[j].Message)
	}

}

func backToSourceAndGetData(ctx context.Context, key string, obiID, ObjType int64, offset int32, msg []byte) (*GetCommentResponse, error) {
	ch := global.SF.DoChan(key, func() (interface{}, error) {
		backToSource(msg, key) //TODO 还是被执行多次
		// 透传到MySQL拿数据
		gcResp := GetCommentResponse{}
		FetchDataFromMySQL(&gcResp, obiID, ObjType, offset)
		return &gcResp, nil
	})

	select {
	case <-ctx.Done():
		return nil, errors.New("ctx_timeout")
	case data, _ := <-ch:
		//global.SF.Forget(key)
		return data.Val.(*GetCommentResponse), nil
	}
}

func (c *CommentMessage) mustEmbedUnimplementedMessageServiceServer() {}

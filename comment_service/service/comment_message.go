package service

import (
	"comment/global"
	"comment/models"
	"context"
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
	var gcr GetCommentResponse
	// 先查对象缓存
	oidType := strconv.FormatInt(request.ObjID, 10) + "_" + strconv.FormatInt(int64(request.ObjType), 10)
	getSub := global.Redis.Get(ctx, oidType)
	if getSub.Err() != nil {
		global.Log.Info("对象表缓存失效")
		// TODO 回源

		// TODO 透传到MySQL拿数据
		var csm models.CommentSubjectModels
		global.DB.Select("count, root_count, all_count").Where("obj_id = ? and obj_type = ?", request.ObjID, request.ObjType).Take(&csm)
		AddSubToResponse(&gcr, csm.Count, csm.RootCount, csm.AllCount)
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
		// TODO 回源
		// TODO 透传到MySQL拿数据
		var cim []models.CommentIndexModels
		global.DB.Select("id, member_id").Where("obj_id = ? and obj_type = ?", request.ObjID, request.ObjType).Offset(int(request.Offset)).Limit(10).Find(&cim)
		for i := 0; i < 10; i++ {
			AddIndexToResponse(&gcr, cim[i].ID, cim[i].MemberID)
		}
	} else {
		length := int32(len(getIndex.Val()))
		// 从Offset开始，一次最多返回10条数据
		for i := request.Offset; i < length && i < request.Offset+10; i++ {
			s := getIndex.Val()[i]
			ids := strings.Split(s, "_")
			id, _ := strconv.Atoi(ids[0])
			memberID, _ := strconv.Atoi(ids[1])
			AddIndexToResponse(&gcr, int64(id), int64(memberID))
		}
	}

	// 最后查评论内容缓存
	content := Content{}
	var comment []string
	for i := 0; i < len(gcr.Id); i++ {
		getComment := global.Redis.Get(ctx, strconv.FormatInt(gcr.Id[i], 10))
		if getComment.Err() != nil {
			global.Log.Info("评论内容表缓存失效")
			// TODO 回源
			// TODO 透传到MySQL拿数据
			var ccm []models.CommentContentModels
			global.DB.Select("id, member_id").Where("comment_id in ?", gcr.Id).Find(&ccm)
			for j := 0; j < 10; j++ {
				AddCommentToResponse(&gcr, ccm[j].Message)
			}

			break
		} else {
			err := proto.Unmarshal([]byte(getComment.Val()), &content)
			if err != nil {
				return nil, err
			}
			comment = append(comment, content.Content)
		}
		AddCommentToResponse(&gcr, comment...)
	}

	return &gcr, nil
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

func (c *CommentMessage) mustEmbedUnimplementedMessageServiceServer() {}

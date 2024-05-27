package comment_api

import (
	"comment/comment_service/service"
	"comment/global"
	"comment/models/res"
	"context"
	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	Root     int64 `json:"root,omitempty"`      //根评论ID,不为0是回复评论
	Parent   int64 `json:"parent,omitempty"`    //父评论ID,为0是root评论
	MemberID int64 `json:"member_id,omitempty"` //发表者用户id
	ObjID    int64 `json:"obj_id,omitempty"`    //对象id 即该条评论对应的对象 冗余设计
	State    int8  `json:"state,omitempty"`     //状态(0-正常、1-隐藏)
	ObjType  int8  `json:"obj_type,omitempty"`  //对象类型 冗余设计
	//Floor    int32 `json:"floor,omitempty"`     //评论楼层

	IP int64 `json:"ip,omitempty"`

	Message string `json:"message,omitempty"`
}

// CreateComment 创建单条评论
func (CommentApi) CreateComment(c *gin.Context) {
	var ccr CreateCommentRequest
	err := c.ShouldBindJSON(&ccr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, CreateCommentRequest{}, c)
		return
	}

	client := service.NewMessageServiceClient(global.GrpcConn)
	resp, err := client.CreateCommentMessage(context.Background(), &service.CreateMessageRequest{
		ObjId:    ccr.ObjID,
		MemberId: ccr.MemberID,
		State:    service.State(ccr.State),
		ObjType:  service.ObjType(ccr.ObjType),
		Root:     ccr.Root,
		Parent:   ccr.Parent,
		//Floor:    ccr.Floor,
		Ip:      ccr.IP,
		Comment: ccr.Message,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithAll(resp, "ok", c)
}

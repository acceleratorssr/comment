package comment_api

import (
	"comment/comment_service/service"
	"comment/global"
	"comment/models/res"
	"context"
	"github.com/gin-gonic/gin"
)

type GetCommentRequest struct {
	ObjID   int64 `json:"obj_id" binding:"required" msg:"缺少对象id"`
	ObjType int8  `json:"obj_type" binding:"required" msg:"缺少对象类型"`
	Offset  int32 `json:"offset"`
}

// GetComment 获取评论
func (CommentApi) GetComment(c *gin.Context) {
	var gcr GetCommentRequest
	err := c.ShouldBindJSON(&gcr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, GetCommentRequest{}, c)
		return
	}

	client := service.NewMessageServiceClient(global.GrpcConn)
	resp, err := client.GetComment(context.Background(), &service.GetCommentRequest{
		ObjID:   gcr.ObjID,
		ObjType: service.ObjType(gcr.ObjType),
		Offset:  gcr.Offset,
	})
	if err != nil {
		global.Log.Error("client.GetComment -> ", err)
		return
	}

	res.OKWithData(resp, c)
}

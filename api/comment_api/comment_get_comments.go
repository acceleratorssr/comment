package comment_api

import (
	"comment/comment_service/service"
	"comment/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetComment 获取评论
func (CommentApi) GetComment(c *gin.Context) {
	client := service.NewMessageServiceClient(global.GrpcConn)
	resp, err := client.GetComment(context.Background(), &service.GetCommentRequest{
		ObjID: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"name": resp,
		},
	})
}

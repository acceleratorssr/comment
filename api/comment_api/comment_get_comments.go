package comment_api

import (
	"comment/comment_service/service"
	"comment/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetComment 获取评论
func (CommentApi) GetComment(c *gin.Context) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(global.Grpc.Addr, opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			global.Log.Error(err)
		}
	}(conn)

	client := service.NewMessageServiceClient(conn)

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

package comment_api

import (
	"github.com/gin-gonic/gin"
)

func (CommentApi) GetComment(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"name": "comment",
			"id":   id,
		},
	})
}

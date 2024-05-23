package routers

import "comment/api"

func (RG RouterGroup) CommentRouter() {
	CommentController := api.Groups.CommentApi
	RG.Router.POST("/comments", CommentController.GetComment)
	RG.Router.POST("/comment", CommentController.CreateComment)
}

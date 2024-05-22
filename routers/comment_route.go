package routers

import "comment/api"

func (RG RouterGroup) CommentRouter() {
	CommentController := api.Groups.CommentApi
	RG.Router.GET("/comment/:id", CommentController.GetComment)
	//RG.Router.POST("/comment", CommentController.CreateComment)
}

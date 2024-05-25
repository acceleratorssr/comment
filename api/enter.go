package api

import (
	"comment/api/comment_api"
	"comment/api/subject_api"
	"comment/api/user_api"
)

type Group struct {
	CommentApi comment_api.CommentApi
	UserApi    user_api.UserApi
	Subject    subject_api.SubjectApi
}

var Groups = new(Group)

package routers

import "comment/api"

func (RG RouterGroup) SubjectRouter() {
	subjectApi := api.Groups.Subject
	RG.Router.POST("/subject_create", subjectApi.SubjectCreateView)
}

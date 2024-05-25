package subject_api

import (
	"comment/global"
	"comment/models"
	"comment/models/res"
	"github.com/gin-gonic/gin"
)

type subjectCreateRequest struct {
	ObjID    int64 `json:"obj_id,omitempty"`
	MemberID int64 `json:"member_id,omitempty"`
	State    int8  `json:"state,omitempty"`
	ObjType  int8  `json:"obj_type,omitempty"` //0为视频 1为文章
}

func (SubjectApi) SubjectCreateView(c *gin.Context) {
	var scq subjectCreateRequest

	err := c.ShouldBindJSON(&scq)
	if err != nil {
		global.Log.Error(err)
		return
	}
	csm := models.CommentSubjectModels{
		ObjID:    scq.ObjID,
		MemberID: scq.MemberID,
		State:    scq.State,
		ObjType:  scq.ObjType,
	}

	err = global.DB.Create(&csm).Error
	if err != nil {
		global.Log.Warn("subjectCreateView -> ", err)
		return
	}

	res.OKWithMessage("成功创建对象", c)

}

package subject_api

import "github.com/gin-gonic/gin"

// Subject 业务实现对象接口，subject_api内的方法仅用于测试comment的正确性
type Subject interface {
	SubjectCreateView(c *gin.Context)
}

type SubjectApi struct {
}

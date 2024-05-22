package comment_api

//
//import (
//	"comment/global"
//	"comment/models"
//	"comment/models/res"
//	"comment/utils/jwts"
//	"github.com/gin-gonic/gin"
//)
//
//// MessageRequest 目前仅支持发送
//type MessageRequest struct {
//	ReceiveUserID uint   `json:"receive_user_id" binding:"required" msg:"请输入私聊对象ID"` // 接收者ID
//	Content       string `json:"content" binding:"required" msg:"不能发送空消息"`           // 消息内容
//}
//
//func (CommentApi) GetComment(c *gin.Context) {
//	// 已登录的用户，选择一个用户（可以是自己）发送一条消息
//	var MR MessageRequest
//	var userModel models.UserModels
//
//	err := c.ShouldBindJSON(&MR)
//	if err != nil {
//		global.Log.Errorln("MessageSend -> 参数绑定失败", err)
//		res.FailWithError(err, MessageRequest{}, c)
//		return
//	}
//
//	_permission, _ := c.Get("parseToken")
//
//	// 注意_permission的类型是 *jwts.Permission
//	permission := _permission.(*jwts.CustomClaims)
//	// 用户登录后，被删除，token挂掉，应该不用判断发送方的状态？
//
//	// 查看接收方是否存在
//	// 但前端如果是微信通讯录那种发起聊天的方式，应该也不用判断用户存在与否？
//	err = global.DB.Take(&userModel, MR.ReceiveUserID).Error
//	if err != nil {
//		global.Log.Errorln("MessageSend -> 查无此用户", err)
//		res.FailWithMessage("参数绑定失败", c)
//		return
//	}
//
//	err = global.DB.Create(&models.MessageModels{
//		SendUserID:    permission.UserID,
//		ReceiveUserID: MR.ReceiveUserID,
//		Content:       MR.Content,
//	}).Error
//	if err != nil {
//		global.Log.Errorln("MessageSend -> 聊天记录保存失败", err)
//		res.FailWithMessage("聊天记录发送失败", c)
//		return
//	}
//	res.OKWithMessage("发送成功", c)
//}
//
//func (CommentApi) CreateComment(c *gin.Context) {
//
//}

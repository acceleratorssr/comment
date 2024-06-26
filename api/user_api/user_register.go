package user_api

import (
	"comment/global"
	"comment/models"
	"comment/models/res"
	"comment/models/stype"
	"comment/utils/jwts"
	"comment/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	Username string `json:"user_name" msg:"缺少用户名"`
	NickName string `json:"nick_name" msg:"缺少昵称（后续随时可改）"`
	Password string `json:"password"  msg:"缺少密码"`
	IP       string `json:"ip"`
}

// UserRegisterView 是一个API视图，用于处理用户注册的请求
//
// @Summary 用户注册
// @Description 用户注册视图，需要用户名、昵称和密码。此处前端验证两次输入密码正确后，才会传回信息；会查表以防用户名重复，头像默认，注册成功后自动登录。
// @Tags 用户
// @Accept json
// @Produce json
// @Param UserRegisterRequest body UserRegisterRequest true "用户名，昵称，密码，IP地址"
// @Success 200 {string} string "注册成功"
// @Router /api/user_register [post]
func (UserApi) UserRegisterView(c *gin.Context) {
	// 注册用户
	var URR UserRegisterRequest
	var userModel models.UserModels

	err := c.ShouldBindJSON(&URR)
	if err != nil {
		global.Log.Warnln("注册失败 UserRegisterView -> ", err)
		res.FailWithError(err, UserRegisterRequest{}, c)
		return
	}

	// 注册到mysql，获取id
	err = global.DB.Take(&userModel, "username = ?", URR.Username).Error
	if err == nil {
		global.Log.Warn("注册 -> 用户名已存在", err)
		res.FailWithMessage("用户名已存在", c)
		return
	}

	URR.Password = pwd.HashAndSalt(URR.Password)
	user := models.UserModels{
		Username:   URR.Username,
		NickName:   URR.NickName,
		Password:   URR.Password,
		Token:      "",
		IP:         "",
		SignStatus: stype.SignNotStatus,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		res.FailWithMessage("用户名注册失败", c)
		return
	}
	// 返回用户信息
	token, err := jwts.GenToken(jwts.JwtPayload{
		Username:    URR.Username,
		UserID:      uint(user.ID),
		Permissions: int(user.Permission),
		NickName:    URR.NickName,
	})
	if err != nil {
		global.Log.Error("token -> 生成失败", err)
		res.FailWithMessage("登录失败", c)
		return
	}
	res.OKWithData(token, c)
}

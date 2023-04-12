// @User CPR
package api

import (
	"VideoEdit/model/req"
	"VideoEdit/utils"
	"VideoEdit/utils/r"
	"github.com/gin-gonic/gin"
)

type UserAuth struct {
}

// 用户登录
func (*UserAuth) Login(c *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](c)
	loginVo, code := userService.Login(c, loginReq.UserName, loginReq.PassWord)
	r.SendData(c, code, loginVo)
}

// 用户注册
func (*UserAuth) Register(c *gin.Context) {
	userService.Register(utils.BindValidJson[req.Register](c))
	r.Success(c)
}

// 发送验证码
func (*UserAuth) SendCode(c *gin.Context) {
	userService.SendCode(c, utils.BindValidJson[req.RegisterCode](c))
	r.Success(c)
}

// 根据Token获取用户信息
//func (*UserAuth) GetInfo(c *gin.Context) {
//	user_id, _ := c.Get("user_id")
//	r.SuccessData(c, userService.GetInfo(user_id.(int)))
//}

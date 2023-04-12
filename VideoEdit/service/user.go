// @User CPR
package service

import (
	"VideoEdit/dao"
	"VideoEdit/model"
	"VideoEdit/model/dto"
	"VideoEdit/model/req"
	"VideoEdit/model/resp"
	"VideoEdit/utils"
	"VideoEdit/utils/r"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

const (
	KEY_VERIFY_CODE = "verify_code_"
)

type User struct {
}

// 登录
func (*User) Login(c *gin.Context, username, password string) (resp.LoginVo, int) {
	// 检查用户是否存在
	userAuth := dao.GetOne[model.UserAuth]("username", username)
	if userAuth.Id == 0 {
		return *new(resp.LoginVo), r.ERROR_USER_NOT_EXIST
	}
	// 检查密码是否正确
	if !utils.Encryptor.BcryptCheck(password, userAuth.Password) {
		return *new(resp.LoginVo), r.ERROR_PASSWORD_WRONG
	}
	// 获取用户详细信息 DTO
	userDetailDTO := convertUserDetailDTO(c, userAuth)

	// 登录信息正确， 生成Token
	// TODO: 目前只给用户设定一个角色, 获取第一个值就行, 后期优化: 给用户设置多个角色
	// UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	uuid := utils.Encryptor.MD5(userDetailDTO.IpAddress + userDetailDTO.Browser + userDetailDTO.OS)
	// 生成Token
	token, err := utils.GetJWT().GenToken(userAuth.Id, uuid)
	if err != nil {
		utils.Logger.Info("登录时生成Token失败", zap.Error(err))
		return *new(resp.LoginVo), r.ERROR_TOKEN_CREATE
	}
	userDetailDTO.Token = token
	// todo: 个人感觉这个应该在用户退出登录时记录
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{Id: userAuth.Id},
		IpAddress:     userDetailDTO.IpAddress,
		IpSource:      userDetailDTO.IpSource,
		LastLoginTime: userDetailDTO.LastLoginTime,
	})

	// 保存用户信息到Session和Redis中
	session := sessions.Default(c)
	// ! session 中只能存储字符串
	sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{UserDetailDTO: userDetailDTO})
	utils.Redis.Set(KEY_USER+uuid, sessionInfoStr, 10*time.Minute)
	session.Set(KEY_USER+uuid, sessionInfoStr)

	session.Save()
	return userDetailDTO.LoginVo, r.OK
}

func convertUserDetailDTO(c *gin.Context, userAuth model.UserAuth) dto.UserDetailDTO {
	// 获取用户IP相关信息
	ipAddress := utils.IP.GetIpAddress(c)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)
	browser, os := "unknown", "unknown"

	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}

	// 获取用户详细信息
	userInfo := dao.GetOne[model.User]("id", userAuth.Id)

	return dto.UserDetailDTO{
		LoginVo: resp.LoginVo{
			UserId:        userAuth.Id,
			Email:         userAuth.Email,
			NickName:      userInfo.NickName,
			IpAddress:     ipAddress,
			IpSource:      ipSource,
			LastLoginTime: time.Now(),
		},
		Password:  userAuth.Password,
		Browser:   browser,
		OS:        os,
		IsDisable: userInfo.IsDisable,
	}
}

func (*User) Logout(c *gin.Context, username string) {
	userAuth := dao.GetOne[model.UserAuth]("username", username)
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{Id: userAuth.Id},
		LastLoginTime: time.Now(),
	})

	uuid := utils.GetFromContext[string](c, "uuid")
	session := sessions.Default(c)
	session.Delete(KEY_USER + uuid)
	session.Save()
	utils.Redis.Del(KEY_USER + uuid)
}

// 用户注册
func (*User) Register(req req.Register) int {
	// 检查用户名已存在 则该账号已经注册过了
	if exist := checkUserExistByName(req.UserName); exist {
		utils.Logger.Info("注册时检查用户名已存在", zap.String("username", req.UserName))
		return r.ERROR_USER_NAME_USED
	}
	// VerifyCode 验证码
	if !VerifyCode(req.Code, req.Email) {
		return r.ERROR_VERIFY_CODE
	}
	user := &model.User{
		NickName:   "user:" + req.UserName,
		Email:      req.Email,
		IsDisable:  0,
		UpVideoNum: 0,
	}
	dao.Create(&user)
	// 设置用户默认角色
	dao.Create(&model.UserAuth{
		UserId:        user.Id,
		Username:      req.UserName,
		Email:         req.Email,
		Password:      utils.Encryptor.BcryptHash(req.Password),
		LoginType:     1,
		LastLoginTime: time.Now(), // 注册时会更新 "上次登录时间"
	})
	return r.OK
}

// 检查用户名是否已经存在
func checkUserExistByName(username string) bool {
	existUser := dao.First[model.UserAuth]("username = ?", username)
	return existUser.Id != 0
}

// 检查邮箱是否已经存在
func checkUserExistByEmail(email string) bool {
	existUser := dao.First[model.UserAuth]("email = ?", email)
	return existUser.Id != 0
}

func VerifyCode(code, email string) bool {
	// 从Redis中获取验证码
	realCode := utils.Redis.HGetCode(KEY_VERIFY_CODE, email)
	return realCode == code
}

func (*User) SendCode(c *gin.Context, req req.RegisterCode) int {
	// 检查用户名已存在 则该账号已经注册过了
	if exist := checkUserExistByName(req.UserName); exist {
		utils.Logger.Info("注册时检查用户名已存在", zap.String("username", req.UserName))
		return r.ERROR_USER_NAME_USED
	}
	// 检查邮箱是否已经注册过了
	if exist := checkUserExistByEmail(req.Email); exist {
		utils.Logger.Info("注册时检查邮箱已存在", zap.String("email", req.Email))
		return r.ERROR_EMAIL_USED
	}
	// 发送验证码
	code := genCode()

	utils.Redis.HSet(KEY_VERIFY_CODE, req.Email, code, time.Duration(5)*time.Minute)
	if err := utils.SendEmail(req.Email, "账号注册验证码", code); err != nil {
		utils.Logger.Info("发送验证码失败", zap.Error(err))
		return r.ERROR_EMAIL_SEND_CODE
	}
	return r.OK
}

func genCode() (code string) {
	rand.Seed(time.Now().UnixNano())
	code = ""
	for i := 0; i < 6; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}
	return
}

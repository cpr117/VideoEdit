// @User CPR
package r

// 错误码汇总
const (
	OK   = 0
	FAIL = 500

	// code= 9000... 通用错误
	ERROR_REQUEST_PARAM = 9001
	ERROR_REQUEST_PAGE  = 9002
	ERROR_INVALID_PARAM = 9003
	ERROR_DB_OPE        = 9004
	ERROR_REQUEST_FORM  = 9005

	// code 91... 文件相关错误
	ERROR_FILE_UPLOAD  = 9100
	EEROR_FILE_RECEIVE = 9101

	// code= 1000... 用户模块的错误
	ERROR_USER_NAME_USED  = 1001
	ERROR_PASSWORD_WRONG  = 1002
	ERROR_USER_NOT_EXIST  = 1003
	ERROR_USER_NO_RIGHT   = 1009
	ERROR_OLD_PASSWORD    = 1010
	ERROR_EMAIL_USED      = 1011
	ERROR_EMAIL_CODE      = 1012
	ERROR_EMAIL_SEND_CODE = 1013
	ERROR_VERIFY_CODE     = 1014

	// code = 1200.. 鉴权相关错误
	ERROR_TOKEN_NOT_EXIST  = 1201
	ERROR_TOKEN_RUNTIME    = 1202
	ERROR_TOKEN_WRONG      = 1203
	ERROR_TOKEN_TYPE_WRONG = 1204
	ERROR_TOKEN_CREATE     = 1205
	ERROR_PERMI_DENIED     = 1206
	FORCE_OFFLINE          = 1207
	LOGOUT                 = 1208

	// code=8000 ... 视频模块的错误
	ERROR_VIDEOID_WRONG   = 1301
	ERROR_VIDEO_NOT_EXIST = 1302
)

var codeMsg = map[int]string{
	OK:   "OK",
	FAIL: "FAIL",

	ERROR_REQUEST_PARAM: "请求参数格式错误",
	ERROR_REQUEST_PAGE:  "分页参数错误",
	ERROR_INVALID_PARAM: "不合法的请求参数",
	ERROR_DB_OPE:        "数据库操作异常",

	EEROR_FILE_RECEIVE: "文件接收失败",
	ERROR_FILE_UPLOAD:  "文件上传失败",

	ERROR_USER_NAME_USED:  "用户名已存在",
	ERROR_USER_NOT_EXIST:  "该用户不存在",
	ERROR_PASSWORD_WRONG:  "密码错误",
	ERROR_USER_NO_RIGHT:   "该用户无权限",
	ERROR_OLD_PASSWORD:    "旧密码不正确",
	ERROR_EMAIL_USED:      "邮箱已被注册",
	ERROR_EMAIL_CODE:      "邮箱验证码错误",
	ERROR_EMAIL_SEND_CODE: "邮箱验证码发送失败",

	ERROR_TOKEN_NOT_EXIST:  "TOKEN 不存在，请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN 已过期，请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN 不正确，请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN 格式错误，请重新登陆",
	ERROR_TOKEN_CREATE:     "TOKEN 生成失败",
	ERROR_PERMI_DENIED:     "权限不足",
	FORCE_OFFLINE:          "您已被强制下线",
	LOGOUT:                 "您已退出登录",

	ERROR_VIDEOID_WRONG:   "视频ID错误",
	ERROR_VIDEO_NOT_EXIST: "视频不存在",
}

func GetMsg(code int) string {
	return codeMsg[code]
}

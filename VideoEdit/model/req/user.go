// @User CPR
package req

type Login struct {
	UserName string `json:"user_name" validate:"required,min=6,max=20" label:"用户名"`
	PassWord string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

type Register struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Code     string `json:"code" validate:"required,min=6,max=6" label:"验证码"`
	Phone    string `json:"phone" validate:"required,min=11,max=11" label:"手机号"`
}

type RegisterCode struct {
	UserName string `json:"username" validate:"required,min=6,max=20" label:"用户名"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
}

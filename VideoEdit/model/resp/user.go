// @User CPR
package resp

import "time"

// LoginVo
// 登录返回信息 一些比较简洁的信息
type LoginVo struct {
	UserId        int       `json:"user_id"`
	Email         string    `json:"email"` // 邮箱
	NickName      string    `json:"nickname"`
	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	Token         string    `json:"token"`
	LastLoginTime time.Time `json:"last_login_time"`
}

// @User CPR
package model

import "time"

// 用户账户信息
type UserAuth struct {
	Universal     `mapstructure:",squash"`
	UserId        int       `json:"user_id" gorm:"column:user_id;type:int(11);not null;default:0;comment:'用户ID'"`
	Username      string    `gorm:"type:varchar(50);comment:用户名" json:"username"`
	Password      string    `gorm:"type:varchar(100);comment:密码" json:"password"`
	Email         string    `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	LoginType     int8      `gorm:"type:tinyint(1);comment:登录类型" json:"login_type"`
	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间" json:"last_login_time"`
}

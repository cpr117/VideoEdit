// @User CPR
package model

// 用户信息只包含了用户 个人主页内显示的信息
type User struct {
	Universal
	NickName   string `gorm:"type:varchar(255);not null;comment:昵称" validate:"" json:"nick_name"`
	Email      string `gorm:"type:varchar(255);not null;comment:邮箱" validate:"email" json:"email"`
	UpVideoNum int    `gorm:"int(11);comment:上传视频数" json:"up_video_num"`
	IsDisable  int8   `gorm:"type:tinyint(1);comment:是否禁用(0-否 1-是)" json:"is_disable"`
}

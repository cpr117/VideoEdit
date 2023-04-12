// @User CPR
package interval

import (
	"VideoEdit/dao"
	"VideoEdit/utils"
)

func InitGlobalVariable() {
	// 初始化Viper
	utils.InitViper()
	// 初始化Logger
	utils.InitLogger()
	// 初始化数据库DB
	dao.DB = utils.InitMySQLDB()
	// 初始化Redis
	utils.InitRedis()
	// 初始化VideoManager
	utils.InitVM()
	//// 初始化Casbin
	//utils.InitCasbin(dao.DB)
	//// 初始化其他参数
	//utils.InitOther()
}

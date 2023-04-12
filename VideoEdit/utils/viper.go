// @User CPR
package utils

import (
	"VideoEdit/config"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func InitViper() {
	var configPath string
	// 命令行输入设置的配置文件地址
	flag.StringVar(&configPath, "c", "config/config.toml", "choose config file.")
	flag.Parse()
	log.Println("命令行读取参数，配置文件路径为:%s\n", configPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败, err:", err)
	}

	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件解析失败, err: ", err)
	}
	// 配置文件热重载
	v.WatchConfig() // 间厅柜配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("检测到配置文件发生变化，重新加载配置文件")
		if err := v.Unmarshal(&config.Cfg); err != nil {
			log.Panic("配置文件解析失败, err: ", err)
		}
	})
	log.Println("配置文件加载成功")
}

// @User CPR
package config

var Cfg *Config

// 配置文件的结构体
type Config struct {
	Server  Server
	JWT     JWT
	Mysql   Mysql
	Redis   Redis
	Session Session
	Upload  Upload
	Zap     Zap
	Qiniu   Qiniu
	Video   Video
	Code    Code
	Email   Email
}

type Email struct {
	User     string
	Password string
	Host     string
	Port     string
}

type Code struct {
	Length   int
	Duration int
}

type Video struct {
	VideoPath string
	MaxByte   int64
}

type Zap struct {
	Level        string // 日志级别
	Prefix       string // 日志前缀
	Format       string // 输出格式
	Directory    string // 日志目录
	MaxAge       int    // 日志留存时间
	ShowLine     bool   // 显示行
	LogInConsole bool   // 输出控制台
}

type Mysql struct {
	Host     string // 服务器地址
	Port     string // 端口
	Config   string // 高级配置
	Dbname   string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
	LogMode  string // 日志级别
}

type Redis struct {
	DB       int    // 指定 Redis 数据库
	Addr     string // 服务器地址:端口
	Password string // 密码
}

type Server struct {
	GLimit    uint8
	AppMode   string
	BackPort  string
	FrontPort string
}

type JWT struct {
	Secret string // JWT 签名
	Expire int64  // 过期时间
	Issuer string // 签发者
}

type Session struct {
	Name   string
	Salt   string
	MaxAge int
}

type Upload struct {
	StorePath    string // 本地文件存储路径
	CompletePath string // 本地文件存储路径
}

type Qiniu struct {
	ImgPath       string // 外链链接
	Zone          string // 存储区域
	Bucket        string // 空间名称
	AccessKey     string // 秘钥AK
	SecretKey     string // 秘钥SK
	UseHTTPS      bool   // 是否使用https
	UseCdnDomains bool   // 上传是否使用 CDN 上传加速
}

// @User CPR
package routes

import (
	"VideoEdit/config"
	"VideoEdit/routes/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 后台管理界面接口路由
func BackRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	//// 使用本地文件上传 需要静态文件服务
	//if config.Cfg.Upload.OssType == "local" {
	//	r.Static("/public", "./public")
	//	r.StaticFS("/dir", http.Dir("./public")) // 将 public 目录内的文件列举展示
	//}

	r.Use(middlewares.Logger())             // 自定义的zap日志中间件
	r.Use(middlewares.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middlewares.Cors())               // 跨域中间件

	// 基于 cookie 存储 session
	store := cookie.NewStore([]byte(config.Cfg.Session.Salt))

	// session 存储时间和JWT（json web token）过期时间一致
	// 这是两种不同的用户检测机制 一者存储在客户端 一者存储在服务端
	store.Options(sessions.Options{MaxAge: int(config.Cfg.JWT.Expire) * 3600})
	r.Use(sessions.Sessions(config.Cfg.Session.Name, store)) // 使用session中间件

	// 无需鉴权的借口
	base := r.Group("/")
	{
		// TODO: 后台登录
		// 这边默认最高用户在一开始就创建好了 下面的用户由高级后台用户向下创建 因此没有注册界面
		base.POST("/login", userAPI.Login)
		base.POST("/sendCode", userAPI.SendCode)
		base.POST("/register", userAPI.Register)

		// 需要鉴权的借口
		auth := base.Group("/api")
		// !注意使用中间件的顺序
		auth.Use(middlewares.JWTAuth())      // JWT 鉴权中间件
		auth.Use(middlewares.ListenOnline()) // 监听在线用户
		auth.Use(middlewares.OperationLog()) // 记录操作日志
		{
			auth.POST("/editvideo", videoAPI.EditVideo)
		}
	}

	return r
}

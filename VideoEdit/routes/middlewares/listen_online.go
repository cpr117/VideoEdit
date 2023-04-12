// @User CPR
package middlewares

import (
	"VideoEdit/model/dto"
	"VideoEdit/utils"
	"VideoEdit/utils/r"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	KEY_USER = "user:"
)

// 监听在线状态中间件
func ListenOnline() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := utils.GetFromContext[string](c, "uuid")
		userKey := KEY_USER + uuid

		session := sessions.Default(c)
		var sessionInfo dto.SessionInfo
		// 尝试从 redis 中取用户的 SessionInfo, 取不到则从 session 中获取
		if s, err := utils.Redis.GetResult(userKey); err == nil {
			utils.Json.Unmarshal(s, &sessionInfo)
		} else { // * session 中还是取不到, 就返回 401 让前端退回登录界面
			// ! 考虑一下
			val := session.Get(userKey)
			if val == nil {
				r.Send(c, http.StatusUnauthorized, r.ERROR_TOKEN_RUNTIME, nil)
				c.Abort()
				return
			}
			utils.Json.Unmarshal(val.(string), &sessionInfo)
		}

		// * 每次发送请求会更新 Redis 中的在线状态: 重新计算 10 分钟
		utils.Redis.Set(KEY_USER+uuid, utils.Json.Marshal(sessionInfo), 10*time.Minute)
		fmt.Println("listen_online")
		c.Next()
	}
}

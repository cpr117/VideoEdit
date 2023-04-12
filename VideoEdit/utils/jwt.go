// @User CPR
package utils

import (
	"VideoEdit/config"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义token相关error
var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)

// 定义JWT中存储的信息
type MyClaims struct {
	UserId int    `json:"user_id"`
	UUID   string `json:"uuid"`
	jwt.RegisteredClaims
}

type MyJWT struct {
	Secret []byte
}

// GetJWT
// JWT 工具类
func GetJWT() *MyJWT {
	return &MyJWT{[]byte(config.Cfg.JWT.Secret)}
}

// GenToken
// 生成 JWT
func (j *MyJWT) GenToken(userId int, uuid string) (string, error) {
	claims := MyClaims{
		UserId: userId,
		UUID:   uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Cfg.JWT.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Cfg.JWT.Expire) * time.Hour)),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的 secret 签名并获得完整编码后的字符串 token
	return token.SignedString(j.Secret)
}

// ParseToken
// 解析 JWT
func (j *MyJWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	if err != nil {
		if vError, ok := err.(*jwt.ValidationError); ok {
			if vError.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if vError.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if vError.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	// 校验 token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

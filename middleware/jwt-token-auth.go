package middleware

import (
	"donkey-ucenter/lib/libutils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

const (
	// user token
	UserAuthKey   = "X-User-Authorization"
	UserJwtSecret = "USER_JWT_SECRET"
	UserCtxKey    = "user_id" // 存用户id的  key， JWT 和 gin.Context 都用此key
)

func init() {

}

func jwtTokenWare(tokenKey string, secret string, storeKey string) func(c *gin.Context) {
	// storeKey gin 和 jwt claims中的key
	// tokenKey header头 或者 cookie 包含jwt串的key
	// secret jwt解密密钥
	return func(c *gin.Context) {
		claims, err := libutils.Jwt.CheckToken(c, tokenKey, secret)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": ""})
			c.Abort()
		}

		value := claims[storeKey]
		c.Set(storeKey, cast.ToInt(value))
		c.Next()
	}
}

func UserToken() func(c *gin.Context) {
	return jwtTokenWare(UserAuthKey, UserJwtSecret, UserCtxKey)
}

func GetUserId(c *gin.Context) int {
	return c.GetInt(UserCtxKey)
}

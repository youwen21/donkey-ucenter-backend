package main

import (
	"donkey-ucenter/lib/libutils"
	"donkey-ucenter/middleware"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func main() {
	tokenStr, _ := TestGenJwt()
}

func TestGenJwt() (string, error) {
	userId := 1
	tokenString, err := libutils.Jwt.GenToken(middleware.UserAuthKey, jwt.MapClaims{middleware.UserCtxKey: userId, "exp": time.Now().Unix() + 86400*7}) // 记住登录7天，或者7天不关闭网页，登录最长有效期7天

	fmt.Println(tokenString, err)
	return tokenString, err
}

func TestParseJwt() (jwt.MapClaims, error) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjUyMzc4MjgsInVzZXJfaWQiOjF9.UjBzjofwL64IY5uBjfUjmGG0dC2Q_WlKwUFPV4aHIm4"
	libutils.Jwt.ParseToken(tokenString)
}

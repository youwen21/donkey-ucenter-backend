package router

import (
	"donkey-ucenter/app/handler/api"
	"donkey-ucenter/middleware"
	"github.com/gin-gonic/gin"
)

/*  */

func initApi(engine *gin.Engine) {
	publicApiGroup := engine.Group("/api").Use(middleware.Cors.GinCors())

	// 注册登录， 可能有多种形式
	// 用户名密码注册， 邮箱注册， 手机号注册， 微信扫码直接注册登录
	// 密码登录，邮箱登录，手机号登录，扫码登录

	// 基础模式 用户名密码注册
	publicApiGroup.POST("/auth/register", api.AuthHandler.Register)
	publicApiGroup.POST("/auth/login", api.AuthHandler.Login)
	publicApiGroup.POST("/auth/logout", api.AuthHandler.Logout)

	// 找回密码
	publicApiGroup.POST("/reset-pwd", api.UserPublicHdl.ResetPwdByEmail)
	publicApiGroup.POST("/reset-pwd-confirm", api.UserPublicHdl.ResetPwdConfirm)

	privateApiGroup := engine.Group("/api/user").Use(middleware.Cors.GinCors()).Use(middleware.UserToken())
	privateApiGroup.GET("info", api.UserHdl.Info)
	privateApiGroup.POST("update-info", api.UserHdl.UpdateInfo)
	privateApiGroup.POST("change-pwd", api.UserHdl.ChangePwd)
	privateApiGroup.POST("bind-email", api.UserHdl.BindEmail)
	privateApiGroup.POST("bind-email-confirm", api.UserHdl.BindEmailConfirm)

}

package router

import (
	"donkey-ucenter/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	//gin.SetMode(gin.DebugMode)
	engine.GET("/ping", func(gtx *gin.Context) { gtx.String(200, "pong") })
	engine.OPTIONS("/*options_support", middleware.Cors.GinCors())

	initApi(engine)
}

package api

import (
	"donkey-ucenter/app/service/auth"
	"donkey-ucenter/app/service/auth/auth_def"
	"donkey-ucenter/middleware"
	"donkey-ucenter/req-resp/appresp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandler struct {
}

var (
	AuthHandler = new(authHandler)
)

func (hdl *authHandler) Register(c *gin.Context) {
	form := new(auth_def.Form)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	tokenString, userInfo, err := auth.Srv.Register(form)
	if err != nil {
		c.JSON(http.StatusOK, appresp.Err(err))
		return
	}
	// 前端自定义参数，原样返回
	// 处理 query params 铭感信息
	qValues := c.Request.URL.Query()
	if qValues.Get("password") != "" {
		qValues.Set("password", "******")
	}

	// 设置cookie, 支持ajax
	c.SetSameSite(http.SameSiteNoneMode)
	// todo Host 白名单

	// cookie有效期 7天免费登录 选项
	cookieMaxAge := 0
	if form.Remember {
		cookieMaxAge = 86400 * 7
	}
	c.SetCookie(middleware.UserAuthKey, tokenString, cookieMaxAge, "/", c.Request.Host, true, true)

	LoginRes := auth_def.LoginRes{
		Params:   qValues,
		Token:    tokenString,
		UserInfo: userInfo,
	}
	c.JSON(http.StatusOK, appresp.Reps(LoginRes, nil))

}

// Login 详细说明 ：
// 跨域基础条件 1
// 当AJAX请求的协议、域名或端口与目标服务器不一致时，就会触发跨域限制1。
//
// secure属性的作用 1
// Cookie.secure=true表示该Cookie仅通过HTTPS协议传输，这是跨域请求中携带Cookie的必要条件之一1。
// 需与SameSite=None属性配合使用，才能在跨域请求中发送Cookie1。
//
// 其他关键配置 1
// 前端设置 ：AJAX请求需设置withCredentials=true（XMLHttpRequest）或credentials: 'include'（Fetch API）1。
// 后端配置 ：服务器需在CORS响应中设置Access-Control-Allow-Credentials=true，并明确指定Access-Control-Allow-Origin为具体源（非*）1。
//
// 注意事项 2
// SameSite=None和Secure属性需同时设置，否则Cookie不会在跨域请求中发送1。
// JSONP已被淘汰，不支持跨域请求携带Cookie1。
func (hdl *authHandler) Login(c *gin.Context) {
	loginForm := &auth_def.Form{}
	err := c.ShouldBind(&loginForm)
	if nil != err {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
	}

	tokenString, userInfo, err := auth.Srv.Login(loginForm)
	if err != nil {
		c.JSON(http.StatusOK, appresp.Err(err))
		return
	}

	// 前端自定义参数，原样返回
	// 处理 query params 铭感信息
	qValues := c.Request.URL.Query()
	if qValues.Get("password") != "" {
		qValues.Set("password", "******")
	}

	// 设置cookie, 支持ajax
	c.SetSameSite(http.SameSiteNoneMode)
	// todo Host 白名单

	// cookie有效期 7天免费登录 选项
	cookieMaxAge := 0
	if loginForm.Remember {
		cookieMaxAge = 86400 * 7
	}
	c.SetCookie(middleware.UserAuthKey, tokenString, cookieMaxAge, "/", c.Request.Host, true, true)
	c.JSON(http.StatusOK, appresp.Reps(gin.H{"token": tokenString, "params": qValues, "info": userInfo}, nil))

}

func (hdl *authHandler) Logout(c *gin.Context) {
	// 设置cookie, 支持ajax
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(middleware.UserAuthKey, "", -1, "/", c.Request.Host, true, true)
	c.JSON(http.StatusOK, appresp.Reps("success", nil))
}

package api

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser/user_api"
	"donkey-ucenter/app/service/iuser/user_def"
	"donkey-ucenter/apperror"
	"donkey-ucenter/middleware"
	"donkey-ucenter/req-resp/appresp"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*  */

type userHdl struct{}

var (
	UserHdl = &userHdl{}
)

func (hdl *userHdl) Info(c *gin.Context) {
	userId := middleware.GetUserId(c)

	userData, err := user_api.ApiSrv.Info(userId)
	if err != nil {
		c.JSON(http.StatusOK, appresp.Err(err))
		return
	}

	c.JSON(http.StatusOK, appresp.Reps(userData, nil))
}

func (hdl *userHdl) UpdateInfo(c *gin.Context) {
	form := new(model.UserInfo)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	userId := middleware.GetUserId(c)
	if userId <= 0 {
		c.JSON(http.StatusOK, appresp.Err(apperror.ErrUserNotLogin))
		return
	}

	affectedRows, err := user_api.ApiSrv.UpdateInfo(userId, form)
	if err != nil {
		c.JSON(http.StatusOK, appresp.Err(err))
		return
	}
	c.JSON(http.StatusOK, appresp.Reps(affectedRows, nil))
}

func (hdl *userHdl) ChangePwd(c *gin.Context) {
	form := new(user_def.ChangePwdForm)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	userId := middleware.GetUserId(c)
	if userId <= 0 {
		c.JSON(http.StatusOK, appresp.Err(apperror.ErrUserNotLogin))
		return
	}

	res, err := user_api.ApiSrv.ChangePwd(userId, form)
	c.JSON(http.StatusOK, appresp.Reps(res, err))
}

func (hdl *userHdl) BindEmail(c *gin.Context) {
	form := new(user_def.BindEmailForm)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	userId := middleware.GetUserId(c)
	if userId <= 0 {
		c.JSON(http.StatusOK, appresp.Err(apperror.ErrUserNotLogin))
		return
	}

	res, err := user_api.ApiSrv.BindEmail(userId, form)
	c.JSON(http.StatusOK, appresp.Reps(res, err))
}

func (hdl *userHdl) BindEmailConfirm(c *gin.Context) {
	form := new(user_def.BindEmailConfirmForm)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	userId := middleware.GetUserId(c)
	if userId <= 0 {
		c.JSON(http.StatusOK, appresp.Err(apperror.ErrUserNotLogin))
		return
	}

	err := user_api.ApiSrv.BindEmailConfirm(userId, form)
	c.JSON(http.StatusOK, appresp.Reps(nil, err))
}

package api

import (
	"donkey-ucenter/app/service/iuser/user_api"
	"donkey-ucenter/app/service/iuser/user_def"
	"donkey-ucenter/req-resp/appresp"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*  */

type userPublicHdl struct{}

var (
	UserPublicHdl = &userPublicHdl{}
)

func (hdl *userPublicHdl) ResetPwdByEmail(c *gin.Context) {
	form := new(user_def.ResetPwdForm)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := user_api.PublicApiSrv.ResetPwdByEmail(form)
	c.JSON(http.StatusOK, appresp.Reps(nil, err))
}

func (hdl *userPublicHdl) ResetPwdConfirm(c *gin.Context) {
	form := new(user_def.ResetPwdConformForm)
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := user_api.PublicApiSrv.ResetPwdConform(form)
	c.JSON(http.StatusOK, appresp.Reps(nil, err))
}

package admin

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iverification"
	"donkey-ucenter/app/service/iverification/verification_admin"
	"donkey-ucenter/app/service/iverification/verification_def"
	"donkey-ucenter/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type userVerificationHdl struct{}

var (
	UserVerificationHdl = &userVerificationHdl{}
)

func (hdl *userVerificationHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *userVerificationHdl) Query(c *gin.Context) {
	form := new(verification_def.verificationQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := verification_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userVerificationHdl) GetList(c *gin.Context) {
	form := new(verification_def.verificationQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := verification_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userVerificationHdl) Get(c *gin.Context) {
	info := new(model.Verification)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := verification_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userVerificationHdl) Add(c *gin.Context) {
	info := new(model.Verification)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.CreateUid = middleware.GetAdminId(c)

	err := iverification.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userVerificationHdl) Update(c *gin.Context) {
	info := new(model.Verification)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.UpdateUid = middleware.GetAdminId(c)

	_, err := iverification.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userVerificationHdl) Delete(c *gin.Context) {
	info := new(model.Verification)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := iverification.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// 优先使用update
// SetInfo 弥补 int=0, string="" update 不更新问题
func (hdl *userVerificationHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := iverification.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}

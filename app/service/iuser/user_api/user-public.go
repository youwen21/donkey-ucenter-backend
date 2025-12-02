package user_api

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser"
	"donkey-ucenter/app/service/iuser/user_def"
	"donkey-ucenter/app/service/iverification"
	"donkey-ucenter/lib/libutils"
	"errors"
	"time"
)

/*  */

type publicApiSrv struct{}

var (
	PublicApiSrv = &publicApiSrv{}
)

func (apiSrv *publicApiSrv) ResetPwdByEmail(f *user_def.ResetPwdForm) error {
	if f.Email == "" {
		return errors.New("邮箱不能为空")
	}

	_, err := iuser.Srv.GetByUsername(f.Email)
	if err != nil {
		return errors.New("邮箱不可找回密码")
	}

	code := libutils.RandString(6)
	// 存储userId , 验证码, 绑定时间, 邮箱
	now := time.Now()
	expires := now.Add(time.Minute * 10)
	verifyData := new(model.Verification)
	verifyData.UserId = 0
	verifyData.Type = "password_reset"
	verifyData.Code = code
	verifyData.Target = f.Email
	verifyData.Status = 1
	verifyData.CreatedAt = &now
	verifyData.ExpiresAt = &expires
	err = iverification.Srv.Insert(verifyData)

	return err
}

func (apiSrv *publicApiSrv) ResetPwdConform(f *user_def.ResetPwdConformForm) error {
	if f.Email == "" {
		return errors.New("邮箱不能为空")
	}

	verifyData, err := iverification.Srv.GetBy(&model.Verification{
		Type:   "password_reset",
		Target: f.Email,
		Code:   f.Code,
		Status: 1,
	})
	if err != nil {
		return errors.New("邮箱或者验证码错误")
	}
	if verifyData.ExpiresAt.Before(time.Now()) {
		return errors.New("验证码已过期")
	}

	userData, err := iuser.Srv.GetByUsername(f.Email)
	if err != nil {
		return errors.New("邮箱用户不存在")
	}

	upDate := model.User{
		UserInfo: model.UserInfo{
			Id: userData.Id,
		},
		Password: libutils.EncryptWord(f.Password),
	}
	_, err = iuser.Srv.Update(&upDate)
	return err
}

package user_api

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser"
	"donkey-ucenter/app/service/iuser/user_def"
	"donkey-ucenter/app/service/iverification"
	"donkey-ucenter/conf"
	"donkey-ucenter/lib/libutils"
	"errors"
	"time"
)

/*  */

type apiSrv struct{}

var (
	ApiSrv = &apiSrv{}
)

func (apiSrv *apiSrv) Info(userId int) (*model.UserInfo, error) {
	user, err := iuser.Srv.Get(userId)
	if err != nil {
		return nil, err
	}
	return user.ToUserInfo(), nil
}

func (apiSrv *apiSrv) UpdateInfo(userId int, f *model.UserInfo) (int64, error) {
	user := f.ToUser()
	user.Id = userId
	affected, err := iuser.Srv.Update(user)
	return affected, err
}
func (apiSrv *apiSrv) BindEmail(userId int, f *user_def.BindEmailForm) (int, error) {
	if f.Email == "" {
		return 0, errors.New("邮箱不能为空")
	}
	// 验证邮箱
	if !libutils.IsEmail(f.Email) {
		return 0, errors.New("邮箱格式错误")
	}

	code := libutils.RandString(6)

	// 存储userId , 验证码, 绑定时间, 邮箱
	now := time.Now()
	expires := now.Add(time.Minute * 10)
	verifyData := new(model.Verification)
	verifyData.UserId = userId
	verifyData.Type = "email_bind"
	verifyData.Code = code
	verifyData.Target = f.Email
	verifyData.Status = 1
	verifyData.CreatedAt = &now
	verifyData.ExpiresAt = &expires
	iverification.Srv.Insert(verifyData)

	// 发送验证码
	emailClient := conf.Config.Smtp.GetClient()
	if err := emailClient.SendTo("绑定邮箱验证码", "验证码："+code, f.Email); err != nil {
		return 0, err
	}

	return verifyData.Id, nil
}

func (apiSrv *apiSrv) BindEmailConfirm(userId int, f *user_def.BindEmailConfirmForm) error {
	if f.Code == "" {
		return errors.New("验证码不能为空")
	}
	// 根据验证码，用户ID 取出绑定时间，邮箱
	verifyData, err := iverification.Srv.GetBy(&model.Verification{
		UserId: userId,
		Type:   "email_bind",
		Target: f.Email,
		Status: 1,
	})
	if err != nil {
		return err
	}

	// 验证码是否正确, 是否过期
	if verifyData.Code != f.Code {
		return errors.New("验证码错误")
	}
	if !verifyData.ExpiresAt.Before(time.Now()) {
		return errors.New("验证码已过期")
	}

	// 绑定邮箱
	user := model.User{}
	user.Id = userId
	user.Email = f.Email
	_, err = iuser.Srv.Update(&user)
	return err
}

func (apiSrv *apiSrv) ChangePwd(userId int, f *user_def.ChangePwdForm) (int64, error) {
	if f.NewPassword == "" {
		return 0, errors.New("新密码不能为空")
	}
	if f.OldPassword == "" {
		return 0, errors.New("旧密码不能为空")
	}
	user, err := iuser.Srv.Get(userId)
	if err != nil {
		return 0, err
	}
	if user.Password != libutils.EncryptWord(f.OldPassword) {
		return 0, errors.New("旧密码错误")
	}

	upDate := model.User{
		UserInfo: model.UserInfo{
			Id: user.Id,
		},
		Password: libutils.EncryptWord(f.NewPassword),
	}

	affected, err := iuser.Srv.Update(&upDate)
	return affected, err
}

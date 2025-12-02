package auth

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/auth/auth_def"
	"donkey-ucenter/app/service/iuser"
	"donkey-ucenter/lib/libutils"
	"donkey-ucenter/middleware"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Register(f *auth_def.Form) (string, *model.UserInfo, error) {
	if f.Username == "" {
		return "", nil, errors.New("用户名不能为空")
	}
	if f.Password == "" {
		return "", nil, errors.New("密码不能为空")
	}

	userData, err := iuser.Srv.GetByUsername(f.Username)
	if nil != userData {
		return "", nil, errors.New("用户名不可注册")
	}

	info := &model.User{
		UserInfo: model.UserInfo{
			Name:     f.Username,
			Nickname: "",
			Status:   1,
		},
	}
	info.Password = libutils.EncryptWord(f.Password)

	err = iuser.Srv.Insert(info)
	if err != nil {
		return "", nil, err
	}
	if info.Id <= 1 {
		return "", nil, errors.New("注册失败")
	}

	tokenString, err := libutils.Jwt.GenToken(middleware.UserJwtSecret, jwt.MapClaims{middleware.UserCtxKey: info.Id, "exp": time.Now().Unix() + 86400*7}) // 记住登录7天，或者7天不关闭网页，登录最长有效期7天
	if err != nil {
		return "", nil, err
	}
	return tokenString, &info.UserInfo, err
}

func (s *srv) Login(f *auth_def.Form) (string, *model.UserInfo, error) {
	if f.Username == "" {
		return "", nil, errors.New("用户名不能为空")
	}
	if f.Password == "" {
		return "", nil, errors.New("密码不能为空")
	}

	userData, err := iuser.Srv.GetByUsername(f.Username)

	if nil == userData || err != nil {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 密码校验
	genPwd := libutils.EncryptWord(f.Password)
	if userData.Password != genPwd {
		return "", nil, errors.New("用户名或密码错误2")
	}

	tokenString, err := libutils.Jwt.GenToken(middleware.UserJwtSecret, jwt.MapClaims{middleware.UserCtxKey: userData.Id, "exp": time.Now().Unix() + 86400*7}) // 记住登录7天，或者7天不关闭网页，登录最长有效期7天
	if err != nil {
		return "", nil, err
	}

	return tokenString, &userData.UserInfo, nil
}

func (s *srv) Logout(v *model.User) error {
	return nil
}

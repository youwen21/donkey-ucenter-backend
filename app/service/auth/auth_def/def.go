package auth_def

import (
	"donkey-ucenter/app/model"
	"net/url"
)

type Form struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Remember bool   `json:"remember" form:"remember"`
}

type LoginRes struct {
	Token    string          `json:"token" form:"token"`
	Params   url.Values      `json:"params" form:"params"`
	UserInfo *model.UserInfo `json:"userInfo" form:"userInfo"`
}

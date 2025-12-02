package user_def

type ResetPwdForm struct {
	//UEP string `json:"UEP" form:"UEP"` // 登陆名 或者 email 或者 手机号
	Email string `json:"email" form:"email"`
}

type ResetPwdConformForm struct {
	//UEP string `json:"UEP" form:"UEP"` // 登陆名 或者 email 或者 手机号
	Email string `json:"email" form:"email"`
	//Token    string `json:"token" form:"token"`
	Code     string `json:"code" form:"code"`
	Password string `json:"password" form:"password"`
}

type ChangePwdForm struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

type BindEmailForm struct {
	Email string `json:"email" form:"email"`
}
type BindEmailConfirmForm struct {
	Email string `json:"email" form:"email"`
	Code  string `json:"code" form:"code"`
}

package user_verification_def

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/lib/libdto"
)

/*  */

type UserVerificationQueryForm struct {
	model.UserVerification
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type UserVerificationQueryRes struct {
	Total int64                    `json:"total" form:"total"`
	List  []model.UserVerification `json:"list" form:"list"`
}

type UserVerificationExDTO struct {
	model.UserVerification

	// extend
}

type UserVerificationQueryResEx struct {
	Total int64                   `json:"total" form:"total"`
	List  []UserVerificationExDTO `json:"list" form:"list"`
}

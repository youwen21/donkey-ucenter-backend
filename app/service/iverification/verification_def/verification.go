package verification_def

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/lib/libdto"
)

/*  */

type VerificationQueryForm struct {
	model.Verification
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type VerificationQueryRes struct {
	Total int64                `json:"total" form:"total"`
	List  []model.Verification `json:"list" form:"list"`
}

type VerificationExDTO struct {
	model.Verification

	// extend
}

type VerificationQueryResEx struct {
	Total int64               `json:"total" form:"total"`
	List  []VerificationExDTO `json:"list" form:"list"`
}

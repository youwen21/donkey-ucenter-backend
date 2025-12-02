package user_def

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/lib/libdto"
)

/*  */

type UserQueryForm struct {
	model.User
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type UserQueryRes struct {
	Total int64        `json:"total" form:"total"`
	List  []model.User `json:"list" form:"list"`
}

type UserExDTO struct {
	model.User

	// extend
}

type UserQueryResEx struct {
	Total int64       `json:"total" form:"total"`
	List  []UserExDTO `json:"list" form:"list"`
}

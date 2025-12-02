package user_verification_admin

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser_verification"
	"donkey-ucenter/app/service/iuser_verification/user_verification_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *user_verification_def.UserVerificationQueryForm) (*user_verification_def.UserVerificationQueryResEx, error) {
	res, err := iuser_verification.Srv.Query(f)
	// biz process
	result := new(user_verification_def.UserVerificationQueryResEx)
	result.Total = res.Total
	result.List = make([]user_verification_def.UserVerificationExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.UserVerification) user_verification_def.UserVerificationExDTO {
	return user_verification_def.UserVerificationExDTO{
		// TODO
		UserVerification: v,
	}
}

func (aSrv *adminSrv) GetList(f *user_verification_def.UserVerificationQueryForm) ([]model.UserVerification, error) {
	res, err := iuser_verification.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.UserVerification, error) {
	res, err := iuser_verification.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.UserVerification, error) {
	res, err := iuser_verification.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.UserVerification, error) {
	res, err := iuser_verification.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.UserVerification) (*model.UserVerification, error) {
	err := iuser_verification.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.UserVerification) (int64, error) {
	affected, err := iuser_verification.Srv.Update(v)
	return affected, err
}

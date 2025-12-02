package verification_admin

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iverification"
	"donkey-ucenter/app/service/iverification/verification_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *verification_def.VerificationQueryForm) (*verification_def.VerificationQueryResEx, error) {
	res, err := iverification.Srv.Query(f)
	// biz process
	result := new(verification_def.VerificationQueryResEx)
	result.Total = res.Total
	result.List = make([]verification_def.VerificationExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.Verification) verification_def.VerificationExDTO {
	return verification_def.VerificationExDTO{
		// TODO
		Verification: v,
	}
}

func (aSrv *adminSrv) GetList(f *verification_def.VerificationQueryForm) ([]model.Verification, error) {
	res, err := iverification.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.Verification, error) {
	res, err := iverification.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.Verification, error) {
	res, err := iverification.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.Verification, error) {
	res, err := iverification.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.Verification) (*model.Verification, error) {
	err := iverification.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.Verification) (int64, error) {
	affected, err := iverification.Srv.Update(v)
	return affected, err
}

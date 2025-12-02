package iuser_verification

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser_verification/internal/dml"
	"donkey-ucenter/app/service/iuser_verification/user_verification_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *user_verification_def.UserVerificationQueryForm) (int64, error) {
	return dml.UserVerificationDml.Count(f)
}

func (s *srv) Query(f *user_verification_def.UserVerificationQueryForm) (*user_verification_def.UserVerificationQueryRes, error) {
	res, err := dml.UserVerificationDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *user_verification_def.UserVerificationQueryForm) ([]model.UserVerification, error) {
	res, err := dml.UserVerificationDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.UserVerification, error) {
	return dml.UserVerificationDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.UserVerification, error) {
	res, err := dml.UserVerificationDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.UserVerification, error) {
	res, err := dml.UserVerificationDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.UserVerification) (*model.UserVerification, error) {
	res, err := dml.UserVerificationDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.UserVerification, error) {
	res, err := dml.UserVerificationDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.UserVerification) error {
	err := dml.UserVerificationDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.UserVerification, batchSize int) (int64, error) {
	return dml.UserVerificationDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.UserVerification) (int64, error) {
	return dml.UserVerificationDml.Update(m)
}

func (s *srv) UpdateBy(f *model.UserVerification, data map[string]any) (int64, error) {
	return dml.UserVerificationDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.UserVerificationDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.UserVerificationDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) (int64, error) {
	return dml.UserVerificationDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.UserVerification, error) {
	return dml.UserVerificationDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.UserVerification, error) {
	return dml.UserVerificationDml.RawFind(sql)
}

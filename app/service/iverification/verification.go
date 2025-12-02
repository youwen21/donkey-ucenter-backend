package iverification

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iverification/internal/dml"
	"donkey-ucenter/app/service/iverification/verification_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *verification_def.VerificationQueryForm) (int64, error) {
	return dml.VerificationDml.Count(f)
}

func (s *srv) Query(f *verification_def.VerificationQueryForm) (*verification_def.VerificationQueryRes, error) {
	res, err := dml.VerificationDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *verification_def.VerificationQueryForm) ([]model.Verification, error) {
	res, err := dml.VerificationDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.Verification, error) {
	return dml.VerificationDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.Verification, error) {
	res, err := dml.VerificationDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.Verification, error) {
	res, err := dml.VerificationDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.Verification) (*model.Verification, error) {
	res, err := dml.VerificationDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.Verification, error) {
	res, err := dml.VerificationDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.Verification) error {
	err := dml.VerificationDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.Verification, batchSize int) (int64, error) {
	return dml.VerificationDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.Verification) (int64, error) {
	return dml.VerificationDml.Update(m)
}

func (s *srv) UpdateBy(f *model.Verification, data map[string]any) (int64, error) {
	return dml.VerificationDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.VerificationDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.VerificationDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) (int64, error) {
	return dml.VerificationDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.Verification, error) {
	return dml.VerificationDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.Verification, error) {
	return dml.VerificationDml.RawFind(sql)
}

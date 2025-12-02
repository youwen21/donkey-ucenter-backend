package iuser

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser/internal/dml"
	"donkey-ucenter/app/service/iuser/user_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *user_def.UserQueryForm) (int64, error) {
	return dml.UserDml.Count(f)
}

func (s *srv) Query(f *user_def.UserQueryForm) (*user_def.UserQueryRes, error) {
	res, err := dml.UserDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *user_def.UserQueryForm) ([]model.User, error) {
	res, err := dml.UserDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.User, error) {
	return dml.UserDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.User, error) {
	res, err := dml.UserDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.User, error) {
	res, err := dml.UserDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.User) (*model.User, error) {
	res, err := dml.UserDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.User, error) {
	res, err := dml.UserDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.User) error {
	err := dml.UserDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.User, batchSize int) (int64, error) {
	return dml.UserDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.User) (int64, error) {
	return dml.UserDml.Update(m)
}

func (s *srv) UpdateBy(f *model.User, data map[string]any) (int64, error) {
	return dml.UserDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.UserDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.UserDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) (int64, error) {
	return dml.UserDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.User, error) {
	return dml.UserDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.User, error) {
	return dml.UserDml.RawFind(sql)
}
func (s *srv) GetByUsername(userName string) (*model.User, error) {
	form := new(model.User)
	form.Name = userName
	return s.GetBy(form)
}

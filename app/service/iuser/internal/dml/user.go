package dml

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser/internal/dml/internal/dal"
	"donkey-ucenter/app/service/iuser/user_def"
	"donkey-ucenter/lib/lru"
	"fmt"
)

/*  */

const (
	userDmlLruKey = "lru:userDml:"
)

type userDml struct{}

var (
	UserDml            = &userDml{}
	userDmlLruCache, _ = lru.New(200)
)

func (dm *userDml) Count(f *user_def.UserQueryForm) (int64, error) {
	return dal.UserDal.Count(f)
}

func (dm *userDml) Query(f *user_def.UserQueryForm) (*user_def.UserQueryRes, error) {
	return dal.UserDal.Query(f)
}

func (dm *userDml) GetList(f *user_def.UserQueryForm) ([]model.User, error) {
	return dal.UserDal.GetList(f)
}

func (dm *userDml) GetLisByPkList(pkList []int) ([]model.User, error) {
	return dal.UserDal.GetLisByPkList(pkList)
}

func (dm *userDml) GetAll() ([]model.User, error) {
	return dal.UserDal.GetAll()
}

func (dm *userDml) Get(pk int) (*model.User, error) {
	return dal.UserDal.Get(pk)
}

func (dm *userDml) GetBy(f *model.User) (*model.User, error) {
	return dal.UserDal.GetBy(f)
}

func (dm *userDml) GetMulti(pkList []int) (map[int]model.User, error) {
	return dal.UserDal.GetMulti(pkList)
}

func (dm *userDml) Insert(m *model.User) error {
	return dal.UserDal.Insert(m)
}

func (dm *userDml) BatchInsert(bm []*model.User, batchSize int) (int64, error) {
	return dal.UserDal.BatchInsert(bm, batchSize)
}

func (dm *userDml) Update(m *model.User) (int64, error) {
	return dal.UserDal.Update(m)
}

func (dm *userDml) UpdateBy(f *model.User, data map[string]any) (int64, error) {
	return dal.UserDal.UpdateBy(f, data)
}

func (dm *userDml) SetInfo(data map[string]any) (int64, error) {
	return dal.UserDal.SetInfo(data)
}

func (dm *userDml) Delete(pk int) error {
	return dal.UserDal.Delete(pk)
}

func (dm *userDml) Exec(sql string, values ...interface{}) (int64, error) {
	return dal.UserDal.Exec(sql, values...)
}

func (dm *userDml) RawGet(sql string) (*model.User, error) {
	return dal.UserDal.RawGet(sql)
}

func (dm *userDml) RawFind(sql string) ([]model.User, error) {
	return dal.UserDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *userDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", userDmlLruKey, key)
}

func (dm *userDml) LruRemove(key string) bool {
	return userDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *userDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */

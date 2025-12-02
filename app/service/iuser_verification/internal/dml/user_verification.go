package dml

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser_verification/internal/dml/internal/dal"
	"donkey-ucenter/app/service/iuser_verification/user_verification_def"
	"donkey-ucenter/lib/lru"
	"fmt"
)

/*  */

const (
	userVerificationDmlLruKey = "lru:userVerificationDml:"
)

type userVerificationDml struct{}

var (
	UserVerificationDml            = &userVerificationDml{}
	userVerificationDmlLruCache, _ = lru.New(200)
)

func (dm *userVerificationDml) Count(f *user_verification_def.UserVerificationQueryForm) (int64, error) {
	return dal.UserVerificationDal.Count(f)
}

func (dm *userVerificationDml) Query(f *user_verification_def.UserVerificationQueryForm) (*user_verification_def.UserVerificationQueryRes, error) {
	return dal.UserVerificationDal.Query(f)
}

func (dm *userVerificationDml) GetList(f *user_verification_def.UserVerificationQueryForm) ([]model.UserVerification, error) {
	return dal.UserVerificationDal.GetList(f)
}

func (dm *userVerificationDml) GetLisByPkList(pkList []int) ([]model.UserVerification, error) {
	return dal.UserVerificationDal.GetLisByPkList(pkList)
}

func (dm *userVerificationDml) GetAll() ([]model.UserVerification, error) {
	return dal.UserVerificationDal.GetAll()
}

func (dm *userVerificationDml) Get(pk int) (*model.UserVerification, error) {
	return dal.UserVerificationDal.Get(pk)
}

func (dm *userVerificationDml) GetBy(f *model.UserVerification) (*model.UserVerification, error) {
	return dal.UserVerificationDal.GetBy(f)
}

func (dm *userVerificationDml) GetMulti(pkList []int) (map[int]model.UserVerification, error) {
	return dal.UserVerificationDal.GetMulti(pkList)
}

func (dm *userVerificationDml) Insert(m *model.UserVerification) error {
	return dal.UserVerificationDal.Insert(m)
}

func (dm *userVerificationDml) BatchInsert(bm []*model.UserVerification, batchSize int) (int64, error) {
	return dal.UserVerificationDal.BatchInsert(bm, batchSize)
}

func (dm *userVerificationDml) Update(m *model.UserVerification) (int64, error) {
	return dal.UserVerificationDal.Update(m)
}

func (dm *userVerificationDml) UpdateBy(f *model.UserVerification, data map[string]any) (int64, error) {
	return dal.UserVerificationDal.UpdateBy(f, data)
}

func (dm *userVerificationDml) SetInfo(data map[string]any) (int64, error) {
	return dal.UserVerificationDal.SetInfo(data)
}

func (dm *userVerificationDml) Delete(pk int) error {
	return dal.UserVerificationDal.Delete(pk)
}

func (dm *userVerificationDml) Exec(sql string, values ...interface{}) (int64, error) {
	return dal.UserVerificationDal.Exec(sql, values...)
}

func (dm *userVerificationDml) RawGet(sql string) (*model.UserVerification, error) {
	return dal.UserVerificationDal.RawGet(sql)
}

func (dm *userVerificationDml) RawFind(sql string) ([]model.UserVerification, error) {
	return dal.UserVerificationDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *userVerificationDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", userVerificationDmlLruKey, key)
}

func (dm *userVerificationDml) LruRemove(key string) bool {
	return userVerificationDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *userVerificationDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */

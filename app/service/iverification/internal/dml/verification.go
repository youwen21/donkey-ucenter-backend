package dml

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iverification/internal/dml/internal/dal"
	"donkey-ucenter/app/service/iverification/verification_def"
	"donkey-ucenter/lib/lru"
	"fmt"
)

/*  */

const (
	verificationDmlLruKey = "lru:verificationDml:"
)

type verificationDml struct{}

var (
	VerificationDml            = &verificationDml{}
	verificationDmlLruCache, _ = lru.New(200)
)

func (dm *verificationDml) Count(f *verification_def.VerificationQueryForm) (int64, error) {
	return dal.VerificationDal.Count(f)
}

func (dm *verificationDml) Query(f *verification_def.VerificationQueryForm) (*verification_def.VerificationQueryRes, error) {
	return dal.VerificationDal.Query(f)
}

func (dm *verificationDml) GetList(f *verification_def.VerificationQueryForm) ([]model.Verification, error) {
	return dal.VerificationDal.GetList(f)
}

func (dm *verificationDml) GetLisByPkList(pkList []int) ([]model.Verification, error) {
	return dal.VerificationDal.GetLisByPkList(pkList)
}

func (dm *verificationDml) GetAll() ([]model.Verification, error) {
	return dal.VerificationDal.GetAll()
}

func (dm *verificationDml) Get(pk int) (*model.Verification, error) {
	return dal.VerificationDal.Get(pk)
}

func (dm *verificationDml) GetBy(f *model.Verification) (*model.Verification, error) {
	return dal.VerificationDal.GetBy(f)
}

func (dm *verificationDml) GetMulti(pkList []int) (map[int]model.Verification, error) {
	return dal.VerificationDal.GetMulti(pkList)
}

func (dm *verificationDml) Insert(m *model.Verification) error {
	return dal.VerificationDal.Insert(m)
}

func (dm *verificationDml) BatchInsert(bm []*model.Verification, batchSize int) (int64, error) {
	return dal.VerificationDal.BatchInsert(bm, batchSize)
}

func (dm *verificationDml) Update(m *model.Verification) (int64, error) {
	return dal.VerificationDal.Update(m)
}

func (dm *verificationDml) UpdateBy(f *model.Verification, data map[string]any) (int64, error) {
	return dal.VerificationDal.UpdateBy(f, data)
}

func (dm *verificationDml) SetInfo(data map[string]any) (int64, error) {
	return dal.VerificationDal.SetInfo(data)
}

func (dm *verificationDml) Delete(pk int) error {
	return dal.VerificationDal.Delete(pk)
}

func (dm *verificationDml) Exec(sql string, values ...interface{}) (int64, error) {
	return dal.VerificationDal.Exec(sql, values...)
}

func (dm *verificationDml) RawGet(sql string) (*model.Verification, error) {
	return dal.VerificationDal.RawGet(sql)
}

func (dm *verificationDml) RawFind(sql string) ([]model.Verification, error) {
	return dal.VerificationDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *verificationDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", verificationDmlLruKey, key)
}

func (dm *verificationDml) LruRemove(key string) bool {
	return verificationDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *verificationDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */

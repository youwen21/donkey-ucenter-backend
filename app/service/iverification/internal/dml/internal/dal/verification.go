package dal

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iverification/verification_def"
	"donkey-ucenter/apperror"
	"donkey-ucenter/conf"
	"errors"

	"gorm.io/gorm"
)

/*  */

type verificationDal struct{}

var (
	VerificationDal = &verificationDal{}
)

func (d *verificationDal) GetSessionByModel(m *model.Verification) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}
	if m.UserId != 0 {
		session.Where("user_id = ?", m.UserId)
	}
	if m.Code != "" {
		session.Where("code = ?", m.Code)
	}
	if m.Type != "" {
		session.Where("type = ?", m.Type)
	}
	if m.Target != "" {
		session.Where("target = ?", m.Target)
	}
	if m.Status != 0 {
		session.Where("status = ?", m.Status)
	}
	if m.ExpiresAt != nil {
		session.Where("expires_at = ?", m.ExpiresAt)
	}
	if m.CreatedAt != nil {
		session.Where("created_at = ?", m.CreatedAt)
	}

	return session
}

func (d *verificationDal) GetSessionByForm(f *verification_def.VerificationQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.Verification)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *verificationDal) Count(f *verification_def.VerificationQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *verificationDal) Query(f *verification_def.VerificationQueryForm) (*verification_def.VerificationQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.Verification

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &verification_def.VerificationQueryRes{Total: total, List: list}, nil
}

func (d *verificationDal) GetList(f *verification_def.VerificationQueryForm) ([]model.Verification, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.Verification

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *verificationDal) GetAll() ([]model.Verification, error) {
	var results []model.Verification

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *verificationDal) Get(pk int) (*model.Verification, error) {
	info := &model.Verification{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *verificationDal) GetBy(m *model.Verification) (*model.Verification, error) {
	session := d.GetSessionByModel(m)

	info := &model.Verification{}

	session.Order("`id` DESC")

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *verificationDal) GetLisByPkList(pkList []int) ([]model.Verification, error) {
	var results []model.Verification

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *verificationDal) GetMulti(pkList []int) (map[int]model.Verification, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.Verification)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *verificationDal) Insert(m *model.Verification) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *verificationDal) BatchInsert(bm []*model.Verification, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *verificationDal) Update(m *model.Verification) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *verificationDal) UpdateBy(f *model.Verification, data map[string]any) (int64, error) {
	// where clause
	session := d.GetSessionByModel(f)

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *verificationDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk id")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *verificationDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.Verification{}).Error
	return err
}

func (d *verificationDal) Exec(sql string, values ...interface{}) (int64, error) {
	session := d.newSession()
	session.Exec(sql, values...)
	return session.RowsAffected, session.Error
}

func (d *verificationDal) RawGet(sql string) (*model.Verification, error) {
	info := &model.Verification{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *verificationDal) RawFind(sql string) ([]model.Verification, error) {
	var results []model.Verification
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *verificationDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *verificationDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_user_verification")
}

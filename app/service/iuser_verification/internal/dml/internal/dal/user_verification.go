package dal

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser_verification/user_verification_def"
	"donkey-ucenter/apperror"
	"donkey-ucenter/conf"
	"errors"

	"gorm.io/gorm"
)

/*  */

type userVerificationDal struct{}

var (
	UserVerificationDal = &userVerificationDal{}
)

func (d *userVerificationDal) GetSessionByModel(m *model.UserVerification) *gorm.DB {
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

func (d *userVerificationDal) GetSessionByForm(f *user_verification_def.UserVerificationQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.UserVerification)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *userVerificationDal) Count(f *user_verification_def.UserVerificationQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *userVerificationDal) Query(f *user_verification_def.UserVerificationQueryForm) (*user_verification_def.UserVerificationQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.UserVerification

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &user_verification_def.UserVerificationQueryRes{Total: total, List: list}, nil
}

func (d *userVerificationDal) GetList(f *user_verification_def.UserVerificationQueryForm) ([]model.UserVerification, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.UserVerification

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userVerificationDal) GetAll() ([]model.UserVerification, error) {
	var results []model.UserVerification

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userVerificationDal) Get(pk int) (*model.UserVerification, error) {
	info := &model.UserVerification{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userVerificationDal) GetBy(m *model.UserVerification) (*model.UserVerification, error) {
	session := d.GetSessionByModel(m)

	info := &model.UserVerification{}

	session.Order("`id` DESC")

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userVerificationDal) GetLisByPkList(pkList []int) ([]model.UserVerification, error) {
	var results []model.UserVerification

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *userVerificationDal) GetMulti(pkList []int) (map[int]model.UserVerification, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.UserVerification)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *userVerificationDal) Insert(m *model.UserVerification) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *userVerificationDal) BatchInsert(bm []*model.UserVerification, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *userVerificationDal) Update(m *model.UserVerification) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *userVerificationDal) UpdateBy(f *model.UserVerification, data map[string]any) (int64, error) {
	// where clause
	session := d.GetSessionByModel(f)

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *userVerificationDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk id")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *userVerificationDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.UserVerification{}).Error
	return err
}

func (d *userVerificationDal) Exec(sql string, values ...interface{}) (int64, error) {
	session := d.newSession()
	session.Exec(sql, values...)
	return session.RowsAffected, session.Error
}

func (d *userVerificationDal) RawGet(sql string) (*model.UserVerification, error) {
	info := &model.UserVerification{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userVerificationDal) RawFind(sql string) ([]model.UserVerification, error) {
	var results []model.UserVerification
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *userVerificationDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *userVerificationDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_user_verification")
}

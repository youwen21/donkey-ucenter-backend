package dal

import (
	"donkey-ucenter/app/model"
	"donkey-ucenter/app/service/iuser/user_def"
	"donkey-ucenter/apperror"
	"donkey-ucenter/conf"
	"errors"

	"gorm.io/gorm"
)

/*  */

type userDal struct{}

var (
	UserDal = &userDal{}
)

func (d *userDal) GetSessionByModel(m *model.User) *gorm.DB {
	session := d.newSession()

	if m.Id != 0 {
		session.Where("id = ?", m.Id)
	}

	return session
}

func (d *userDal) GetSessionByForm(f *user_def.UserQueryForm) *gorm.DB {
	session := d.GetSessionByModel(&f.User)

	if len(f.IdList) > 0 {
		session.Where("id in (?)", f.IdList)
	}

	return session
}

func (d *userDal) Count(f *user_def.UserQueryForm) (int64, error) {
	session := d.GetSessionByForm(f)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *userDal) Query(f *user_def.UserQueryForm) (*user_def.UserQueryRes, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var total int64
	var list []model.User

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &user_def.UserQueryRes{Total: total, List: list}, nil
}

func (d *userDal) GetList(f *user_def.UserQueryForm) ([]model.User, error) {
	session := d.GetSessionByForm(f)

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	var results []model.User

	err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userDal) GetAll() ([]model.User, error) {
	var results []model.User

	session := d.newSession()
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *userDal) Get(pk int) (*model.User, error) {
	info := &model.User{}
	session := d.newSession()
	if err := session.Where("`id`= ?", pk).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userDal) GetBy(m *model.User) (*model.User, error) {
	session := d.GetSessionByModel(m)

	info := &model.User{}

	if err := session.First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userDal) GetLisByPkList(pkList []int) ([]model.User, error) {
	var results []model.User

	session := d.newSession()
	query := session.Where("`id` IN ?", pkList)
	err := query.Find(&results).Error
	return results, err
}

func (d *userDal) GetMulti(pkList []int) (map[int]model.User, error) {
	if len(pkList) == 0 {
		return nil, apperror.PkListEmpty
	}

	var mMap = make(map[int]model.User)

	results, err := d.GetLisByPkList(pkList)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		mMap[v.Id] = v
	}
	return mMap, nil
}

func (d *userDal) Insert(m *model.User) error {
	session := d.newSession()
	err := session.Create(m).Error
	return err
}

func (d *userDal) BatchInsert(bm []*model.User, batchSize int) (int64, error) {
	session := d.newSession()
	err := session.CreateInBatches(bm, batchSize).Error
	return session.RowsAffected, err
}

func (d *userDal) Update(m *model.User) (int64, error) {
	session := d.newSession()
	err := session.Updates(m).Error
	return session.RowsAffected, err
}

func (d *userDal) UpdateBy(f *model.User, data map[string]any) (int64, error) {
	// where clause
	session := d.GetSessionByModel(f)

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

// SetInfo 允许 0 和 "" 值，优先使用 Update
func (d *userDal) SetInfo(data map[string]any) (int64, error) {
	session := d.newSession()

	if _, ok := data["id"]; !ok {
		return 0, errors.New("lost pk id")
	}

	// where clause
	session.Where("`id` = ?", data["id"]).Omit("id")

	err := session.Updates(data).Error
	return session.RowsAffected, err
}

func (d *userDal) Delete(pk int) error {
	session := d.newSession()
	err := session.Where("`id` = ?", pk).Delete(model.User{}).Error
	return err
}

func (d *userDal) Exec(sql string, values ...interface{}) (int64, error) {
	session := d.newSession()
	session.Exec(sql, values...)
	return session.RowsAffected, session.Error
}

func (d *userDal) RawGet(sql string) (*model.User, error) {
	info := &model.User{}
	session := d.newSession()
	if err := session.Raw(sql).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *userDal) RawFind(sql string) ([]model.User, error) {
	var results []model.User
	session := d.newSession()
	err := session.Raw(sql).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *userDal) newEngine() *gorm.DB {
	return conf.Config.MysqlDefault.GetDb()
}

func (d *userDal) newSession() *gorm.DB {
	return conf.Config.MysqlDefault.GetSession().Table("t_user")
}

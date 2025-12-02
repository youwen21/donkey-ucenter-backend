package model

import "time"

/*  */

type UserInfo struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`
	Name       string     `json:"name" form:"name"`         // 登陆名
	Nickname   string     `json:"nickname" form:"nickname"` // 称呼
	Avatar     string     `json:"avatar" form:"avatar"`     // 头像
	Email      string     `json:"email" form:"email"`       // 邮箱
	Phone      string     `json:"phone" form:"phone"`       // 手机号
	Status     int        `json:"status" form:"status"`     // 状态
	CreateTime *time.Time `json:"create_time" form:"create_time"`
}

func (m *UserInfo) ToUser() *User {
	var user User
	user.Id = m.Id
	user.Name = m.Name
	user.Nickname = m.Nickname
	user.Avatar = m.Avatar
	user.Email = m.Email
	user.Phone = m.Phone
	user.Status = m.Status
	user.CreateTime = m.CreateTime

	return &user
}

type User struct {
	UserInfo

	Password   string     `json:"password" form:"password"`                                            // 密码
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *User) TableName() string {
	return "t_user"
}

func (m *User) ToUserInfo() *UserInfo {
	var userInfo UserInfo
	userInfo.Id = m.Id
	userInfo.Name = m.Name
	userInfo.Nickname = m.Nickname
	userInfo.Avatar = m.Avatar
	userInfo.Email = m.Email
	userInfo.Phone = m.Phone
	userInfo.Status = m.Status
	userInfo.CreateTime = m.CreateTime

	return &userInfo
}

package model

import "time"

/*  */

type Verification struct {
	Id        int        `json:"id" form:"id" gorm:"autoIncrement"`                  // 主键ID
	UserId    int        `json:"user_id" form:"user_id"`                             // 用户ID，0表示未注册用户（注册验证场景）
	Type      string     `json:"type" form:"type"`                                   // 验证类型: email_register, phone_register, email_bind, phone_bind, email_reset, phone_reset 等
	Code      string     `json:"code" form:"code"`                                   // 验证码或令牌
	Target    string     `json:"target" form:"target"`                               // 目标值: 邮箱地址、手机号等
	Status    int8       `json:"status" form:"status"`                               // 状态: 0-未使用, 1-已使用, 2-已过期
	ExpiresAt *time.Time `json:"expires_at" form:"expires_at"`                       // 过期时间
	CreatedAt *time.Time `json:"created_at" form:"created_at" gorm:"autoCreateTime"` // 创建时间
	UsedAt    *time.Time `json:"used_at" form:"used_at"`                             // 使用时间
}

// 自定义表名
func (m *Verification) TableName() string {
	return "t_verification"
}

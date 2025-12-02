package libdto

type OrderForm struct {
	OrderBy     []string `form:"orderBy" json:"orderBy" `        // 字段
	OrderDirect string   `form:"orderDirect" json:"orderDirect"` // 方式
}

type PageForm struct {
	Page     int `form:"page,default=1" json:"page,default=1"`
	PageSize int `form:"pageSize,default=20" json:"pageSize,default=20"`
}

// Limit 获取每页记录数
func (p PageForm) Limit() int {
	if p.PageSize == 0 {
		return 20
	}
	return p.PageSize
}

// Offset 获取当前页开始ID
func (p PageForm) Offset() int {
	pSize := p.PageSize
	if pSize == 0 {
		pSize = 20
	}

	if p.Page == 0 {
		return 0
	}

	return (p.Page - 1) * pSize
}

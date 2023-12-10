package dto

// 通用user ID DTO
type UserIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// page DTO
type Pagination struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (p Pagination) GetPage() int {
	if p.Page < 1 {
		return 1
	}
	return p.Page
}

func (p Pagination) GetLimit() int {
	if p.Limit < 1 {
		return 1
	}
	return p.Limit
}

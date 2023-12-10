package dao

import (
	"gogofly/service/dto"

	"gorm.io/gorm"
)

// 通用分页函数
func Paginate(p dto.Pagination) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetPage()).Limit(int(p.GetLimit()))
	}
}

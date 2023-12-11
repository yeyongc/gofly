package model

// entity: 对应数据库sys_user_address表
type UserAddress struct {
	ID            uint   `gorm:"primarykey" json:"address_id" form:"address_id"`
	Mobile        string `gorm:"size:11,not null" json:"mobile" form:"mobile"`
	UserName      string `gorm:"size:64" json:"user_name" form:"user_name"`
	Country       string `gorm:"size:128,default:中国" json:"country" from:"country"`
	Province      string `gorm:"size:128" json:"province" form:"province"`
	City          string `gorm:"size:128" json:"city" form:"city"`
	County        string `gorm:"size:128" json:"county" form:"county"`
	DetailAddress string `gorm:"size:128" json:"detail_address" form:"detail_address"`
	Mark          string `gorm:"size:10" json:"mark" form:"mark"`
	UserID        uint   `gorm:"not null" json:"owner_id" form:"owner_id" uri:"id"`
}

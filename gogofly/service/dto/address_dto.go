package dto

import "gogofly/model"

type AddressAddDTO struct {
	Mobile        string `gorm:"size:11" json:"mobile" form:"address"`
	UserName      string `gorm:"size:64,not null" json:"user_name" form:"user_name"`
	Country       string `gorm:"size:128" json:"country" from:"country"`
	Province      string `gorm:"size:128" json:"province" form:"province"`
	City          string `gorm:"size:128" json:"city" form:"city"`
	County        string `gorm:"size:128" json:"county" form:"county"`
	DetailAddress string `gorm:"size:128" json:"detail_address" form:"detail_address"`
	Mark          string `gorm:"size:10,default null" json:"mark" form:"mark"`
	UserID        uint   `gorm:"not null" json:"user_id" uri:"id" form:"user_id" binding:"required"`
}

type AddressUpdateDTO struct {
	AddressAddDTO
	ID uint `json:"address_id" form:"address_id" binding:"required"`
}

// 通用address ID DTO
type AddressDeleteDTO struct {
	IDs    []uint `json:"ids" form:"ids"`
	UserID int    `json:"user_id" form:"user_id" uri:"id"`
}

func (m *AddressAddDTO) ConvertToModel(userAddress *model.UserAddress) {
	userAddress.City = m.City
	userAddress.Country = m.Country
	userAddress.County = m.County
	userAddress.DetailAddress = m.DetailAddress
	userAddress.Mark = m.Mark
	userAddress.Mobile = m.Mobile
	userAddress.Province = m.Province
	userAddress.UserName = m.UserName
	userAddress.UserID = m.UserID
}

func (m *AddressUpdateDTO) ConvertToModel(userAddress *model.UserAddress) {
	userAddress.City = m.City
	userAddress.Country = m.Country
	userAddress.County = m.County
	userAddress.DetailAddress = m.DetailAddress
	userAddress.Mark = m.Mark
	userAddress.Mobile = m.Mobile
	userAddress.Province = m.Province
	userAddress.UserName = m.UserName
	userAddress.ID = m.ID
	userAddress.UserID = m.UserID
}

package dao

import (
	"fmt"
	"gogofly/model"
	"gogofly/service/dto"
)

type AddressDao struct {
	BaseDao
}

func (m *AddressDao) AddAddress(dto *dto.AddressAddDTO) error {
	var err error
	var userAddress model.UserAddress

	dto.ConvertToModel(&userAddress)

	err = m.Orm.Save(&userAddress).Error

	return err
}
func (m *AddressDao) UpdateAddress(dto *dto.AddressUpdateDTO) error {
	var err error
	var userAddress model.UserAddress
	dto.ConvertToModel(&userAddress)
	err = m.Orm.Save(&userAddress).Where(userAddress.ID).Error
	return err
}

func (m *AddressDao) GetAll(ownerId uint) ([]model.UserAddress, error) {
	var (
		userAddressLsit []model.UserAddress
		err             error
	)
	fmt.Println(ownerId)
	m.Orm.Model(&model.UserAddress{}).Where("user_id = ?", ownerId).Find(&userAddressLsit)
	return userAddressLsit, err
}

func (m *AddressDao) Detete(DTO *dto.AddressDeleteDTO) error {
	var (
		err       error
		addresses []model.UserAddress
	)
	for _, id := range DTO.IDs {
		addresses = append(addresses, model.UserAddress{ID: id})
	}
	err = m.Orm.Delete(&addresses, DTO.IDs).Error
	return err
}

var addressDao *AddressDao

func NewAddressDao() *AddressDao {
	if addressDao == nil {
		addressDao = &AddressDao{
			BaseDao: *NewBaseDao(),
		}
	}
	return addressDao
}

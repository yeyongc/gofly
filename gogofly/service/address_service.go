package service

import (
	"gogofly/dao"
	"gogofly/model"
	"gogofly/service/dto"
)

type AddressService struct {
	UserService
	Dao *dao.AddressDao
}

func (m *AddressService) AddAddress(addressAddDTO *dto.AddressAddDTO) error {
	return m.Dao.AddAddress(addressAddDTO)
}
func (m *AddressService) UpdateAddress(addressUpdateDTO *dto.AddressUpdateDTO) error {
	return m.Dao.UpdateAddress(addressUpdateDTO)
}
func (m *AddressService) GetAll(ownerId uint) ([]model.UserAddress, error) {
	return m.Dao.GetAll(ownerId)
}
func (m *AddressService) Delete(addressDeleteDTO *dto.AddressDeleteDTO) error {
	return m.Dao.Detete(addressDeleteDTO)
}

var addressService *AddressService

func NewAddressService(Dao *dao.AddressDao) *AddressService {
	if addressService == nil {
		addressService = &AddressService{
			Dao: Dao,
		}
	}
	return addressService
}

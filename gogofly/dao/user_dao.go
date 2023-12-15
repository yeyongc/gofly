package dao

import (
	"gogofly/model"
	"gogofly/service/dto"
)

type UserDao struct {
	BaseDao
}

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		return &UserDao{
			BaseDao: *NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(username, password string) model.User {
	var user model.User
	m.Orm.Model(&user).Where("name=? and password=?", username, password).Find(&user)
	return user
}

func (m *UserDao) AddUser(userDTO *dto.UserAddDTO) error {
	var user model.User
	userDTO.ConvertToModel(&user)

	err := m.Orm.Save(&user).Error

	if err == nil {
		userDTO.ID = user.ID
		userDTO.Password = ""
	}

	return err
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := m.Orm.First(&user, id).Error
	return user, err
}

func (m *UserDao) GetUserList(userListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var userList []model.User
	var userNum int64

	err := m.Orm.Model(&model.User{}).Scopes(Paginate(userListDTO.Pagination)).Find(&userList).Offset(-1).Limit(-1).Count(&userNum).Error

	return userList, userNum, err

}
func (m *UserDao) UpdateUser(userUpdateDTO *dto.UserUpdatetDTO) error {
	var user model.User
	m.Orm.First(&user, userUpdateDTO.ID)
	userUpdateDTO.ConvertToModel(&user)
	return m.Orm.Save(&user).Error
}

func (m *UserDao) DeleteUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}

func (m *UserDao) GetUserByName(name string) (model.User, error) {
	var user model.User
	err := m.Orm.Model(&model.User{}).Where("name=?", name).Find(&user).Error
	return user, err
}

func (m *UserDao) HasUserName(username string) bool {
	var recordNum int64
	m.Orm.Model(&model.User{}).Where("name=?", username).Count(&recordNum)

	return recordNum > 0
}

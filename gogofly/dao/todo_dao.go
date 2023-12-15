package dao

import (
	"gogofly/model"
	"gogofly/service/dto"
)

type UserTodoDao struct {
	BaseDao
}

var userTodoDao *UserTodoDao

func NewUserTodoDao() *UserTodoDao {
	if userTodoDao == nil {
		userTodoDao = &UserTodoDao{
			BaseDao: *NewBaseDao(),
		}
	}
	return userTodoDao
}

func (u *UserTodoDao) AddTodo(userTodoDTO *dto.UserTodoDTO) error {
	var userTodo model.UserTodo
	userTodoDTO.ConvertToModel(&userTodo)

	err := u.Orm.Save(&userTodo).Error
	return err
}

func (u *UserTodoDao) DeleteTodo(id uint) error {
	return u.Orm.Delete(&model.UserTodo{}, id).Error
}

func (u *UserTodoDao) GetTodoList(userId uint) ([]model.UserTodo, int64, error) {
	var (
		userTodoList []model.UserTodo
		todoNum      int64
		err          error
	)

	err = u.Orm.Model(&model.UserTodo{}).Where("user_id=?", userId).Find(&userTodoList).Error
	return userTodoList, todoNum, err
}

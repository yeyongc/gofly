package service

import (
	"gogofly/dao"
	"gogofly/model"
	"gogofly/service/dto"
)

type UserTodoService struct {
	BaseService
	Dao *dao.UserTodoDao
}

func (u *UserTodoService) AddTodo(userTodoDTO *dto.UserTodoDTO) error {
	return u.Dao.AddTodo(userTodoDTO)
}
func (u *UserTodoService) DeleteTodo(id uint) error {
	return u.Dao.DeleteTodo(id)
}
func (u *UserTodoService) GetTodoList(userTodoDTO *dto.UserTodoDTO) ([]model.UserTodo, int64, error) {
	return u.Dao.GetTodoList(userTodoDTO.UserID)
}

var userTodoService *UserTodoService

func NewUserTodoService(dao *dao.UserTodoDao) *UserTodoService {
	if userTodoService == nil {
		userTodoService = &UserTodoService{
			Dao: dao,
		}
	}
	return userTodoService
}

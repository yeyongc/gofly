package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gogofly/dao"
	"gogofly/global"
	"gogofly/global/constants"
	"gogofly/model"
	"gogofly/service/dto"
	"gogofly/utils"
	"time"

	"github.com/spf13/viper"
)

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func (m *UserService) Login(userDto *dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string
	user, err := m.Dao.GetUserByName(userDto.Name)
	if err != nil || utils.IsValidPassword(user.Password, userDto.Password) {
		errResult = errors.New(fmt.Errorf("invalid username or password").Error())
	} else {
		// 校验成功，生成token,并将token缓存进redis
		token, err = GenerateAndCacheTokenToRedis(user.ID, user.Name)
		if err != nil {
			errResult = errors.New(fmt.Errorf("generate token error: %s", err).Error())
		}
	}
	return user, token, errResult
}
func (m *UserService) AddUser(userDTO *dto.UserAddDTO) error {
	if m.Dao.HasUserName(userDTO.Name) {
		return errors.New("user name has existed")
	}
	return m.Dao.AddUser(userDTO)
}
func (m *UserService) GetUserById(idDTO *dto.UserIDDTO) (model.User, error) {
	return m.Dao.GetUserById(idDTO.ID)
}

func (m *UserService) GetUserList(userListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(userListDTO)
}

func (m *UserService) UpdateUser(userUpdateDTO *dto.UserUpdatetDTO) error {
	if userUpdateDTO.ID == 0 {
		return errors.New("invalid User ID")
	}

	return m.Dao.UpdateUser(userUpdateDTO)
}

func (m *UserService) DeleteUser(userIDDTO *dto.UserIDDTO) error {
	return m.Dao.DeleteUserById(userIDDTO.ID)
}

// 生成并将token缓存进redis:  `user:token` -> `name`
func GenerateAndCacheTokenToRedis(userId uint, name string) (string, error) {
	token, err := utils.GenerateToken(userId, name)

	// 生成token出错，直接返回
	if err != nil {
		return token, err
	}

	cacheValues := model.LoginUser{
		Name: name,
		ID:   userId,
	}
	jsonData, _ := json.Marshal(cacheValues)

	// 将token缓存进redis
	err = global.RedisClient.Set(constants.REDIS_USER_PREFIX+token, string(jsonData), viper.GetDuration("jwt.Expire")*constants.REDIS_DURATION_UNIT)

	return token, err
}

// 根据token从redis中获取用户名
func GetUserByTokenFromRedis(token string) (model.LoginUser, error) {
	var user model.LoginUser
	val, err := global.RedisClient.Get(constants.REDIS_USER_PREFIX + token)
	json.Unmarshal([]byte(val), &user)
	return user, err
}

func GetCachedTokenDuration(token string) (time.Duration, error) {
	return global.RedisClient.TTL(context.Background(), constants.REDIS_USER_PREFIX+token)
}

var userService *UserService

func NewUserService(dao *dao.UserDao) *UserService {
	if userService == nil {
		return &UserService{
			Dao: dao,
		}
	}
	return userService
}

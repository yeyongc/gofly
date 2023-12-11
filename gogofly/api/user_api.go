package api

import (
	"fmt"
	"gogofly/api/goflyhttp"
	"gogofly/dao"
	"gogofly/service"
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
}

func AddUser(ctx *gin.Context) {
	var userAddDTO dto.UserAddDTO
	resp := HandleActionAndGenerateResponse(Action{
		Ctx: ctx,
		Req: &Request{
			DTO: &userAddDTO,
		},
		ErrCode: goflyhttp.ERR_CODE_SERVICE_USER_ADD,
		ServiceFunc: func() (any, error) {
			m := service.NewUserService(dao.NewUserDao())
			err := m.AddUser(&userAddDTO)
			return "success", err
		},
	})
	SendResponse(ctx, resp)
}

func UpdateUser(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdatetDTO

	resp := HandleActionAndGenerateResponse(Action{
		Ctx: ctx,
		Req: &Request{
			BindUri: true,
			DTO:     &userUpdateDTO,
		},
		ErrCode: goflyhttp.ERR_CODE_SERVICE_USER_UPDATE,
		ServiceFunc: func() (any, error) {
			m := service.NewUserService(dao.NewUserDao())
			err := m.UpdateUser(&userUpdateDTO)
			return "success", err
		},
	})
	SendResponse(ctx, resp)
}

func DeleteUserById(ctx *gin.Context) {
	var userIDDTO dto.UserIDDTO

	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_USER_DELETE,
		Req: &Request{
			DTO:     &userIDDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewUserService(dao.NewUserDao())
			err := m.DeleteUser(&userIDDTO)
			return "success", err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}

func LoginUser(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_USER_DELETE,
		Req: &Request{
			DTO:     &userLoginDTO,
			BindUri: true,
		},
		Ctx: ctx,
		ServiceFunc: func() (any, error) {
			m := service.NewUserService(dao.NewUserDao())
			user, token, err := m.Login(&userLoginDTO)
			fmt.Println(user, "sdfgh")
			data := map[string]interface{}{
				"user":  user,
				"token": token,
			}
			return data, err
		},
	})
	SendResponse(ctx, res)
}

func GetUserById(ctx *gin.Context) {
	var userIDDTO dto.UserIDDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_USER_DELETE,
		Req: &Request{
			DTO:     &userIDDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewUserService(dao.NewUserDao())
			user, err := m.GetUserById(&userIDDTO)
			data := map[string]interface{}{
				"user": user,
			}
			return data, err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}

func GetUserList(ctx *gin.Context) {
	var userListDTO dto.UserListDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_USER_DELETE,
		Req: &Request{
			DTO:     &userListDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewUserService(dao.NewUserDao())
			users, length, err := m.GetUserList(&userListDTO)
			data := map[string]interface{}{
				"users":  users,
				"length": length,
			}
			return data, err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}

package api

import (
	"gogofly/api/goflyhttp"
	"gogofly/dao"
	"gogofly/service"
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

type UserTodoApi struct {
	BaseApi
}

func AddTodo(ctx *gin.Context) {
	var userTodoDTO dto.UserTodoDTO
	resp := HandleActionAndGenerateResponse(Action{
		Ctx:     ctx,
		ErrCode: goflyhttp.ERR_CODE_SERVICE_TODO_ADD,
		Req: &Request{
			DTO: &userTodoDTO,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewUserTodoService(dao.NewUserTodoDao())
			err := m.AddTodo(&userTodoDTO)
			return "success", err
		},
	})
	SendResponse(ctx, resp)
}

func DeleteTodo(ctx *gin.Context) {
	var userTodoDTO dto.UserTodoDTO
	resp := HandleActionAndGenerateResponse(Action{
		Ctx:     ctx,
		ErrCode: goflyhttp.ERR_CODE_SERVICE_TODO_DELETE,
		Req: &Request{
			DTO: &userTodoDTO,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewUserTodoService(dao.NewUserTodoDao())
			err := m.Dao.DeleteTodo(userTodoDTO.ID)
			return "success", err
		},
	})
	SendResponse(ctx, resp)
}

func GetTodoList(ctx *gin.Context) {
	var userTodoDTO dto.UserTodoDTO
	resp := HandleActionAndGenerateResponse(Action{
		Ctx:     ctx,
		ErrCode: goflyhttp.ERR_CODE_SERVICE_TODO_GET_LIST,
		Req: &Request{
			DTO: &userTodoDTO,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewUserTodoService(dao.NewUserTodoDao())
			res, num, err := m.Dao.GetTodoList(userTodoDTO.UserID)
			data := map[string]interface{}{
				"length": num,
				"todos":  res,
			}
			return data, err
		},
	})
	SendResponse(ctx, resp)
}

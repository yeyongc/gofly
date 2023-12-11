package api

import (
	"gogofly/api/goflyhttp"
	"gogofly/dao"
	"gogofly/service"
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

type AddressApi struct {
	BaseApi
}

func GetAddressLsitByUserId(ctx *gin.Context) {
	var userIDDTO dto.UserIDDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_ADDRESS_GETALL,
		Req: &Request{
			DTO:     &userIDDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewAddressService(dao.NewAddressDao())
			res, err := m.GetAll(userIDDTO.ID)
			data := map[string]interface{}{
				"addresses": res,
				"length":    len(res),
			}
			return data, err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}
func UpdateAddress(ctx *gin.Context) {
	var addressUpdateDTO dto.AddressUpdateDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_ADDRESS_UPDATE,
		Req: &Request{
			DTO:     &addressUpdateDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewAddressService(dao.NewAddressDao())
			err := m.UpdateAddress(&addressUpdateDTO)
			data := "success"
			return data, err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}
func DeleteAddress(ctx *gin.Context) {
	var addressDeleteDTO dto.AddressDeleteDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_ADDRESS_DELETE,
		Req: &Request{
			DTO:     &addressDeleteDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewAddressService(dao.NewAddressDao())
			err := m.Delete(&addressDeleteDTO)
			data := "success"
			return data, err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}

func AddAddress(ctx *gin.Context) {
	var addressAddDTO dto.AddressAddDTO
	res := HandleActionAndGenerateResponse(Action{
		ErrCode: goflyhttp.ERR_CODE_SERVICE_ADDRESS_ADD,
		Req: &Request{
			DTO:     &addressAddDTO,
			BindUri: true,
		},
		ServiceFunc: func() (any, error) {
			m := service.NewAddressService(dao.NewAddressDao())
			err := m.AddAddress(&addressAddDTO)
			data := "success"
			return data, err
		},
		Ctx: ctx,
	})
	SendResponse(ctx, res)
}

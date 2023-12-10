package api

import (
	"gogofly/service"
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

type HostApi struct {
	Service *service.HostService
	BaseApi
}

func NewHostApi(ctx *gin.Context) *HostApi {
	return &HostApi{
		Service: service.NewHostService(),
		BaseApi: NewBaseApi(ctx),
	}
}

func ShutdownHost(ctx *gin.Context) {
	var shutdownDTO dto.ShutdownHostDTO

	m := NewHostApi(ctx)

	if err := m.BuildRequest(&BuildRequestOption{DTO: &shutdownDTO}).GetErrors(); err != nil {
		return
	}

	err := m.Service.Shutdown(&shutdownDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg:  err.Error(),
			Code: 1001,
		})
		return
	}

	m.Ok(ResponseJson{
		Msg: "Shudown success",
	})
}

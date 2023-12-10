package service

import (
	"context"
	"fmt"
	"gogofly/service/dto"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
)

type HostService struct {
	BaseService
}

var hostService *HostService

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (m *HostService) Shutdown(dto *dto.ShutdownHostDTO) error {
	var errResult error
	hostIP := dto.HostIP
	fmt.Println(hostIP)

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user.name"),
	}
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  hostIP,
		ModuleName: "command",
		Args:       viper.GetString("ansible.Shutdown.Args"),
		ExtraVars: map[string]interface{}{
			"ansible_password": viper.GetString("ansible.user.password"),
		},
	}

	adc := &adhoc.AnsibleAdhocCmd{
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "online",
	}

	errResult = adc.Run(context.TODO())

	return errResult
}

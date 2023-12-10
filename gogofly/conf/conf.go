package conf

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")

	err := viper.ReadInConfig()

	if err != nil {
		panic("Read config error: " + err.Error())
	}
}

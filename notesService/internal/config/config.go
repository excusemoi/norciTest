package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

func InitConfig(filePath string) error {
	viper.AddConfigPath(filePath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

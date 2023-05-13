package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Values struct {
	HttpPort   string `envconfig:"http_port" mapstructure:"HTTP_PORT"`
	GrpcPort   string `envconfig:"grpc_port" mapstructure:"GRPC_PORT"`
	DBDriver   string `envconfig:"db_driver" mapstructure:"DB_DRIVER" default:"mysql"`
	DBUsername string `envconfig:"db_username" mapstructure:"DB_USERNAME"`
	DBPassword string `envconfig:"db_password" mapstructure:"DB_PASSWORD"`
	DBHost     string `envconfig:"db_host" mapstructure:"DB_HOST"`
	DBPort     string `envconfig:"db_port" mapstructure:"DB_PORT" default:"3306"`
	DBName     string `envconfig:"db_name" mapstructure:"DB_NAME"`
	DBTest     string `envconfig:"db_test" mapstructure:"DB_TEST"`
}

func NewValues() (*Values, error) {
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			return nil, err
		}
	}

	v := &Values{}

	if err := viper.Unmarshal(v); err != nil {
		return nil, err
	}

	return v, nil
}

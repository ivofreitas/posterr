package config

import (
	"github.com/spf13/viper"
	"sync"
	"time"
)

var (
	c    *config
	once sync.Once
)

type config struct {
	Server Server `mapstructure:"server"`
	MySQL  MySQL  `mapstructure:"mysql"`
}

type Server struct {
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
	Port     string `mapstructure:"port"`
}

type MySQL struct {
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	Host         string        `mapstructure:"host"`
	Database     string        `mapstructure:"database"`
	PoolConn     int           `mapstructure:"pool_conn"`
	ConnLifetime time.Duration `mapstructure:"conn_lifetime"`
}

func Get() *config {

	once.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("json")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			return
		}

		viper.Unmarshal(&c)
	})

	return c
}

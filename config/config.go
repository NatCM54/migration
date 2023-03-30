package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbMysqlUsername string `mapstructure:"DB_MYSQL_USERNAME"`
	DbMysqlPassword string `mapstructure:"DB_MYSQL_PASSWORD"`
	DbMysqlPort     string `mapstructure:"DB_MYSQL_PORT"`
	DbMysqlHost     string `mapstructure:"DB_MYSQL_HOST"`
	DbMysqlDatabase string `mapstructure:"DB_MYSQL_DATABASE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("local.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	return
}

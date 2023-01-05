package config

import (
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Redis  RedisConf  `mapstructure:"redis"`
	Server ServerConf `mapstructure:"server"`
}

type RedisConf struct {
	ExpireTime int    `mapstructure:"expire_time"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Db         int    `mapstructure:"db"`
	Password   string `mapstructure:"password"`
}

type ServerConf struct {
	Port   int    `mapstructure:"port"`
	ApiKey string `mapstructure:"api_key"`
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("retriever") // automatically turns to capitalized
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	return
}

package config

import (
	constants "ecom/constants"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ConfigDefaultFile = "config/config.yml"
	ConfigReleaseFile = "config/config.release.yml"
	ConfigDevFile     = "config/config.dev.yml"
	configType        = "yml"
)

type (
	Config struct {
		Debug          bool     `mapstructure:"debug"`
		ContextTimeout int      `mapstructure:"contextTimeout"`
		Server         Server   `mapstructure:"server"`
		Services       Services `mapstructure:"services"`
		Database       Database `mapstructure:"database"`
		Logger         Logger   `mapstructure:"logger"`
		Jwt            Jwt      `mapstructure:"jwt"`
		Redis          Redis    `mapstructure:"redis"`
	}

	Server struct {
		Host     string `mapstructure:"host"`
		Env      string `mapstructure:"env"`
		UseRedis bool   `mapstructure:"useRedis"`
		Port     int    `mapstructure:"port"`
	}

	Database struct {
		Driver             string `mapstructure:"driver"`
		Host               string `mapstructure:"host"`
		Port               int    `mapstructure:"port"`
		Username           string `mapstructure:"username"`
		Password           string `mapstructure:"password"`
		Name               string `mapstructure:"name"`
		MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
		MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
		ConnMaxLifetime    int    `mapstructure:"connMaxLifetime"`
		TimeZone           string `mapstructure:"timeZone"`
		SSLMode            string `mapstructure:"sslMode"`
	}

	Jwt struct {
		Secret                string `mapstructure:"secret"`
		AccessTokenExpiresIn  int64  `mapstructure:"accessTokenExpiresIn"`
		RefreshTokenExpiresIn int64  `mapstructure:"refreshTokenExpiresIn"`
	}

	Logger struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		Prefix string `mapstructure:"prefix"`
	}

	Redis struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		Password     string `mapstructure:"password"`
		PoolSize     int    `mapstructure:"poolSize"`
		MinIdleConns int    `mapstructure:"minIdleConns"`
		DB           int    `mapstructure:"db"`
		Name         string `mapstructure:"name"`
	}

	Services struct {
	}
)

func NewConfig() *Config {
	initConfig()
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return conf
}

func initConfig() {
	var configFile string
	switch os.Getenv("ENV") {
	case constants.Prod:
		configFile = ConfigReleaseFile
		fmt.Printf("gin mode %s,config%s\n", gin.EnvGinMode, ConfigReleaseFile)
	case constants.Dev:
		configFile = ConfigDevFile
		fmt.Printf("gin mode: %s,config file: %s\n", gin.EnvGinMode, ConfigDevFile)
	default:
		configFile = ConfigDefaultFile
		fmt.Printf("gin mode %s,config%s\n", gin.EnvGinMode, ConfigDefaultFile)
	}
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}

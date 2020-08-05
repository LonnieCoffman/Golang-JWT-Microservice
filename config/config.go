package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	DB          *gorm.DB
	DBErr       error
	SERVER_HOST string
	SERVER_PORT int
	DB_HOST     string
	DB_DRIVER   string
	DB_USER     string
	DB_PASS     string
	DB_NAME     string
	DB_PORT     int64
	DB_DSN      string
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	Config.DB_DSN = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		v.Get("DB_USER").(string),
		v.Get("DB_PASS").(string),
		v.Get("DB_HOST").(string),
		v.GetInt("DB_PORT"),
		v.Get("DB_NAME").(string),
	)

	return v.Unmarshal(&Config)
}

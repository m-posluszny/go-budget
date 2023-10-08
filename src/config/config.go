package config

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

var config *Config

func Get() Config {
	return *config
}

func getAppMode() string {
	modes := []string{"dev", "prod", "test"}
	appMode := strings.ToLower(os.Getenv("APP_MODE"))
	if slices.Contains(modes, appMode) {
		panic(errors.New("you need to set APP_MODE to one of prod,develop,test"))
	}
	return appMode
}

func Init() Config {
	replacer := strings.NewReplacer(".", "_")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	viper.Debug()
	viper.AutomaticEnv()
	return *config
}

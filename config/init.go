package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Current *AppConfig

func init() {
	current, err := readConfig()
	if err != nil {
		log.Fatalf("application config could not be loaded, cause: %s", err)
	}
	Current = &current
}

func readConfig() (AppConfig, error) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("GO_APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, fmt.Errorf("fatal error config file: %w \n", err)
	}

	var appConfig AppConfig
	if err = viper.Unmarshal(&appConfig); err != nil {
		return AppConfig{}, err
	}

	return appConfig, nil
}

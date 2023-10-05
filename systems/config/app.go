package config

import (
	"fmt"
	"go/build"
	"os"

	"github.com/spf13/viper"
	//driver
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// set config based on env
	LoadEnvVars()
	OpenDbPool()
}

func LoadEnvVars() {
	var AppPath string
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	// Bind OS environment variable
	viper.SetEnvPrefix("app")
	viper.BindEnv("env")
	viper.BindEnv("path") // bind APP_PATH variable

	viper.SetConfigName("config")
	dir, _ := os.Getwd()
	AppPath = dir

	viper.SetConfigType("json")
	viper.AddConfigPath(AppPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

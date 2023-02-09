package config

import (
	"fmt"

	"github.com/spf13/viper"
	yml "gopkg.in/yaml.v3"
)

/*
SpeedtestWrapperConfig is an abstraction for the app config
*/
type SpeedtestWrapperConfig struct {
	DBFile string `yaml:"db_file"`
}

/*
AppConfig loads the config from file
*/
func AppConfig() SpeedtestWrapperConfig {
	var config SpeedtestWrapperConfig
	if err := viper.UnmarshalKey("cfg", &config); err != nil {
		panic(err.Error())
	}

	return config
}

/*
WriteConfig writes config to file
*/
func WriteConfig(cfg SpeedtestWrapperConfig) string {
	if cfgBytes, err := yml.Marshal(cfg); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(string(cfgBytes))
		return string(cfgBytes)
	}
}

/*
GetDBFile returns the DB file path
*/
func GetDBFile() string {
	return viper.GetString("cfg.db_file")
}

var (
	Version   string
	BuildDate string
)

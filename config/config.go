package config

import (
	"fmt"

	"github.com/spf13/viper"
	yml "gopkg.in/yaml.v3"
)

/*
MQTTConfig is a config abstraction for the MQTT client
*/
type MQTTConfig struct {
	Host       string
	Port       string
	User       string
	Pass       string
	ClientName string
	Topic      string
	Qos        byte
}

/*
SpeedtestWrapperConfig is an abstraction for the app config
*/
type SpeedtestWrapperConfig struct {
	MQTT MQTTConfig
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

var (
	Version   string
	BuildDate string
)

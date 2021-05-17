package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

const (
	configSample = "samples/config_test.yml"
)

func setup(t *testing.T) {
	currPath, err := os.Getwd()
	if err != nil {
		t.Errorf("Error trying to find current dir")
	}
	fmt.Println("current path", currPath)
	configFile := filepath.Join(currPath, configSample)
	fmt.Println("config file", configFile)

	if _, err := os.Stat(configFile); err != nil {
		t.Errorf("Config file doesn't exists:\n'%s'", configFile)
	}

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		t.Errorf("Failed using config file: %v\n%v", viper.ConfigFileUsed(), err.Error())
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestAppConfig(t *testing.T) {
	setup(t)
	cfg := AppConfig()

	fmt.Sprintln(cfg.MQTT.Host)
	if cfg.MQTT.Host != "localhost" {
		t.Errorf("Wrong host value: %s", cfg.MQTT.Host)
	}
	fmt.Sprintln(cfg.MQTT.Qos)
	if cfg.MQTT.Qos != 1 {
		t.Errorf("Wrong Qos value: %d", cfg.MQTT.Qos)
	}
}

package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func BindEnvVars() {
	bindEnv("cfg.mqtt.host", "SPEEDTEST_MQTT_HOST")
	bindEnv("cfg.mqtt.port", "SPEEDTEST_MQTT_PORT")
	bindEnv("cfg.mqtt.user", "SPEEDTEST_MQTT_USER")
	bindEnv("cfg.mqtt.pass", "SPEEDTEST_MQTT_PASS")
	bindEnv("cfg.mqtt.clientName", "SPEEDTEST_MQTT_CLIENTNAME")
	bindEnv("cfg.mqtt.topic", "SPEEDTEST_MQTT_TOPIC")
	bindEnv("cfg.mqtt.qos", "SPEEDTEST_MQTT_QOS")
}

func bindEnv(key string, envVar string) {
	if err := viper.BindEnv(key, envVar); err != nil {
		log.Panic(fmt.Sprintf("Failed to bind config key '%s' to environment variable '%s'", key, envVar))
	}
}

func SetDefaults() {
	viper.SetDefault("cfg.mqtt.port", "1883")
}

func InitConfig(cfgFile string) {
	BindEnvVars()
	SetDefaults()
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".raspberry-network-monitor" (without extension).
		//viper.AddConfigPath(home)
		viper.AddConfigPath(filepath.Join(home, ".speedtest-wrapper"))
		viper.AddConfigPath("/etc/speedtest-wrapper")
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

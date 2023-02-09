package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func BindEnvVars() {
	bindEnv("cfg.db_file", "SPEEDTEST_DB_FILE")
}

func bindEnv(key string, envVar string) {
	if err := viper.BindEnv(key, envVar); err != nil {
		log.Panic(fmt.Sprintf("Failed to bind config key '%s' to environment variable '%s'", key, envVar))
	}
}

func SetDefaults() {
	viper.SetDefault("cfg.db_file", "./speedtest.db")
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
		viper.AddConfigPath(home)
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

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App         AppConfig
	Environment EnvironmentConfig
	AppDb       MySQLConfig     `mapstructure:"app_db"`
	DbWatcher   DbWatcherConfig `mapstructure:"db_watcher"`
}

type AppConfig struct {
	Name string
}

type EnvironmentConfig struct {
	Production bool
	SeedDb     bool
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type DbWatcherConfig struct {
	Schema string
	Table  string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("app.name", "default app name")

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode config into struct, %v", err)
	}

	return config
}

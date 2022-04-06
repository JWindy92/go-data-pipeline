package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config c.Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("app.name", "default app name")

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode config into struct, %v", err)
	}

	fmt.Printf("Starting %s", config.App.Name)
	fmt.Printf("Production: %s", config.Environment.Production)
	fmt.Printf("Seed DB: %s", config.Environment.SeedDb)
}

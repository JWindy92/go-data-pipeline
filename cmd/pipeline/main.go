package main

// will subscribe to kafka topic and grab records from the db's accordingly

import (
	"fmt"

	c "github.com/JWindy92/go-data-pipeline/config"
)

func main() {

	config := c.LoadConfig()

	fmt.Printf("Starting %s\n", config.App.Name)
	fmt.Printf("Production: %v\n", config.Environment.Production)
	fmt.Printf("Seed DB: %v\n", config.Environment.SeedDb)
	fmt.Printf(config.AppDb.Host)
}

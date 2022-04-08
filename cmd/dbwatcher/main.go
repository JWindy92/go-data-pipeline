package main

import (
	"fmt"

	"github.com/JWindy92/go-data-pipeline/pkg/binlogger"
)

func main() {
	fmt.Printf("starting dbwatcher")

	binlogger.Listen()
}

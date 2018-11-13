package main

import (
	"fmt"

	"github.com/Ilyes512/goreleaserpoc/pkg/configuration"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Version: ", configuration.Version)
	fmt.Println("Commit: ", configuration.Commit)
	fmt.Println("Date: ", configuration.Date)
}

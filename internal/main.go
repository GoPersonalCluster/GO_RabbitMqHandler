package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Getenv("APP_ENV")

	switch env {
	case "production":
		fmt.Println("Running in production")
	case "development":
		fmt.Println("Running in development")
	default:
		fmt.Println("Running in local mode")
	}
}
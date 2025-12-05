package main

import (
	"effMobile/internal/app"
	"effMobile/pkg/config"
	"log"
)

func init() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Error from config: %v", err)
	}
}

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}
}

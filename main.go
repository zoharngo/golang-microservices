package main

import (
	"log"
	"os"

	"github.com/zoharngo/golang-microservices/mvc/app"
)

var (
	logger = log.New(os.Stdout, "main => ", log.Ltime)
)

func main() {
	logger.Println("Starting ..")
	if err := app.StartUp(); err != nil {
		logger.Fatalf("Failed to start application service: [%s]", err.Error())
		os.Exit(1)
	}
}

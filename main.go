package main

import (
	"HaiooTechnicalTest-Golang/answer"
	"HaiooTechnicalTest-Golang/api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	answer.MyAnswer()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api.Init()
}

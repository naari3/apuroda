package main

import (
	"apuroda/db"
	"apuroda/server"
	"apuroda/stores"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	}
	db.Init()
	stores.Init()
	server.Start()
	db.Close()
}

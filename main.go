package main

import (
	"log"

	"github.com/Suryaakula888/database"
	"github.com/Suryaakula888/routers"
)

func main() {
	database.Setup()
	engine := routers.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}

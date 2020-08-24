package main

import (
	"backend/api"
	"backend/storage"
	"log"
)

func main() {
	router := api.Router{
		Port: 8080,
	}
	router.Run()

	readConf()
}

func readConf() {
	db, err := storage.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

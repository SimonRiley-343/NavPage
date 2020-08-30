package main

import (
	"backend/api"
	"backend/storage"
	"log"
)

func main() {
	err := storage.CheckTable()
	if err != nil {
		log.Fatalln(err)
	}

	router := api.Router{
		Port: 8080,
	}
	router.Init()

	router.SetPost("/api/login", api.Login)
	router.SetPost("/api/pages", api.Pages)

	err = router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"backend/api"
	"backend/storage"
	"flag"
	"log"
)

func main() {
	err := storage.CheckTable()
	if err != nil {
		log.Fatalln(err)
	}

	port := *flag.Int("p", 0, "Set port")

	if port == 0 {
		conf := storage.ConfData{}
		port, err = conf.Port()
		if err != nil {
			log.Fatalln(err)
		}
	}

	router := api.Router{
		Port: port,
	}
	router.Init()

	router.SetPost("/api/login", api.Login)
	router.SetPost("/api/pages", api.Pages)

	err = router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

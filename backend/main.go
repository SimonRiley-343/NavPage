package main

import (
	"backend/api"
	"log"
)

func main() {
	data := api.Data{}
	err := data.CheckData()
	if err != nil {
		log.Fatalln(err)
	}

	router := api.Router{
		Port: 8080,
	}
	router.Init()
	err = router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

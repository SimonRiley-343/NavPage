package main

import (
	"backend/api"
	"backend/storage"
	"fmt"
	"log"
)

func main() {
	err := storage.CheckTable()
	if err != nil {
		log.Fatalln(err)
	}

	ud := storage.UserData{}
	auth, err := ud.CheckPasswd(1001, "passwda")
	if err != nil {
		log.Fatalln(err)
	}
	if auth {
		fmt.Println("Right password")
	} else {
		fmt.Println("Wrong password")
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

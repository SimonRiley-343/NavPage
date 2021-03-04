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

    var port int
    flag.IntVar(&port, "p", 0, "Set port")

    flag.Parse()

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

    router.SetPost("/login", api.Login)
    router.SetPost("/pages", api.Pages)

    err = router.Run()
    if err != nil {
        log.Fatalln(err)
    }
}

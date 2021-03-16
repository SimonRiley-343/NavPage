package main

import (
    "backend/api"
    "backend/model"
    "backend/storage"
    "flag"
    "log"
    "strconv"
)

func main() {
    err := storage.CheckTable()
    if err != nil {
        log.Fatalln(err)
    }

    var port string
    flag.StringVar(&port, "p", "", "Set port")

    flag.Parse()

    if port == "" {
        conf := storage.ConfData{}
        port, err = conf.GetConf(model.DB_CONFIG_PORT)
        if err != nil {
            panic(err)
        }
        if port == "" {
            port = model.DB_CONFIG_DEFAULT_SESSIONSECRET
        }
    }

    portNum, _ := strconv.Atoi(port)

    router := api.Router{
        Port: portNum,
    }
    router.Init()

    routerApi := router.Router.Group("/api")
    {
        routerApi.POST("/login", api.Login)
        routerApi.POST("/pages", api.Pages)
    }

    routerAdminApi := router.Router.Group("/api")
    {
        routerAdminApi.Use(router.Auth())
        routerAdminApi.POST("/addPage", api.AddPage)
    }

    err = router.Run()
    if err != nil {
        log.Fatalln(err)
    }
}

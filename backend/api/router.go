package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Router struct {
	Port int
}

func (r Router) Run() {
	router := gin.Default()
	router.LoadHTMLGlob("webdist/*.html")
	router.Static("/", "webdist")

	router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Run(":" + strconv.Itoa(r.Port))
}
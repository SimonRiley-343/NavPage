package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Router struct {
	Port   int
	Router *gin.Engine
}

func (r *Router) Init() {
	r.Router = gin.Default()
	r.Router.LoadHTMLGlob("webdist/*.html")
	r.Router.Static("/", "webdist")

	r.Router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func (r *Router) SetPost(relativePath string, handlerFunc gin.HandlerFunc) {
	r.Router.POST(relativePath, handlerFunc)
}

func (r *Router) Run() error {
	return r.Router.Run(":" + strconv.Itoa(r.Port))
}

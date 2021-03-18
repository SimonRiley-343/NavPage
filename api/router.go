package api

import (
    "backend/storage"
    "github.com/gin-contrib/cors"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/memstore"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "time"

    "backend/model"
)

type Router struct {
    Port   int
    Router *gin.Engine
}

func (r *Router) Init() {
    conf := storage.ConfData{}

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

    sessionSecret, err := conf.GetConf(model.DB_CONFIG_SESSIONSECRET)
    if err != nil {
        panic(err)
    }

    if sessionSecret == "" {
        sessionSecret = model.DB_CONFIG_DEFAULT_SESSIONSECRET
    }

    store := memstore.NewStore([]byte(sessionSecret))
    r.Router.Use(sessions.Sessions(model.SESSION_NAME, store))
}

func (r *Router) Run() error {
    return r.Router.Run(":" + strconv.Itoa(r.Port))
}

func (r *Router) Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        resultData := model.ResultGinAuth{
            Code: model.REQ_CODE_LOGIN_FAILED,
            Msg:  model.REQ_MSG_LOGIN_FAILED,
        }

        sessionId, err := c.Cookie(model.COOKIE_KEY_SESSIONID)
        if err != nil || (storage.GetSession(c, sessionId) == "") {
            c.Abort()
            c.JSON(http.StatusUnauthorized, resultData)
            return
        }

        c.Next()
    }
}

package storage

import (
    "backend/model"
    "fmt"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, value string) {
    session := sessions.Default(c)
    session.Set(key, value)
    err := session.Save()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func GetSession(c *gin.Context, key string) string {
    session := sessions.Default(c)
    sessionStr := ""
    sessionValue := session.Get(key)

    if sessionValue != nil {
        sessionStr = sessionValue.(string)
    }

    return sessionStr
}

func GetUserIdbyCookie(c *gin.Context) (string, error) {
    sessionId, err := c.Cookie(model.COOKIE_KEY_SESSIONID)
    if err != nil {
        return "", err
    }
    return GetSession(c, sessionId), nil
}

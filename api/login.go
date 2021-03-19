package api

import (
    "backend/model"
    "backend/storage"
    "backend/utils"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "strconv"
    "time"
)

func Login(c *gin.Context) {
    passwd := c.PostForm(model.POST_FORM_PASSWD)

    reqMsg := model.ResultMsgLogin{}
    reqMsgLoginFailed := model.ResultMsgLogin{
        Code:  model.REQ_CODE_LOGIN_FAILED,
        Msg:   model.REQ_MSG_LOGIN_FAILED,
    }

    loggedSessionId, _ := c.Cookie(model.COOKIE_KEY_SESSIONID)
    if loggedSessionId != "" {
        if storage.GetSession(c, loggedSessionId) != "" {
            reqMsg.Code = model.REQ_CODE_SUCCESS
            reqMsg.Msg = model.REQ_MSG_SUCCESS
            reqMsg.Login = true
        } else {
            reqMsg.Code = model.REQ_CODE_LOGIN_FAILED
            reqMsg.Msg = model.REQ_MSG_LOGIN_FAILED
            reqMsg.Login = false
        }
        c.JSON(http.StatusOK, reqMsg)
        return
    }

    if passwd == "" {
        reqMsg.Code = model.REQ_CODE_PARAM_ERR
        reqMsg.Msg = model.REQ_MSG_PARAM_ERR
        c.JSON(http.StatusOK, reqMsg)
        return
    }

    conf := storage.ConfData{}

    if isLogin, err := conf.CheckPasswd(passwd); err != nil {
        log.Fatalln(err)
    } else if !isLogin {
        c.JSON(http.StatusOK, reqMsgLoginFailed)
        return
    } else if isLogin {
        reqMsg.Code = model.REQ_CODE_SUCCESS
        reqMsg.Msg = model.REQ_MSG_SUCCESS
        reqMsg.Login = true

        sessionId := utils.RandomStr(model.SESSIONID_LEN)
        storage.SetSession(c, sessionId, strconv.FormatInt(time.Now().Unix(), 10))

        c.SetCookie(model.COOKIE_KEY_SESSIONID, sessionId, 0,
            "/", "", false, false)

        c.JSON(http.StatusOK, reqMsg)
        return
    }

    reqMsg.Code = model.REQ_CODE_UNKNOWN_ERR
    reqMsg.Msg = model.REQ_MSG_UNKNOWN_ERR
    reqMsg.Login = false
    c.JSON(http.StatusInternalServerError, reqMsg)
    return
}

package api

import (
	"backend/model"
	"backend/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	passwd := c.PostForm(model.POST_FORM_PASSWD)

	reqMsg := model.ResultMsgLogin{}
	reqMsgLoginFailed := model.ResultMsgLogin{
		Code:  model.REQ_CODE_LOGIN_FAILED,
		Login: false,
		Msg:   model.REQ_MSG_LOGIN_FAILED,
	}

	if passwd == "" {
		reqMsg.Code = model.REQ_CODE_PARAM_ERR
		reqMsg.Login = false
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
		reqMsg.Login = true
		reqMsg.Msg = model.REQ_MSG_SUCCESS
		c.JSON(http.StatusOK, reqMsg)
		return
	}
	reqMsg.Code = model.REQ_CODE_UNKNOWN_ERR
	reqMsg.Login = false
	reqMsg.Msg = model.REQ_MSG_UNKNOWN_ERR
	c.JSON(http.StatusInternalServerError, reqMsg)
	return
}

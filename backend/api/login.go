package api

import (
	"backend/model"
	"backend/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	user := c.PostForm(model.POST_FORM_USER)
	passwd := c.PostForm(model.POST_FORM_PASSWD)

	reqMsg := model.ResultMsg{}
	reqMsgLoginFailed := model.ResultMsg{
		Code: model.REQ_CODE_LOGIN_FAILED,
		Msg:  model.REQ_MSG_LOGIN_FAILED,
	}

	if user == "" || passwd == "" {
		reqMsg.Code = model.REQ_CODE_PARAM_ERR
		reqMsg.Msg = model.REQ_MSG_PARAM_ERR
		c.JSON(http.StatusOK, reqMsg)
		return
	}

	ud := storage.UserData{}
	id, err := ud.GetUserId(user)
	if err != nil {
		log.Fatalln(err)
	}
	if id == 0 {
		c.JSON(http.StatusOK, reqMsgLoginFailed)
		return
	}

	if isCorrect, err := ud.CheckPasswd(id, passwd); err != nil {
		log.Fatalln(err)
	} else if !isCorrect {
		c.JSON(http.StatusOK, reqMsgLoginFailed)
		return
	} else if isCorrect {
		reqMsg.Code = model.REQ_CODE_SUCCESS
		reqMsg.Id = id
		reqMsg.Msg = model.REQ_MSG_LOGIN_SUCCESS
		c.JSON(http.StatusOK, reqMsg)
		return
	}

	reqMsg.Code = model.REQ_CODE_UNKNOWN_ERR
	reqMsg.Msg = model.REQ_MSG_UNKNOWN_ERR
	c.JSON(http.StatusInternalServerError, reqMsg)
	return
}

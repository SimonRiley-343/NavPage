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

	reqMsg := model.ResultMsgLogin{}
	reqMsgLoginFailed := model.ResultMsgLogin{
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

	if userId, err := ud.CheckPasswd(user, passwd); err != nil {
		log.Fatalln(err)
	} else if userId == 0 {
		c.JSON(http.StatusOK, reqMsgLoginFailed)
	} else {
		reqMsg.Code = model.REQ_CODE_SUCCESS
		reqMsg.Id = userId
		reqMsg.Msg = model.REQ_MSG_SUCCESS
		c.JSON(http.StatusOK, reqMsg)
	}
	return
}

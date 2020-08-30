package api

import (
	"backend/model"
	"backend/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Pages(c *gin.Context) {
	reqMsg := model.ResultMsgPages{}

	ud := storage.PageData{}
	pageList, err := ud.GetAllPage()
	if err != nil {
		log.Fatalln(err)
	}

	if len(pageList) == 0 {
		reqMsg.Code = model.REQ_CODE_NO_PAGE
		reqMsg.Msg = model.REQ_MSG_PAGE_NO
		c.JSON(http.StatusOK, reqMsg)
		return
	}

	reqMsg.Code = model.REQ_CODE_SUCCESS
	reqMsg.Msg = model.REQ_MSG_SUCCESS
	reqMsg.Pages = pageList
	c.JSON(http.StatusOK, reqMsg)
	return
}

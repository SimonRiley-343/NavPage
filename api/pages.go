package api

import (
    "backend/model"
    "backend/storage"
    "github.com/gin-gonic/gin"
    bolt "go.etcd.io/bbolt"
    "log"
    "net/http"
)

func Pages(c *gin.Context) {
    reqMsg := model.ResultMsgPages{
        Code: model.REQ_CODE_UNKNOWN_ERR,
        Msg: model.REQ_MSG_UNKNOWN_ERR,
    }

    pageData := storage.PageData{}
    pageList, err := pageData.GetAllPage()
    if err != nil {
        log.Fatalln(err)
    }

    if len(pageList) == 0 {
        reqMsg.Code = model.REQ_CODE_NO_PAGE
        reqMsg.Msg = model.REQ_MSG_NO_PAGE
        c.JSON(http.StatusOK, reqMsg)
        return
    }

    reqMsg.Code = model.REQ_CODE_SUCCESS
    reqMsg.Msg = model.REQ_MSG_SUCCESS
    reqMsg.Pages = pageList
    c.JSON(http.StatusOK, reqMsg)
    return
}

func AddPage(c *gin.Context) {
    s, err := storage.Open()
    if err != nil {
        log.Fatalln(err)
    }
    defer s.Close()

    reqMsg := model.ResultMsgAddPage{
        Code: model.REQ_CODE_UNKNOWN_ERR,
        Msg: model.REQ_MSG_UNKNOWN_ERR,
    }

    pageName := c.PostForm(model.POST_FORM_PAGE_NAME)
    pageCat := c.PostForm(model.POST_FORM_PAGE_CAT)
    pageDesc := c.PostForm(model.POST_FORM_PAGE_DESC)
    pageUrl := c.PostForm(model.POST_FORM_PAGE_URL)

    if (len(pageName) == 0) || (len(pageCat) == 0) || (len(pageDesc) == 0) || (len(pageUrl) == 0) {
        reqMsg.Code = model.REQ_CODE_PARAM_ERR
        reqMsg.Msg = model.REQ_MSG_PARAM_ERR
        c.JSON(http.StatusBadRequest, reqMsg)
        return
    }

    var pageId string

    err = s.DB.Update(func(tx *bolt.Tx) error {
        bucketPage := tx.Bucket([]byte(model.DB_NAME_PAGE))

        pageData := storage.PageData{}

        pageId, err = pageData.InsertPage(bucketPage, pageName, pageCat, pageDesc, pageUrl)
        return err
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, reqMsg)
        log.Fatalln(err)
    }

    reqMsg.Code = model.REQ_CODE_SUCCESS
    reqMsg.Msg = model.REQ_MSG_SUCCESS
    reqMsg.PageId = pageId
    c.JSON(http.StatusOK, reqMsg)
    return
}

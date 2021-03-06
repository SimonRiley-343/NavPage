package model

type ResultMsgLogin struct {
    Code  int    `json:"code"`
    Msg   string `json:"msg"`
    Login bool   `json:"login"`
}

type ResultMsgPages struct {
    Code  int     `json:"code"`
    Msg   string  `json:"msg"`
    Pages []Pages `json:"pages,omitempty"`
}

type ResultMsgAddPage struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    PageId string `json:"pageId,omitempty"`
}

type ResultGinAuth struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
}

type Pages struct {
    Id   string `json:"id"`
    Name string `json:"name"`
    Cat  string `json:"cat"`
    Desc string `json:"desc"`
    Url  string `json:"url"`
}

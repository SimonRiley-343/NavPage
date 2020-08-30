package model

type ResultMsgLogin struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Id   int    `json:"id,omitempty"`
}

type ResultMsgPages struct {
	Code  int     `json:"code"`
	Msg   string  `json:"msg"`
	Pages []Pages `json:"pages,omitempty"`
}

type Pages struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

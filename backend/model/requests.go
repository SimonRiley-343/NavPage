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
	Url  string `json:"url"`
	Img  string `json:"img"`
}

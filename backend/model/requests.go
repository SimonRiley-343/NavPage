package model

type ResultMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Id   int    `json:"id,omitempty"`
}

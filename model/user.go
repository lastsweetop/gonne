package model

type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

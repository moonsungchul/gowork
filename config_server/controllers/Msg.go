package controllers

type Msg struct {
	MsgID   int    `json:"msg_id"`
	Message string `json:"message"`
}

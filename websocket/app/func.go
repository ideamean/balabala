package app

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error_code int
	Error_msg  string
}

type RetMsg struct {
	OnlineUserNum int
	State         int
}

func Resp_Error(w http.ResponseWriter, code int, msg string) {
	err := &Error{Error_code: code, Error_msg: msg}
	Resp_Json(w, err)
}

func Resp_Json(w http.ResponseWriter, data interface{}) {
	ret, _ := json.Marshal(data)
	w.Write([]byte(ret))
}

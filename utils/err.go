package utils

import "net/http"

const ERR_OBJ_NOT_FOUND = 101
const ERR_JSON_PARSE = 102
const ERR_UNKNOWN = 999

type AppErr interface {
	Error() string
	ErrCode() int
	HTTPStatus() int
	HTTPMsg() string
}

type appErr struct {
	msg  string
	code int
}

func NewAppErr(msg string, code int) appErr {
	return appErr{msg: msg, code: code}
}

func (r appErr) Error() string {
	return r.msg
}

func (r appErr) ErrCode() int {
	return r.code
}

func (r appErr) HTTPStatus() int {
	if r.code == ERR_OBJ_NOT_FOUND {
		return http.StatusNotFound
	} else if r.code == ERR_JSON_PARSE {
		return http.StatusBadRequest
	} else {
		return http.StatusInternalServerError
	}
}

func (r appErr) HTTPMsg() string {
	if r.code == ERR_OBJ_NOT_FOUND {
		return "Item Not Found"
	} else {
		return r.Error()
	}
}

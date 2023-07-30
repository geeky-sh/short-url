package utils

import (
	"encoding/json"
	"net/http"
)

func WriteMsgRes(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	en := json.NewEncoder(w)
	en.Encode(map[string]string{"msg": msg})
}

func WriteAppErrRes(w http.ResponseWriter, err AppErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatus())
	en := json.NewEncoder(w)
	en.Encode(map[string]string{"msg": err.HTTPMsg()})
}

func WriteRes(w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	en := json.NewEncoder(w)
	en.Encode(res)
}

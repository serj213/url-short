package httpserver

import (
	"encoding/json"
	"net/http"
)


type ResponseErr struct {
	Status string `json:"status"`
	Msg string	`json:"msg"`
}

type ResponseSuccess struct {
	Status string `json:"status"`
	Alias string `json:"alias"`
}

func ErrorResponse(slug string, w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(
		ResponseErr{
			Status: "failed",
			Msg: slug,
		},
	);
}


func ResponseOk(data any, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}
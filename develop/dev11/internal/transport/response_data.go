package transport

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type SuccessResponse struct {
	Result interface{} `json:"result,omitempty"`
}

type BadResponse struct {
	Err string `json:"err,omitempty"`
}

func badResponse(w http.ResponseWriter, code int, response string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	body, err := json.Marshal(&BadResponse{response})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if _, err = w.Write(body); err != nil {
		slog.Error(err.Error())
	}
}

func successResponse(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	body, err := json.Marshal(&SuccessResponse{Result: "Success"})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if _, err = w.Write(body); err != nil {
		slog.Error(err.Error())
	}
}

func successResponseData(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	body, err := json.Marshal(&SuccessResponse{data})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if _, err = w.Write(body); err != nil {
		slog.Error(err.Error())
	}
}

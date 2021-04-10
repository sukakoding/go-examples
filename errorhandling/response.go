package main

import (
	"encoding/json"
	"net/http"
)

type (
	errorResponse struct {
		Code    uint32 `json:"code"`
		Message string `json:"message"`
	}

	successResponse struct {
		Message string `json:"message"`
	}
)

// JSON write status and JSON data to http response writer
func JSON(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(b)
}

// Error write error, status to http response writer with JSON format: {"code": status, "message": error}
func Error(w http.ResponseWriter, err error) {

	c := codeUnknown
	s := http.StatusInternalServerError
	m := "Unknown error, please contact administrator!"

	if v, ok := err.(Problem); ok {
		c = v.Code()
		s = v.Status()
		m = v.Message()
	}

	// TODO log, warning
	r := &errorResponse{
		Code:    c,
		Message: m,
	}

	JSON(w, s, r)
}

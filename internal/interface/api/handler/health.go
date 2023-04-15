package handler

import (
	"net/http"
)

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

type HealthHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
}

func (h healthHandler) Get(w http.ResponseWriter, r *http.Request) {
	//nolint:errcheck
	w.Write([]byte("ok"))
}

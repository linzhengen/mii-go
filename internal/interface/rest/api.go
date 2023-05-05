package rest

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/linzhengen/mii-go/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func ResJSON(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	w.WriteHeader(status)
	render.JSON(w, r, v)
}

func ParseReqBody(r *http.Request, req any) error {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(reqBody, &req); err != nil {
		logger.WithContext(r.Context()).Error("failed decode request body", zap.Error(err))
		return err
	}
	return nil
}

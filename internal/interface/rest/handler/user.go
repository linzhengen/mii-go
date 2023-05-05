package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/linzhengen/mii-go/internal/interface/rest"
	"github.com/linzhengen/mii-go/internal/interface/rest/request"
	"github.com/linzhengen/mii-go/internal/interface/rest/response"
	"github.com/linzhengen/mii-go/internal/usecase"
	"github.com/linzhengen/mii-go/pkg/logger"
	"net/http"
)

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{userUseCase: userUseCase}
}

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func (h userHandler) Get(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	u, err := h.userUseCase.GetUser(r.Context(), userId)
	if err != nil {
		logger.WithContext(r.Context()).Errorf("get user error: %v", err)
		rest.ResJSON(w, r, http.StatusBadRequest, nil)
		return
	}
	rest.ResJSON(w, r, http.StatusOK, response.GetUserRes{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Status: string(u.Status),
	})
}

func (h userHandler) Post(w http.ResponseWriter, r *http.Request) {
	var req request.PostUserReq
	if err := rest.ParseReqBody(r, &req); err != nil {
		rest.ResJSON(w, r, http.StatusBadRequest, nil)
		return
	}
	if err := h.userUseCase.CreateUser(r.Context(), req.Name, req.Password, req.Email); err != nil {
		logger.WithContext(r.Context()).Errorf("create user error: %v", err)
		rest.ResJSON(w, r, http.StatusInternalServerError, nil)
		return
	}
	rest.ResJSON(w, r, http.StatusOK, nil)
}

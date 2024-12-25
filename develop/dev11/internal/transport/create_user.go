package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
	"strconv"
)

type CreateUserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type createUserHandler struct {
	service service.UserService
}

func NewUserCreateHandler(s service.UserService) CreateUserHandler {
	return &createUserHandler{
		service: s,
	}
}

func (h *createUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		badResponse(w, http.StatusBadRequest, "user_id must be integer")
		return
	}

	_, err = h.service.CreateUser(r.Context(), user_id)
	if err != nil {
		badResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	successResponse(w, http.StatusOK)
}

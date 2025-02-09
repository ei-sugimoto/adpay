package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/usecase"
)

type UserController struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

func (c UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
		return
	}

	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	// リクエストボディをデコード
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	newUser := entity.NewUserWithoutID(req.Name, req.Password)
	err := c.UserUsecase.Save(r.Context(), newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintln(err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

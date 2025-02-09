package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/infra/persistence"
	"github.com/ei-sugimoto/adpay/apps/backend/usecase"
	"github.com/ei-sugimoto/adpay/apps/backend/utils"
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
		if err == persistence.ErrExistUser {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "User already exists"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintln(err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
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

	token, err := c.UserUsecase.Login(r.Context(), req.Name, req.Password)
	if err != nil {
		switch err {
		case utils.ErrFailedSignToken:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to sign token"})
		case utils.ErrNoSetSecretKey:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "JWT_SECRET_KEY is not set"})

		default:
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "name or password is incorrect"})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

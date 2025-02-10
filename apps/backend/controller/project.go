package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/usecase"
)

type ProjectController struct {
	ProjectUsecase *usecase.ProjectUsecase
}

func NewProjectController(projectUsecase *usecase.ProjectUsecase) *ProjectController {
	return &ProjectController{
		ProjectUsecase: projectUsecase,
	}
}

func (c *ProjectController) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
		return
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	userID := r.Context().Value("userID").(int64)

	newProject := entity.NewProjectWithoutID(req.Name, userID)
	err := c.ProjectUsecase.Save(r.Context(), newProject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintln(err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

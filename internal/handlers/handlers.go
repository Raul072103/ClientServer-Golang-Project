package handlers

import (
	"SCS_project/internal/config"
	"SCS_project/internal/helpers"
	"SCS_project/internal/repository"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App      *config.AppConfig
	commRepo repository.CommunicationRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		a,
		repository.NewCommRepo(a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		helpers.ServerError(w, err)
	}
}

package repository

import (
	"SCS_project/internal/config"
)

type CommunicationRepo interface {
	ConfirmCommunication() (bool, error)
}

type CommRepo struct {
	App *config.AppConfig
}

func NewCommRepo(a *config.AppConfig) CommunicationRepo {
	return &CommRepo{a}
}

func (r *CommRepo) ConfirmCommunication() (bool, error) {
	return false, nil
}

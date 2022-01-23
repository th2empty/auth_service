package service

import (
	"github.com/th2empty/auth-server/pkg/models"
	"github.com/th2empty/auth-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

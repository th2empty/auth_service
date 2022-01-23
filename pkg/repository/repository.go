package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/th2empty/auth-server/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

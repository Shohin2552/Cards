package repository

import (
	"cards/internal/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgx.Conn
}

type RepositoryInterface interface {
	CreateUser(user models.User)
	SelectUser(user models.User)
	DeleteUser(user models.User)
	UpdateUser(user models.User)
}

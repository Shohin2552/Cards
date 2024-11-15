package repository

import (
	"cards/internal/models"
	"context"
)

func (repo *Repository) CreateUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO users(name, number, create_date,
		)VALUES($1, $2, NOW())
	`, user.Name, user.Number)

	return
}

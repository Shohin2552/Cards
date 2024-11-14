package repository

import (
	"cards/internal/models"
	"context"
)

func (repo *Repository) CreateUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO users(name, number, create_date,
		)VALUES($1, $2, $3,)
	`, user.Name, user.Number, user.Create_Date)

	return
}

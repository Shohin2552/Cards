package repository

import (
	"cards/internal/models"
	"context"
)

func (repo *Repository) SelectUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		SELECT * FROM users;
	`, user.Id, user.Name, user.Number, user.Create_Date)

	return
}

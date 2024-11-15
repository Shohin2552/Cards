package repository

import (
	"cards/internal/models"
	"context"
)

func (repo *Repository) UpdateUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		UPDATE users
		SET
			name = $2,
			number = $3,
		WHERE id = $1
	`, user.Id, user.Name, user.Number)

	return
}

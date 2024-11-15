package repository

import (
	"cards/internal/models"
	"context"
)

func (repo *Repository) DeleteUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		DELETE * FROM users;
		WHERE id = $1
	`, user.Id)

	return
}

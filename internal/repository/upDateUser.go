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
			create_date = $4,
		WHERE id = $1
	`, user.Id, user.Name, user.Number, user.Create_Date)

	return
}

package repository

import (
	"database/sql"
	"errors"

	"github.com/saparaly/forum/models"
)

type TokenRepo struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

type Token interface {
	GetUserIDByToken(token string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}

func (r *TokenRepo) GetUserIDByToken(token string) (*models.User, error) {
	stmt := `SELECT u.id, u.role, u.email, u.username, u.password FROM session s JOIN users u ON s.userId = u.id WHERE s.Token = ?`

	row := r.db.QueryRow(stmt, token)
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Role, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return user, nil
}

func (r *TokenRepo) GetUserById(id int) (*models.User, error) {
	stmt, err := r.db.Prepare("SELECT role, username FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.Role, &user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

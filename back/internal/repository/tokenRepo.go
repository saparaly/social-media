package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/saparaly/forum/models"
)

type TokenRepo struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

type Token interface {
	GetUserByToken(token string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}

func (r *TokenRepo) GetUserByToken(token string) (*models.User, error) {
	var followers, following sql.NullString

	stmt := `SELECT u.id, u.role, u.email, u.username, u.password, u.followers, u.following FROM session s JOIN users u ON s.userId = u.id WHERE s.Token = ?`

	row := r.db.QueryRow(stmt, token)
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Role, &user.Email, &user.Username, &user.Password, &followers, &following)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	if followers.Valid {
		user.Followers = stringsToInts(strings.Split(followers.String, ","))
	}
	if following.Valid {
		user.Following = stringsToInts(strings.Split(following.String, ","))
	}
	return user, nil
}

func (r *TokenRepo) GetUserById(id int) (*models.User, error) {
	// stmt, err := r.db.Prepare("SELECT role, username FROM users WHERE id = ?")
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// var user models.User
	// err = stmt.QueryRow(id).Scan(&user.Role, &user.Username)
	// if err != nil {
	// 	return nil, err
	// }

	// return &user, nil
	var user models.User
	var followers, following sql.NullString
	err := r.db.QueryRow("SELECT id, username, email, password, role, followers, following FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &followers, &following)
	if err != nil {
		return &models.User{}, err
	}
	if followers.Valid {
		user.Followers = stringsToInts(strings.Split(followers.String, ","))
	}
	if following.Valid {
		user.Following = stringsToInts(strings.Split(following.String, ","))
	}
	return &user, nil
}

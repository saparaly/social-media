package repository

import (
	"database/sql"
	"fmt"

	"github.com/saparaly/forum/models"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateSession(session models.Session) error
	DeleteSession(token string) error
}

func (r *AuthRepo) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (role, username, email, password) VALUES (?, ?, ?, ?) RETURNING id")
	row := r.db.QueryRow(query, user.Role, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepo) GetUserByUsername(username string) (models.User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	row := r.db.QueryRow(query, username)

	var fullUser models.User
	err := row.Scan(&fullUser.Id, &fullUser.Role, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return models.User{}, err
	}
	return fullUser, nil
}

func (r *AuthRepo) GetUserByEmail(email string) (models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	var fullUser models.User
	err := row.Scan(&fullUser.Id, &fullUser.Role, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return models.User{}, err
	}

	return fullUser, nil
}

func (r *AuthRepo) CreateSession(session models.Session) error {
	exists, err := r.getSessionByUserId(session.UserId)
	if err != nil {
		fmt.Println("session exists")
		return err
	}
	if exists != nil {
		getId, err := r.db.Query("SELECT Id from session WHERE UserID = ?", session.UserId)
		var sesId int
		for getId.Next() {
			var id int
			getId.Scan(&id)
			sesId = id
		}

		stmt, err := r.db.Prepare("UPDATE session Set Token = ?, ExpirationDate = ? WHERE ID = ?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(session.Token, session.ExpirationDate, sesId)
		if err != nil {
			return err
		}
		return nil
	}
	stmt, err := r.db.Prepare("INSERT INTO session(UserID, Token, ExpirationDate) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.UserId, session.Token, session.ExpirationDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) getSessionByUserId(userId int) (*models.Session, error) {
	row := r.db.QueryRow("SELECT * FROM session WHERE UserID = ? ORDER BY ExpirationDate DESC LIMIT 1", userId)

	var session models.Session
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpirationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func (r *AuthRepo) DeleteSession(token string) error {
	_, err := r.db.Exec("DELETE FROM session WHERE token = $1", token)
	if err != nil {
		return err
	}
	return nil
}

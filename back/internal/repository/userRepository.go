package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/saparaly/forum/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

type User interface {
	GetUsers() ([]models.User, error)
	AddUser(username string, currectUserId int) error
	GetUserIdByUsername(username string) (int, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(user models.User) error
}

func (r *UserRepo) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, username, email, password, role, followers, following FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var followers, following sql.NullString
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &followers, &following); err != nil {
			return nil, err
		}
		if followers.Valid {
			user.Followers = stringsToInts(strings.Split(followers.String, ","))
		}
		if following.Valid {
			user.Following = stringsToInts(strings.Split(following.String, ","))
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func stringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func (r *UserRepo) AddUser(toFollowId string, currectUserId int) error {
	return nil
}

func (r *UserRepo) GetUserIdByUsername(username string) (int, error) {
	var userID int
	err := r.db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *UserRepo) GetUserById(id int) (models.User, error) {
	var user models.User
	var followers, following sql.NullString
	err := r.db.QueryRow("SELECT id, username, email, password, role, followers, following FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &followers, &following)
	if err != nil {
		return models.User{}, err
	}
	if followers.Valid {
		user.Followers = stringsToInts(strings.Split(followers.String, ","))
	}
	if following.Valid {
		user.Following = stringsToInts(strings.Split(following.String, ","))
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(user models.User) error {
	// Convert the followers and following lists to comma-separated strings
	followersStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(user.Followers)), ","), "[]")
	followingStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(user.Following)), ","), "[]")

	// Update the user information in the database
	stmt, err := r.db.Prepare("UPDATE users SET username=?, email=?, password=?, role=?, followers=?, following=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Role, followersStr, followingStr, user.Id)
	if err != nil {
		return err
	}

	return nil
}

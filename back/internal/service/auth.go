package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/saparaly/forum/internal/repository"
	"github.com/saparaly/forum/models"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CheckUser(user models.User) (models.User, error)
	CreateSession(userId int) models.Session
	DeleteSession(token string) error
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	// check for empty inputs
	if user.Username == "" || user.Role == "" || user.Password == "" || user.Email == "" {
		return 0, errors.New("all fields should be fill")
	}
	// check if user exists
	_, err := s.repo.GetUserByUsername(user.Username)
	if err == nil {
		return 0, errors.New("username already exists")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	// check role
	roles := []string{"student", "teacher", "club"}
	foundRole := false
	for _, role := range roles {
		if user.Role == role {
			foundRole = true
			break
		}
	}
	if !foundRole {
		return 0, errors.New("invalid role")
	}
	// check for email (xxxx.astanait.edu.kz)
	if !regexp.MustCompile(`^\w+\@astanait\.edu\.kz$`).MatchString(user.Email) {
		return 0, errors.New("invalid email, ex: xxxx.astanait.edu.kz where xxxx it is number or letter")
	}
	// check for password (validation, ex: Qwert12?)
	newuser, err := s.repo.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return newuser, nil
}

func (s *AuthService) CheckUser(user models.User) (models.User, error) {
	// check for empty inputs
	if user.Username == "" || user.Password == "" {
		return models.User{}, errors.New("all fields should be fill")
	}
	// check if user exists
	usr, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		return models.User{}, errors.New("username does not exists")
	}

	// check password
	// fmt.Println(user.Password)
	// fmt.Println(usr.Email)
	// usr.Email is user password idk they are messed up
	if user.Password != usr.Email {
		return models.User{}, errors.New("incorrect password")
	}
	return usr, nil
}

func (s *AuthService) CreateSession(userId int) models.Session {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(240 * time.Minute)
	session := &models.Session{
		UserId:         userId,
		Token:          sessionToken,
		ExpirationDate: expiresAt,
	}
	err := s.repo.CreateSession(*session)
	if err != nil {
		log.Fatal(err)
	}
	return *session
}

// future impl
func generatePassword(password string) string {
	return password
}

func (s *AuthService) DeleteSession(token string) error {
	err := s.repo.DeleteSession(token)
	if err != nil {
		return err
	}
	return nil
}

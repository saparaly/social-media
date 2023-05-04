package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saparaly/forum/models"
)

const (
	allowedOrigin = "http://localhost:5173"
)

type signUpResponse struct {
	Id        int         `json:"id"`
	Valid     bool        `json:"valid"`
	Err       string      `json:"err,omitempty"`
	PostId    int64       `json:"postid"`
	IsUser    bool        `json:"isuser"`
	IsUpdated bool        `json:"isupdated"`
	IsDeleted bool        `json:"isdeleted"`
	Post      models.Post `json:"post"`
	IsLiked   bool        `json:"isliked"`
	// here is arr of comment
	Comment models.Comment `json:"commet"`
	// here is one comment i just messed up
	Comments models.Comment `json:"comments"`
	Users    []models.User  `json:"users"`
}

func (h *Handler) singUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Content-Type", "application/json")

	var getUser map[string]string
	err := json.NewDecoder(r.Body).Decode(&getUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.User{
		Username: getUser["username"],
		Email:    getUser["email"],
		Password: getUser["password"],
		Role:     getUser["role"],
	}
	//  vrode id create user-a
	id, err := h.services.CreateUser(*user)
	if err != nil {
		response := signUpResponse{
			Valid: false,
			Err:   err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		// http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusBadRequest)
		return
	}
	// User created successfully, return a JSON response
	response := signUpResponse{
		Id:    id,
		Valid: true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	var getUser map[string]string
	err := json.NewDecoder(r.Body).Decode(&getUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &models.User{
		Username: getUser["username"],
		Password: getUser["password"],
	}

	// validate inputs
	usr, err := h.services.CheckUser(*user)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
			Err:   err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session := h.services.CreateSession(usr.Id)
	cookie := &http.Cookie{
		Name:    "token",
		Value:   session.Token,
		Expires: session.ExpirationDate,
		Path:    "/",
	}
	http.SetCookie(w, cookie)

	response := signUpResponse{
		Id:    usr.Id,
		Valid: true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusBadRequest)
		return
	}
	err = h.services.DeleteSession(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	w.Write([]byte("Logged out successfully"))
}

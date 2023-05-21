package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/saparaly/forum/models"
)

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	currentUser, err := h.GetUserByToken(w, r)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.services.GetUsers()
	if err != nil {
		fmt.Println(err, " is this err")
		return
	}

	// fmt.Println(user)

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Id:    currentUser.Id,
		Valid: true,
		Users: user,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) followUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	currentUser, err := h.GetUserByToken(w, r)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// user to follow
	var user *models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	toFollowId, err := h.services.GetUserIdByUsername(user.Username)
	if err != nil {
		fmt.Println(err)
		return
	}

	toFollowUser, err := h.services.GetUserById(toFollowId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the user is already following the target user
	alreadyFollowing := false
	for _, id := range currentUser.Following {
		if id == toFollowId {
			alreadyFollowing = true
			break
		}
	}

	if !alreadyFollowing {
		currentUser.Following = append(currentUser.Following, toFollowId)
		err = h.services.UpdateUser(*currentUser)
		if err != nil {
			return
		}
	}

	// Check if the user is already a follower of the target user
	alreadyFollowed := false
	for _, id := range toFollowUser.Followers {
		if id == currentUser.Id {
			alreadyFollowed = true
			break
		}
	}

	if !alreadyFollowed {
		toFollowUser.Followers = append(toFollowUser.Followers, currentUser.Id)
		toFollowUser.Id = toFollowId
		err = h.services.UpdateUser(*toFollowUser)
		if err != nil {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid: true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) unfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	currentUser, err := h.GetUserByToken(w, r)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// user to unfollow
	var user *models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	toUnfollowId, err := h.services.GetUserIdByUsername(user.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	toUnfollowUser, err := h.services.GetUserById(toUnfollowId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the user is already following the target user
	followingIndex := -1
	for i, id := range currentUser.Following {
		if id == toUnfollowId {
			followingIndex = i
			break
		}
	}

	if followingIndex != -1 {
		currentUser.Following = append(currentUser.Following[:followingIndex], currentUser.Following[followingIndex+1:]...)
		err = h.services.UpdateUser(*currentUser)
		if err != nil {
			return
		}
	}

	followerIndex := -1
	for i, id := range toUnfollowUser.Followers {
		if id == currentUser.Id {
			followerIndex = i
			break
		}
	}

	if followerIndex != -1 {
		toUnfollowUser.Followers = append(toUnfollowUser.Followers[:followerIndex], toUnfollowUser.Followers[followerIndex+1:]...)
		err = h.services.UpdateUser(*toUnfollowUser)
		if err != nil {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid: true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	currentUser, err := h.GetUserByToken(w, r)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println(currentUser.Following)

	post, err := h.services.GetFollowedUsersPost(*currentUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(post)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) followedUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// user to follow
	var user *models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if len(user.Following) > 0 && user.Following[0] == 0 {
		user.Following = user.Following[1:]
	}
	fmt.Println(user.Following)
	var users []models.User

	for i := 0; i < len(user.Following); i++ {
		userOne, err := h.services.GetUserById(user.Following[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, *userOne)
	}

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid: true,
		Users: users,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	user, err := h.GetUserByToken(w, r)
	if err != nil {
		fmt.Println(err, "123")
		response := signUpResponse{
			Valid: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	userPost, err := h.services.GetUserCreatedPosts(user.Id)
	if err != nil {
		fmt.Println(err, "321")
		return
	}

	userLikedPosts, err := h.services.GetLikedPosts(user.Id)
	if err != nil {
		fmt.Println(err, "333")
		return
	}

	fmt.Println(user)
	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid:        true,
		User:         *user,
		CreatedPosts: userPost,
		LikedPosts:   userLikedPosts,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) user(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userPost, err := h.services.GetUserCreatedPosts(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	userLikedPosts, err := h.services.GetLikedPosts(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := h.services.GetUserById(id)

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid:        true,
		User:         *user,
		CreatedPosts: userPost,
		LikedPosts:   userLikedPosts,
	}
	json.NewEncoder(w).Encode(response)
}

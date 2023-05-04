package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	user, err := h.services.GetUsers()
	if err != nil {
		fmt.Println(err, " is this err")
		return
	}

	// fmt.Println(user)

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
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
		// Update the user's following list
		currentUser.Following = append(currentUser.Following, toFollowId)
		// Save the updated user information back to the database
		err = h.services.UpdateUser(*currentUser)
		if err != nil {
			return
		}
	}

	/**/
	// Check if the user is already a follower of the target user
	alreadyFollowed := false
	for _, id := range toFollowUser.Followers {
		if id == currentUser.Id {
			alreadyFollowed = true
			break
		}
	}

	if !alreadyFollowed {
		// Update the user to follow's followers list
		toFollowUser.Followers = append(toFollowUser.Followers, currentUser.Id)
		// Save the updated user information back to the database
		err = h.services.UpdateUser(*toFollowUser)
		if err != nil {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid: true,
		// PostId: postId,
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
		// Remove the user to unfollow from the following list
		fmt.Println(currentUser.Following, "1")
		currentUser.Following = append(currentUser.Following[:followingIndex], currentUser.Following[followingIndex+1:]...)
		// Save the updated user information back to the database
		fmt.Println(currentUser.Following, "2")
		err = h.services.UpdateUser(*currentUser)
		if err != nil {
			return
		}
	}

	// Check if the user is already a follower of the target user
	followerIndex := -1
	for i, id := range toUnfollowUser.Followers {
		if id == currentUser.Id {
			followerIndex = i
			break
		}
	}

	if followerIndex != -1 {
		// Remove the user to unfollow from the followers list
		toUnfollowUser.Followers = append(toUnfollowUser.Followers[:followerIndex], toUnfollowUser.Followers[followerIndex+1:]...)
		// Save the updated user information back to the database
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

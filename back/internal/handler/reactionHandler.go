package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/saparaly/forum/models"
)

func (h *Handler) postReaction(w http.ResponseWriter, r *http.Request) {
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

	userId, err := h.GetUserByToken(w, r)
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
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var post *models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// ***************************************************************************************

	if post.ReactionType == "like" {
		err = h.services.LikePost(post.Id, userId.Id)
	} else if post.ReactionType == "dislike" {
		err = h.services.DislikePost(post.Id, userId.Id)
	}
	if err != nil {
		fmt.Println("error is here")
		fmt.Println(err)
		return
	}

	p, err := h.services.GetPost(post.Id)

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid:     true,
		Post:      p,
		IsUser:    true,
		IsUpdated: true,
		IsLiked:   true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) commentReaction(w http.ResponseWriter, r *http.Request) {
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

	userId, err := h.GetUserByToken(w, r)
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
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var comment *models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// ***************************************************************************************

	if comment.ReactionType == "like" {
		err = h.services.LikeComment(comment.Id, userId.Id)
	} else if comment.ReactionType == "dislike" {
		err = h.services.DislikeComment(comment.Id, userId.Id)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := h.services.GetOneComment(comment.Id)

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid:    true,
		Comments: c,
		IsUser:   true,
		IsLiked:  true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
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

	userId, err := h.GetUserByToken(w, r)
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
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var comment *models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	comment.UserId = userId.Id
	comment.Username = userId.Username
	comment.UserRole = userId.Role

	err = h.services.CreateComment(*comment)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
			Err:   err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid:   true,
		Comment: *comment,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getComments(w http.ResponseWriter, r *http.Request) {
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

	postId, err := strconv.Atoi(r.URL.Query().Get("postId"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comments, err := h.services.GetComment(postId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

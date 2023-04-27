package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	// fmt.Println(string(body))

	var post *models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// ***************************************************************************************

	fmt.Println(post.ReactionType, "action")
	if post.ReactionType == "like" {
		// fmt.Println("liked")
		err = h.services.LikePost(post.Id, userId.Id)
	} else if post.ReactionType == "dislike" {
		fmt.Println("dislike")
		err = h.services.DislikePost(post.Id, userId.Id)
	}
	if err != nil {
		fmt.Println("error is here")
		fmt.Println(err)
		return
	}

	p, err := h.services.GetPost(post.Id)
	fmt.Println(p.Like, " final check")

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

func (h *Handler) commentReaction(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {}

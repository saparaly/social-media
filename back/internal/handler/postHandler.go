package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/saparaly/forum/models"
)

// _ httprouter.Params

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
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
			// Err:   err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
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
	post.UserId = userId.Id
	postId, err := h.services.CreatePost(*post)
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
	// w.Write([]byte(`{"valid": true}`))
	response := signUpResponse{
		Valid:  true,
		PostId: postId,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getPost(w http.ResponseWriter, r *http.Request) {
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

	post, err := h.services.GetPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
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

	if userId.Id != post.UserId {
		response := signUpResponse{
			IsUser: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.services.DeletePost(*post)
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
		Valid:     true,
		IsUser:    true,
		IsDeleted: true,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
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
	// fmt.Println("token problem")
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println("not token problem")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// fmt.Println("smth")

	var post *models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	// fmt.Println("is error here?")

	// fmt.Println(userId.Id)
	// fmt.Println(post.UserId)
	if userId.Id != post.UserId {
		response := signUpResponse{
			IsUser: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println("no, how can i finf err")

	err = h.services.UpdatePost(*post)
	if err != nil {
		fmt.Println(err)
		response := signUpResponse{
			Valid: false,
			Err:   err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println("post updated")
	// fmt.Println(post)

	w.WriteHeader(http.StatusOK)
	response := signUpResponse{
		Valid:     true,
		Post:      *post,
		IsUser:    true,
		IsUpdated: true,
	}
	json.NewEncoder(w).Encode(response)
}

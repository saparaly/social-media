package handler

import (
	"net/http"

	"github.com/saparaly/forum/internal/service"
	"github.com/saparaly/forum/models"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitPoutes() http.Handler {
	// mux := httprouter.New()
	// mux.POST("/create-post", h.createPost)
	// mux.GET("/post/:%d", h.getPost)

	mux := http.NewServeMux()
	mux.HandleFunc("/signup", h.singUp)
	mux.HandleFunc("/signin", h.signIn)
	mux.HandleFunc("/logout", h.logout)
	mux.HandleFunc("/create-post", h.createPost)
	mux.HandleFunc("/post", h.getPost)
	mux.HandleFunc("/update-post", h.updatePost)
	mux.HandleFunc("/delete-post", h.deletePost)
	mux.HandleFunc("/post-reaction", h.postReaction)
	mux.HandleFunc("/comment-reaction", h.commentReaction)
	mux.HandleFunc("create-comment", h.createComment)

	return mux
}

func (h *Handler) GetUserByToken(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return &models.User{}, err
	}

	user, err := h.services.ChechUserByToken(cookie.Value)
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
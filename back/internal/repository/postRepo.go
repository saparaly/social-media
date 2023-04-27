package repository

import (
	"database/sql"
	"strings"

	"github.com/saparaly/forum/models"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

type Post interface {
	CreatePost(post models.Post) (int64, error)
	GetPost(postId int) (models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(post models.Post) error
	LikePost(postId, userId int) error
}

func (r *PostRepo) CreatePost(post models.Post) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO post(userId, title, img, description, tags, created, date, location) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.UserId, post.Title, post.Img, post.Description, strings.Join(post.Tags, ","), post.Created, post.Date, post.Location)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostRepo) GetPost(postId int) (models.Post, error) {
	stmt, err := r.db.Prepare("SELECT id, userId, title, img, description, tags, created, date, location, like, dislike FROM post WHERE id = ?")
	if err != nil {
		return models.Post{}, err
	}
	defer stmt.Close()

	var post models.Post
	var tagsString string
	err = stmt.QueryRow(postId).Scan(&post.Id, &post.UserId, &post.Title, &post.Img, &post.Description, &tagsString, &post.Created, &post.Date, &post.Location, &post.Like, &post.Dislike)
	if err != nil {
		return models.Post{}, err
	}

	post.Tags = strings.Split(tagsString, ",")
	return post, nil
}

func (r *PostRepo) UpdatePost(post models.Post) error {
	stmt, err := r.db.Prepare("UPDATE post Set title = ?, description = ?, location = ?, date = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Description, post.Location, post.Date, post.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) DeletePost(post models.Post) error {
	stmt, err := r.db.Prepare("DELETE FROM post WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepo) LikePost(postId, userId int) error {
	// Check if user has already liked the post
	row := r.db.QueryRow("SELECT COUNT(*) FROM post_likes WHERE postId = $1 AND userId = $2", postId, userId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// User has already liked the post, remove the like
		_, err := r.db.Exec("DELETE FROM post_likes WHERE postId = $1 AND userId = $2", postId, userId)
		if err != nil {
			return err
		}
	} else {
		// User has not yet liked the post, add the like
		_, err := r.db.Exec("INSERT INTO post_likes (postId, userId) VALUES ($1, $2)", postId, userId)
		if err != nil {
			return err
		}
	}

	// Get the like count for the post
	row = r.db.QueryRow("SELECT COUNT(*) FROM post_likes WHERE postId = $1", postId)
	var likeCount int
	err = row.Scan(&likeCount)
	if err != nil {
		return err
	}

	// Remove the dislike if there is any
	_, err = r.db.Exec("DELETE FROM post_dislikes WHERE postId = $1 AND userId = $2", postId, userId)
	if err != nil {
		return err
	}

	// Update the post record with the new like and dislike counts
	_, err = r.db.Exec("UPDATE post SET like = $1, dislike = (SELECT COUNT(*) FROM post_dislikes WHERE postId = $2) WHERE id = $2", likeCount, postId)
	if err != nil {
		return err
	}

	return nil
}

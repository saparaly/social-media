package repository

import (
	"database/sql"

	"github.com/saparaly/forum/models"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

type Comment interface {
	CreateComment(comment models.Comment) error
	GetComments(postID int) ([]*models.Comment, error)
	GetOneComment(commentID int) (models.Comment, error)
}

func (r *CommentRepo) CreateComment(comment models.Comment) error {
	stmt := "INSERT INTO comments (text, userId, username, role, postId) VALUES (?, ?,?, ?, ?)"

	result, err := r.db.Exec(stmt, comment.Text, comment.UserId, comment.Username, comment.UserRole, comment.PostId)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	comment.Id = int(lastInsertId)
	return nil
}

func (r *CommentRepo) GetComments(postID int) ([]*models.Comment, error) {
	// Prepare the SQL query to retrieve all comments for a post.
	query := `
        SELECT id, postId, userId, username, role, text, like, dislike
        FROM comments
        WHERE postId = $1
    `

	// Execute the query and retrieve the results.
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the rows and build the slice of comments.
	comments := []*models.Comment{}
	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Username, &comment.UserRole, &comment.Text, &comment.Like, &comment.Dislike)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepo) GetOneComment(commentID int) (models.Comment, error) {
	// Prepare the SQL query to retrieve a single comment by ID.
	query := `
        SELECT id, postId, userId, username, role, text, like, dislike
        FROM comments
        WHERE id = $1
    `

	// Execute the query and retrieve the result.
	row := r.db.QueryRow(query, commentID)
	comment := models.Comment{}
	err := row.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Username, &comment.UserRole, &comment.Text, &comment.Like, &comment.Dislike)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/saparaly/forum/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

type User interface {
	GetUsers() ([]models.User, error)
	AddUser(username string, currectUserId int) error
	GetUserIdByUsername(username string) (int, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(user models.User) error
	GetFollowedUsersPost(currectUser models.User) ([]models.Post, error)
	getPostIdsByUserId(userId int) ([]int, error)
	GetPost(postId int) (models.Post, error)
	GetUserCreatedPosts(id int) ([]models.Post, error)
	GetLikedPosts(id int) ([]models.Post, error)
}

func (r *UserRepo) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, username, email, password, role, followers, following FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var followers, following sql.NullString
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &followers, &following); err != nil {
			return nil, err
		}
		if followers.Valid {
			user.Followers = stringsToInts(strings.Split(followers.String, ","))
		}
		if following.Valid {
			user.Following = stringsToInts(strings.Split(following.String, ","))
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func stringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func (r *UserRepo) AddUser(toFollowId string, currectUserId int) error {
	return nil
}

func (r *UserRepo) GetUserIdByUsername(username string) (int, error) {
	var userID int
	err := r.db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *UserRepo) GetUserById(id int) (models.User, error) {
	var user models.User
	var followers, following sql.NullString
	err := r.db.QueryRow("SELECT id, username, email, password, role, followers, following FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &followers, &following)
	if err != nil {
		return models.User{}, err
	}
	if followers.Valid {
		user.Followers = stringsToInts(strings.Split(followers.String, ","))
	}
	if following.Valid {
		user.Following = stringsToInts(strings.Split(following.String, ","))
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(user models.User) error {
	// Convert the followers and following lists to comma-separated strings
	followersStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(user.Followers)), ","), "[]")
	followingStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(user.Following)), ","), "[]")

	// Update the user information in the database followers
	stmt, err := r.db.Prepare("UPDATE users SET username=?, email=?,followers=?, following=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, followersStr, followingStr, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetFollowedUsersPost(currentUser models.User) ([]models.Post, error) {
	var posts []models.Post

	// postIds created by user
	// currentUser.Following [0, 1, 2]
	if len(currentUser.Following) > 0 && currentUser.Following[0] == 0 {
		currentUser.Following = append(currentUser.Following[:0], currentUser.Following[1:]...)
	}

	var postsIds []int
	for i := 0; i < len(currentUser.Following); i++ {
		postId, err := r.getPostIdsByUserId(currentUser.Following[i])
		if err != nil {
			return nil, err
		}
		postsIds = append(postsIds, postId...)
	}

	for i := 0; i < len(postsIds); i++ {
		post, err := r.GetPost(postsIds[i])
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *UserRepo) getPostIdsByUserId(userId int) ([]int, error) {
	var postIds []int

	query := "SELECT id FROM post WHERE userId = ?"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var postId int
		if err := rows.Scan(&postId); err != nil {
			return nil, err
		}
		postIds = append(postIds, postId)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return postIds, nil
}

func (r *UserRepo) GetPost(postId int) (models.Post, error) {
	stmt, err := r.db.Prepare("SELECT id, userId, username, role, title, img, description, tags, created, date, location, like, dislike FROM post WHERE id = ?")
	if err != nil {
		return models.Post{}, err
	}
	defer stmt.Close()

	var post models.Post
	var tagsString string
	err = stmt.QueryRow(postId).Scan(&post.Id, &post.UserId, &post.Username, &post.UserRole, &post.Title, &post.Img, &post.Description, &tagsString, &post.Created, &post.Date, &post.Location, &post.Like, &post.Dislike)
	if err != nil {
		return models.Post{}, err
	}

	post.Tags = strings.Split(tagsString, ",")
	return post, nil
}

func (r *UserRepo) GetUserCreatedPosts(id int) ([]models.Post, error) {
	stmt := `SELECT id, userId, username, role, title, img, description, tags, created, date, location, like, dislike
	FROM post WHERE userId = ?`

	rows, err := r.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.Post{}

	for rows.Next() {
		post := &models.Post{}
		var tagsString string
		err := rows.Scan(&post.Id, &post.UserId, &post.Username, &post.UserRole, &post.Title, &post.Img, &post.Description, &tagsString, &post.Created, &post.Date, &post.Location, &post.Like, &post.Dislike)
		if err != nil {
			return nil, err
		}
		post.Tags = strings.Split(tagsString, ",")
		posts = append(posts, *post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *UserRepo) GetLikedPosts(id int) ([]models.Post, error) {
	// Get the IDs of all posts that have been liked by the user
	rows, err := r.db.Query("SELECT postId FROM post_likes WHERE userId = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the liked posts
	likedPosts := []models.Post{}

	// Iterate through the rows and add each post to the likedPosts slice
	for rows.Next() {
		var postID int
		err := rows.Scan(&postID)
		if err != nil {
			return nil, err
		}

		// Use the Get method to get the post by its ID
		post, err := r.GetPost(postID)
		if err != nil {
			return nil, err
		}

		likedPosts = append(likedPosts, post)
	}

	// Check for any errors that occurred while iterating through the rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return likedPosts, nil
}

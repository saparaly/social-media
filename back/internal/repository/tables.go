package repository

const (
	users = `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		role TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	session = `CREATE TABLE IF NOT EXISTS session (
		id INTEGER PRIMARY KEY,
		UserID INTEGERNOT NULL,
		Token TEXT NOT NULL UNIQUE,
		ExpirationDate TIMESTAMP NOT NULL
	);`
	post = `CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY,
		userId INTEGER NOT NULL,
		title TEXT NOT NULL,
		img TEXT,
		description TEXT NOT NULL,
		tags TEXT,
		created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		date TEXT,
		location TEXT,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		FOREIGN KEY (userId) REFERENCES users(id)
	);`
	postLikes = `CREATE TABLE IF NOT EXISTS post_likes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		postId INTEGER NOT NULL,
		FOREIGN KEY (userId) REFERENCES users(id),
		FOREIGN KEY (postId) REFERENCES post(id),
		UNIQUE (userId, postId)
	);`
	postDislikes = ` CREATE TABLE IF NOT EXISTS post_dislikes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postId INTEGER,
		userId INTEGER,
		FOREIGN KEY(postId) REFERENCES post(id),
		FOREIGN KEY(userId) REFERENCES users(id),
		UNIQUE (userId, postId)
	);`
)

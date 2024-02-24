package post

import "database/sql"

type Post struct {
	ID      int
	Content string
	Author  string
	DB      *sql.DB
}

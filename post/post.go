package post

import "database/sql"

type Post struct {
	ID      int
	Content string
	Author  string
	DB      *sql.DB
}

func (p *Post) Create() error {
	_, err := p.DB.Exec("INSERT INTO posts (content, author) VALUES (?, ?)", p.Content, p.Author)
	return err

}

func (p *Post) GetPost(id int) (Post, error) {
	row := p.DB.QueryRow("SELECT id, content, author FROM posts WHERE id = ?", id)
	err := row.Scan(&p.ID, &p.Content, &p.Author)
	return *p, err
}

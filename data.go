package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user= dbname= password ssl-mode=false")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (Post, error) {
	post := Post{}

	err := Db.QueryRow(`INSERT INTO posts (content, author) VLAUES ($1)`, id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (post *Post) create() error {
	var err error
	statement := `INSERT INTO posts (content, author) VALUES ($1, $2) returning id`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return err
}

func (post *Post) update() error {
	_, err := Db.Exec("UPDATE posts SET content=$2, author=$3 WHERE id=$1", post.Id, post.Content, post.Author)
	return err
}

func (post *Post) delete() error {
	_, err := Db.Exec("DELETE FROM posts where id=$1", post.Id)
	return err
}

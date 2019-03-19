// Package main provides ...
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "os"
)

// DefaultDB 默认数据库
var DefaultDB *sql.DB

// InitDatabase init database for comment.
func InitDatabase() {
	// os.Remove("./main.db")
	db, err := sql.Open("sqlite3", "./main.db")
	DefaultDB = db
	if err != nil {
		log.Fatal(err)
	}
	// defer DefaultDB.Close()
	createUserTable()
	createCommentTable()
}

func createUserTable() {
	sql := `
    create table if not exists user (id integer primary key autoincrement, 
                                     name text default '', 
                                     uuid text default '', 
                                     email text default '',
                                     create_date datetime default current_timestamp, 
                                     update_date datetime , 
                                     blog text default '', 
                                     icon_url text default '');`
	_, err := DefaultDB.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}

func createCommentTable() {
	sql := `
    create table if not exists comment (comment_id integer primary key autoincrement,
                                        article_id text default '',
                                        parent_id text default '',
                                        forefather_id text default '',
                                        uuid text default '',
                                        content text default '',
                                        is_delete integer default 0,
                                        vote_plus integer default 0,
                                        vote_minus integer default 0);`
	_, err := DefaultDB.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}

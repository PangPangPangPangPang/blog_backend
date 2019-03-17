// Package main provides ...
package main

import (
	"database/sql"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "time"
	// "github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// DefaultDB 默认数据库
var DefaultDB *sql.DB

// InitDatabase init database for comment.
func InitDatabase() {
	os.Remove("./main.db")
	db, err := sql.Open("sqlite3", "./main.db")
	DefaultDB = db
	if err != nil {
		log.Fatal(err)
	}
	// defer DefaultDB.Close()
	createUserTable()

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
                                     icon_url text default '');
    delete from user;`
	_, err := DefaultDB.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}

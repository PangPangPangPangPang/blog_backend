// Package main provides ...
package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func register(c *gin.Context) {
	name, check := CheckPostParamsValid(c, "name", "Invalid user name.")
	if !check {
		return
	}

	// Check user exist.
	query := fmt.Sprintf("select name from user where name = '%s'", name)
	rows, err := DefaultDB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	exist := false
	for rows.Next() {
		var n string
		err := rows.Scan(&n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n)
		if n == name {
			exist = true
		}
	}
	if exist {
		c.JSON(http.StatusOK, gin.H{
			"result":   "",
			"errorno":  1,
			"errormsg": "User already exist.",
		})
		return
	}

	// Exec insert.
	email := c.Query("email")
	blog := c.Query("blog")
	iconURL := c.Query("icon_url")
	updateDate := c.Query("update_date")
	uuid, _ := generateUserID()
	insert := fmt.Sprintf("insert into user(name, uuid, email, blog, icon_url, update_date) values('%s', '%s', '%s', '%s', '%s', '%s')", name, uuid, email, blog, iconURL, updateDate)
	_, err = DefaultDB.Exec(insert)
	if err != nil {
		log.Fatal(err)
	}
	// rowsid, _ := res.RowsAffected()
	// res.RowsAffected()
	c.JSON(http.StatusOK, gin.H{
		"result":   uuid,
		"errorno":  0,
		"errormsg": "success",
	})
}

func generateUserID() (id string, err error) {
	var ret uuid.UUID
	temp, err := uuid.NewV4()
	ret = temp
	if err != nil {
		return ret.String(), err
	}
	return ret.String(), nil
}

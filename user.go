// Package main provides ...
package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	u "github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func upload(c *gin.Context) {

}

func uploadIcon(c *gin.Context) (string, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	filename := header.Filename

	out, err := os.Create("./avatar/" + filename)
	defer out.Close()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return filename, nil

}

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
			"result":    "",
			"errorcode": 1,
			"errormsg":  "User already exist.",
		})
		return
	}

	iconURL, err := uploadIcon(c)
	// if err != nil {
	// c.JSON(http.StatusOK, gin.H{
	// "result":    "",
	// "errorcode": 1,
	// "errormsg":  "Upload icon failed.",
	// })
	// return
	// }

	// Exec insert.
	email := c.PostForm("email")
	blog := c.PostForm("blog")
	updateDate := time.Now().UTC()
	uuid, _ := generateUserID()
	insert := fmt.Sprintf("insert into user(name, uuid, email, blog, icon_url, update_date) values('%s', '%s', '%s', '%s', '%s', '%s')", name, uuid, email, blog, iconURL, updateDate)
	_, err = DefaultDB.Exec(insert)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":    "",
			"errorcode": 1,
			"errormsg":  "Registe user failed.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": gin.H{
			"uuid":     uuid,
			"blog":     blog,
			"email":    email,
			"icon_url": iconURL,
			"name":     name,
		},
		"errorcode": 0,
		"errormsg":  "success",
	})
}

func generateUserID() (id string, err error) {
	var ret u.UUID
	temp, err := u.NewV4()
	ret = temp
	if err != nil {
		return ret.String(), err
	}
	return ret.String(), nil
}

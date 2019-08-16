// Package main provides ...
package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	u "github.com/satori/go.uuid"
)

func upload(c *gin.Context) {

}

func uploadIcon(c *gin.Context) (string, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	filename := header.Filename
	avatarPath := VolumnPath("avatar")
	if _, err := os.Stat(avatarPath); os.IsNotExist(err) {
		os.Mkdir(avatarPath, os.ModePerm)
	}

	out, err := os.Create(avatarPath + "/" + filename)
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
	query := fmt.Sprintf("select name, email, uuid from user where name = '%s'", name)
	rows, err := DefaultDB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	exist := false
	var fetchEmail string
	var fetchUUID string
	for rows.Next() {
		var n string
		var e string
		var u string
		err := rows.Scan(&n, &e, &u)
		if err != nil {
			log.Fatal(err)
		}
		if n == name {
			fetchEmail = e
			fetchUUID = u
			exist = true
		}
	}
	email := c.PostForm("email")
	if exist {
		// Login if 'name' and 'email' matched.
		if fetchEmail == email {
			c.JSON(http.StatusOK, gin.H{
				"result": gin.H{
					"uuid":  fetchUUID,
					"email": fetchEmail,
				},
				"errorcode": 0,
				"errormsg":  "Login success.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result":    "",
				"errorcode": 1,
				"errormsg":  "Account already exist.",
			})
		}
		return
	}

	iconURL, err := uploadIcon(c)

	// Exec insert.
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
	temp := u.NewV4()
	ret = temp
	return ret.String(), nil
}

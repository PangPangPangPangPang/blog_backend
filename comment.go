// Package main provides ...
package main

import (
	"github.com/gin-gonic/gin"

	// "log"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

// Comment comment
type Comment struct {
	CommentID    int    `json:"comment_id"`
	ArticleID    string `json:"article_id"`
	ParentID     int    `json:"parent_id"`
	ForefatherID string `json:"forefather_id"`
	UUID         string `json:"uuid"`
	Content      string `json:"content"`
	Name         string `json:"name"`
	Blog         string `json:"blog"`
	IconURL      string `json:"icon_url"`
	CreateDate   int64  `json:"create_date"`
	IDDelete     int    `json:"is_delete"`
	VotePlus     int    `json:"vote_plus"`
	VoteMinus    int    `json:"vote_minus"`
}

// Comments  comment list
type Comments []Comment

// AddComment for add comment
func AddComment(c *gin.Context) {
	articleID, check := CheckPostParamsValid(c, "article_id", "Invalid article.")
	if !check {
		return
	}
	uuid, check := CheckPostParamsValid(c, "uuid", "Invalid user.")
	if !check {
		return
	}
	content, check := CheckPostParamsValid(c, "content", "Invalid content.")
	if !check {
		return
	}

	now := time.Now().Unix()
	parentID := c.PostForm("parent_id")
	forefatherID := c.PostForm("forefather_id")
	insert := fmt.Sprintf(`insert into comment
                           (content, 
                           create_date,
                           article_id,
                           parent_id,
                           forefather_id,
                           uuid) 
                           values('%s', '%d', '%s', '%s', '%s', '%s')`,
		content, now, articleID, parentID, forefatherID, uuid)
	res, err := DefaultDB.Exec(insert)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":    err,
			"errorcode": 1,
			"errormsg":  "Insert comment failed.",
		})
		return
	}
	rowsid, _ := res.RowsAffected()
	res.RowsAffected()
	c.JSON(http.StatusOK, gin.H{
		"result":    rowsid,
		"errorcode": 0,
		"errormsg":  "success",
	})

}

// FetchComment fetch comment
func FetchComment(c *gin.Context) {
	articleID, check := CheckGetParamsValid(c, "article_id", "Invalid article.")
	if !check {
		return
	}
	fetch := fmt.Sprintf(`SELECT comment.*, user.icon_url, user.name, user.blog
	                      FROM comment
						  LEFT OUTER JOIN user on comment.uuid = user.uuid 
						  WHERE comment.article_id = '%s'`, articleID)
	rows, err := DefaultDB.Query(fetch)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":    err,
			"errorcode": 1,
			"errormsg":  "Fetch comment failed.",
		})
		return
	}
	var list Comments
	for rows.Next() {
		var commentID int
		var articleID string
		var parentID int
		var forefatherID string
		var uuid string
		var content string
		var createDate int64
		var idDelete int
		var votePlus int
		var voteMinus int
		var iconURL sql.NullString
		var name sql.NullString
		var blog sql.NullString
		err := rows.Scan(
			&commentID,
			&articleID,
			&parentID,
			&forefatherID,
			&uuid,
			&content,
			&createDate,
			&idDelete,
			&votePlus,
			&voteMinus,
			&iconURL,
			&name,
			&blog,
		)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"result":    err,
				"errorcode": 1,
				"errormsg":  "Fetch comment failed.",
			})
			return
		}
		var comment = Comment{
			commentID,
			articleID,
			parentID,
			forefatherID,
			uuid,
			content,
			name.String,
			blog.String,
			iconURL.String,
			createDate,
			idDelete,
			votePlus,
			voteMinus,
		}

		list = append(list, comment)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": gin.H{
			"comments":   list,
			"article_id": articleID,
		},
		"errorcode": 0,
		"errormsg":  "Success",
	})
	return
}

// DeleteComment fetch comment
func DeleteComment(c *gin.Context) {
	commentID, check := CheckPostParamsValid(c, "comment_id", "Invalid comment_id.")
	if !check {
		return
	}
	uuid, check := CheckPostParamsValid(c, "uuid", "Invalid uuid.")
	if !check {
		return
	}
	update := fmt.Sprintf(`update comment 
                           set is_delete = 1 
                           where comment_id = '%s' and uuid = '%s'`,
		commentID, uuid)
	_, err := DefaultDB.Exec(update)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":    err,
			"errorcode": 1,
			"errormsg":  "Fetch comment failed.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":    "",
		"errorcode": 0,
		"errormsg":  "Success",
	})
	return
}

// CheckPostParamsValid check if valid
func CheckPostParamsValid(c *gin.Context, param string, errormsg string) (string, bool) {
	p := c.PostForm(param)
	if p == "" || p == "undefined" {
		c.JSON(http.StatusOK, gin.H{
			"result":    "",
			"errorcode": 1,
			"errormsg":  errormsg,
		})
		return "", false
	}
	return p, true
}

// CheckGetParamsValid check if valid
func CheckGetParamsValid(c *gin.Context, param string, errormsg string) (string, bool) {
	p := c.Query(param)
	if p == "" || p == "undefined" {
		c.JSON(http.StatusOK, gin.H{
			"result":    "",
			"errorcode": 1,
			"errormsg":  errormsg,
		})
		return "", false
	}
	return p, true
}

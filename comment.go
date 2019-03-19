// Package main provides ...
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

// Comment comment
type Comment struct {
	CommentID    int    `json:"comment_id"`
	ArticleID    string `json:"article_id"`
	ParentID     string `json:"parent_id"`
	ForefatherID string `json:"forefather_id"`
	UUID         string `json:"uuid"`
	Content      string `json:"content"`
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

	parentID := c.PostForm("parent_id")
	forefatherID := c.PostForm("forefather_id")
	insert := fmt.Sprintf(`insert into comment
                           (content, 
                           article_id,
                           parent_id,
                           forefather_id,
                           uuid) 
                           values('%s', '%s', '%s', '%s', '%s')`,
		content, articleID, parentID, forefatherID, uuid)
	res, err := DefaultDB.Exec(insert)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":   err,
			"errorno":  1,
			"errormsg": "Insert comment failed.",
		})
		return
	}
	rowsid, _ := res.RowsAffected()
	res.RowsAffected()
	c.JSON(http.StatusOK, gin.H{
		"result":   rowsid,
		"errorno":  0,
		"errormsg": "success",
	})

}

// FetchComment fetch comment
func FetchComment(c *gin.Context) {
	articleID, check := CheckGetParamsValid(c, "article_id", "Invalid article.")
	if !check {
		return
	}

	fetch := fmt.Sprintf(`select * 
                          from comment 
                          where article_id = '%s' and is_delete = 0`, articleID)
	rows, err := DefaultDB.Query(fetch)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":   err,
			"errorno":  1,
			"errormsg": "Fetch comment failed.",
		})
		return
	}
	var list Comments
	for rows.Next() {
		var commentID int
		var articleID string
		var parentID string
		var forefatherID string
		var uuid string
		var content string
		var idDelete int
		var votePlus int
		var voteMinus int
		err := rows.Scan(
			&commentID,
			&articleID,
			&parentID,
			&forefatherID,
			&uuid,
			&content,
			&idDelete,
			&votePlus,
			&voteMinus,
		)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result":   err,
				"errorno":  1,
				"errormsg": "Fetch comment failed.",
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
			idDelete,
			votePlus,
			voteMinus,
		}
		list = append(list, comment)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": gin.H{
			"comments": list,
		},
		"errorno":  0,
		"errormsg": "Success",
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
			"result":   err,
			"errorno":  1,
			"errormsg": "Fetch comment failed.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":   "",
		"errorno":  0,
		"errormsg": "Success",
	})
	return
}

// CheckPostParamsValid check if valid
func CheckPostParamsValid(c *gin.Context, param string, errormsg string) (string, bool) {
	p := c.PostForm(param)
	if p == "" {
		c.JSON(http.StatusOK, gin.H{
			"result":   "",
			"errorno":  1,
			"errormsg": errormsg,
		})
		return "", false
	}
	return p, true
}

// CheckGetParamsValid check if valid
func CheckGetParamsValid(c *gin.Context, param string, errormsg string) (string, bool) {
	p := c.Query(param)
	if p == "" {
		c.JSON(http.StatusOK, gin.H{
			"result":   "",
			"errorno":  1,
			"errormsg": errormsg,
		})
		return "", false
	}
	return p, true
}

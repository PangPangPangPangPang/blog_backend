package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// Article 获取文章
func Article(c *gin.Context) {

	CheckUpdate()

	id := c.Query("id")
	article := Articles[id]

	if article.Content == "" {
		articlesPath := WorkPath("articles")
		path := articlesPath + "/" + id + ".md"
		content, _ := ReadFile(path)
		article.Content = content
	}

	j, _ := json.Marshal(&article)

	c.JSON(http.StatusOK, gin.H{
		"result":    string(j),
		"errorcode": 0,
		"errormsg":  ""})
}

// ReadFile read  file
func ReadFile(path string) (string, error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

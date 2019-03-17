package main

import (
	"encoding/json"
	"fmt"
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
		content, _ := readFile(path)
		article.Content = content
	}

	j, _ := json.Marshal(&article)

	fmt.Println(j)
	c.JSON(http.StatusOK, gin.H{
		"result":   string(j),
		"errorno":  0,
		"errormsg": ""})
}

func readFile(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

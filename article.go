package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
    "net/http"
    "io/ioutil"
)

// Article aaa
func Article(c *gin.Context) {
	if ListInitStatus {
        ListInitStatus = false
		list, m := GenerateList()
		Articles = m
		ob, error := json.Marshal(&list)
		if error != nil {
			fmt.Println(error)
        }
		listJSON = string(ob)
	}

    id := c.Query("id")
    article := Articles[id]

    if article.Content == "" {
        articlesPath := WorkPath("articles")
        path := articlesPath + "/" + id + ".md"
        content, _ := readFile(path)
        article.Content = content
    }

    j, _ := json.Marshal(&article)

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

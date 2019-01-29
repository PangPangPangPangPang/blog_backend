package main

import (
	"github.com/gin-gonic/gin"
    "net/http"
)

//List 获取文章列表
func List(c *gin.Context) {
    CheckUpdate()
	c.JSON(http.StatusOK, gin.H{
		"result": ListJSON,
		"errorno": 0,
		"errormsg": ""})
}


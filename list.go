package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//List 获取文章列表
func List(c *gin.Context) {
	CheckUpdate()
	c.JSON(http.StatusOK, gin.H{
		"result":    ListJSON,
		"errorcode": 0,
		"errormsg":  ""})
}

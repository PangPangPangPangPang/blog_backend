package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init  database for comment.
	InitDatabase()

	// Run http service.
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/list", List)
	router.POST("/register", register)
	router.POST("/upload", upload)
	router.POST("/addcomment", AddComment)
	router.POST("/deletecomment", DeleteComment)
	router.GET("/fetchcomment", FetchComment)
	router.GET("/article", Article)
	router.GET("/update/:secretkey", Update)
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	router.Use(static.Serve("/avatar", static.LocalFile("./avatar", true)))

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

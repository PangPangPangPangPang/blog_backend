package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func main() {
	router := gin.New()
    router.Use(gin.Logger())
    router.GET("/list", List)
    router.GET("/article", Article)
    router.Use(static.Serve("/", static.LocalFile("./static", true)))

    router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

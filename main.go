package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
	})
}

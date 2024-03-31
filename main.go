package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")

	})

	r.GET("/image", func(ctx *gin.Context) {
		if currentImage == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		currentImage.Write(ctx.Writer)

	})

	r.GET("/thumbnail", func(ctx *gin.Context) {
		if currentImage == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		t, err := imageupload.ThumbnailJPEG(currentImage, 150, 150, 1)
		if err != nil {
			panic(err)
		}

		t.Write(ctx.Writer)
	})

	r.POST("/upload", func(ctx *gin.Context) {
		img, err := imageupload.Process(ctx.Request, "file")
		if err != nil {
			panic(err)
		}

		currentImage = img

		ctx.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":3000")

}

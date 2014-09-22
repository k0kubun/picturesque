package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	locals := gin.H{
		"url": "https://github.com/k0kubun/picturesque",
	}
	c.HTML(200, "/home/k0kubun/picturesque/views/index.html", locals) // FIXME
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Simple security
	basePath := getEnv("KEY", "k0kubun")

	r := gin.Default()
	r.LoadHTMLTemplates("views/*")
	r.GET("/", showIndex)
	r.GET("/:image_path", showImage)
	r.POST("/"+basePath+"/twitpic", uploadTwitpic)

	r.Run(":" + getEnv("PORT", "5000"))
}

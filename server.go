package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	locals := gin.H{
		"url": "https://github.com/k0kubun/picturesque",
	}
	c.HTML(200, "index.html", locals)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.Default()
	r.LoadHTMLTemplates("views/*")
	r.GET("/", showIndex)
	r.POST("/twitpic", uploadTwitpic)

	r.Run(":" + getEnv("PORT", "3000"))
}

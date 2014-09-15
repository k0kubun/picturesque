package main

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

var imageDir = getEnv("IMAGE_DIR", "/tmp/")

// This action should be served by reverse proxy
func showImage(c *gin.Context) {
	path := c.Params.ByName("image_path")
	c.File(imageDir + path)
}

func saveImage(file multipart.File, path string) error {
	dst, err := os.Create(imageDir + path)
	defer dst.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}
	return nil
}

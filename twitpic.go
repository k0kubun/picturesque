package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type image struct {
	Id        int    `xml:"id"`
	Text      string `xml:"text"`
	Url       string `xml:"url"`
	Width     int    `xml:"width"`
	Height    int    `xml:"height"`
	Size      int    `xml:"size"`
	Type      string `xml:"type"`
	Timestamp string `xml:"timestamp"`
	User      user   `xml:"user"`
}

type user struct {
	Id         int    `xml:"id"`
	ScreenName string `xml:"screen_name"`
}

// For YoruFukurou image upload (only xml supported)
// API specification: http://dev.twitpic.com/docs/2/upload/
const uploadMaxBytes = 5 * 1024 * 1024

func uploadTwitpic(c *gin.Context) {
	req := c.Request
	err := req.ParseMultipartForm(uploadMaxBytes)
	if err != nil {
		c.String(404, "File size is too big")
		return
	}

	files := req.MultipartForm.File["media"]
	if len(files) < 1 {
		c.String(404, "File is not uploaded")
		return
	}

	file, err := files[0].Open()
	defer file.Close()
	if err != nil {
		c.String(404, "Uploaded file open error")
		return
	}

	hash, _ := randomString(15)
	ext := fileExtension(files[0].Filename)
	path := hash + "." + ext

	err = saveImage(file, path)
	if err != nil {
		c.String(404, "Server storage is full")
		return
	}

	// FIXME: return valid params
	c.XML(200, image{
		Id:        0,
		Text:      "text",
		Url:       "http://pic.k0kubun.com/" + path, // FIXME
		Width:     0,
		Height:    0,
		Size:      0,
		Type:      ext,
		Timestamp: "Wed, 05 May 2010 16:11:15 +0000",
		User: user{
			Id:         1,
			ScreenName: "k0kubun",
		},
	})
}

func fileExtension(filename string) string {
	texts := strings.Split(filename, ".")
	if len(texts) == 0 {
		return ""
	}
	return texts[len(texts)-1]
}

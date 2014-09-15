package main

import "github.com/gin-gonic/gin"

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

func uploadTwitpic(c *gin.Context) {
	c.XML(200, image{
		Id:        1,
		Text:      "test",
		Url:       "http://i.gyazo.com/cfe106ee557900c6cbbc206177913a55.png",
		Width:     651,
		Height:    420,
		Size:      192839,
		Type:      "png",
		Timestamp: "Wed, 05 May 2010 16:11:15 +0000",
		User: user{
			Id:         1,
			ScreenName: "k0kubun",
		},
	})
}

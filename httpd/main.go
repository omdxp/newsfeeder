package main

import (
	"fmt"

	"github.com/Omar-Belghaouti/newsfeeder/platform/newsfeed"
)

func main() {
	// r := gin.Default()

	// r.GET("/ping", handler.PingGet())

	// r.Run(":8080")

	feed := newsfeed.New()

	feed.Add(newsfeed.Item{
		Title: "First item",
		Post:  "This is the first post",
	})

	fmt.Println(feed.GetAll())

}

package handler

import (
	"net/http"

	"github.com/Omar-Belghaouti/newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title" binding:"required"`
	Post  string `json:"post" binding:"required"`
}

func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		err := c.Bind(&requestBody)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}

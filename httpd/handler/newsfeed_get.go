package handler

import (
	"net/http"

	"github.com/Omar-Belghaouti/newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}

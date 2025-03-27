package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllQuotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

func FilteredQuotes(c *gin.Context) {
	// episode_id := c.Param("episode_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

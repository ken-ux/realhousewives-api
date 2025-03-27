package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllEpisodes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

func OneEpisode(c *gin.Context) {
	// episode_id := c.Param("episode_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

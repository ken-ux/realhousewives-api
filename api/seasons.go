package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllSeasons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

func OneSeason(c *gin.Context) {
	// season_id := c.Param("season_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

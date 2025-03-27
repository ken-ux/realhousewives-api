package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllSeries(c *gin.Context) {
	abbr := c.Query("abbr")

	if abbr != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "to do",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

func OneSeries(c *gin.Context) {
	// series_id := c.Param("series_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

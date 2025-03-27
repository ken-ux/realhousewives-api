package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllHousewives(c *gin.Context) {
	// limit := c.DefaultQuery("limit", 20)
	// name := c.Query("name")
	// series := c.Query("series")

	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

func OneHousewife(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

func FilteredHousewives(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "to do",
	})
}

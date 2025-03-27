package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ken-ux/realhousewives-api/api"
)

func main() {
	r := gin.Default()
	{
		v1 := r.Group("/v1")

		// Series
		v1.GET("/series/all", api.AllSeries)
		v1.GET("/series/:series_id", api.OneSeries)

		// Seasons
		v1.GET("/series/:series_id/seasons/all")
		v1.GET("/series/:series_id/seasons/:season_number")

		// Episodes
		v1.GET("/series/:series_id/seasons/:season_number/episodes/all")
		v1.GET("/series/:series_id/seasons/:season_number/episodes/:episode_number")

		// Housewives
		v1.GET("/housewives")
		v1.GET("/housewives/all")
		v1.GET("/housewives/:housewife_id")

		// Quotes
		v1.GET("/housewives/all/quotes")
		v1.GET("/housewives/:housewife_id/quotes")
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

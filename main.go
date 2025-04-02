package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ken-ux/realhousewives-api/api"
	"github.com/ken-ux/realhousewives-api/db"
	"github.com/ken-ux/realhousewives-api/middleware"
)

func main() {
	// Load dev environment.
	env := os.Getenv("ENV_NAME")
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}
	}

	// Connect to database.
	err := db.Init()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
		os.Exit(1)
	}

	defer db.Close()

	// Create router.
	r := gin.Default()

	// Enable CORS policy with all origins allowed.
	r.Use(cors.Default())
	r.Use(middleware.RateLimiterMiddleware())

	{
		v1 := r.Group("/v1")

		// Series
		v1.GET("/series/all", api.AllSeries)
		v1.GET("/series/:series_id", api.OneSeries)

		// Seasons
		v1.GET("/series/:series_id/seasons/all", api.AllSeasons)
		v1.GET("/series/:series_id/seasons/:season_number", api.OneSeason)

		// Episodes
		v1.GET("/series/:series_id/seasons/:season_number/episodes/all", api.AllEpisodes)
		v1.GET("/series/:series_id/seasons/:season_number/episodes/:episode_number", api.OneEpisode)

		// Housewives
		v1.GET("/housewives/all", api.AllHousewives)
		v1.GET("/housewives/:housewife_id", api.OneHousewife)

		// Quotes
		v1.GET("/housewives/all/quotes", api.AllQuotes)
		v1.GET("/housewives/:housewife_id/quotes", api.FilteredQuotes)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

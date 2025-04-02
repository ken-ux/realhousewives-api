package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ken-ux/realhousewives-api/db"
	"github.com/ken-ux/realhousewives-api/defs"
)

func AllEpisodes(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number, err := strconv.Atoi(c.Param("season_number"))
	if err != nil {
		c.Error(err)
		return
	}
	query := fmt.Sprintf(
		`SELECT *
		FROM episodes
		WHERE season_id = (
			SELECT season_id
			FROM seasons
			WHERE UPPER(series_id) = UPPER('%s')
				AND season_number = %d
		)`,
		series_id, season_number)

	data, err := queryEpisodes(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func OneEpisode(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number, err := strconv.Atoi(c.Param("season_number"))
	if err != nil {
		c.Error(err)
		return
	}

	episode_number, err := strconv.Atoi(c.Param("episode_number"))
	if err != nil {
		c.Error(err)
		return
	}

	query := fmt.Sprintf(
		`SELECT *
		FROM episodes
		WHERE season_id = (
			SELECT season_id
			FROM seasons
			WHERE UPPER(series_id) = UPPER('%s')
				AND season_number = %d
		)
			AND episode_number = %d`,
		series_id, season_number, episode_number)

	data, err := queryEpisodes(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func queryEpisodes(c *gin.Context, query string) ([]defs.Episode, error) {
	rows, err := db.Pool.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[defs.Episode])
}

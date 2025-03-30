package api

import (
	"fmt"
	"net/http"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ken-ux/realhousewives-api/db"
	"github.com/ken-ux/realhousewives-api/defs"
)

func AllSeasons(c *gin.Context) {
	series_id := strings.ToUpper(c.Param("series_id"))
	query := fmt.Sprintf(`SELECT * FROM seasons WHERE series_id = '%v'`, series_id)

	data, err := querySeasons(c, query)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("unable to query seasons: %v", err))
		return
	}

	c.JSON(http.StatusOK, data)
}

func OneSeason(c *gin.Context) {
	series_id := strings.ToUpper(c.Param("series_id"))
	season_number, err := strconv.Atoi(c.Param("season_number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to query seasons: non-number input for season number")
		return
	}
	query := fmt.Sprintf(`SELECT * FROM seasons WHERE series_id = '%s' AND season_number = %d`, series_id, season_number)

	data, err := querySeasons(c, query)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("unable to query seasons: %v", err))
		return
	}

	c.JSON(http.StatusOK, data)
}

func querySeasons(c *gin.Context, query string) ([]defs.Season, error) {
	rows, err := db.Pool.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[defs.Season])
}

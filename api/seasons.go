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

func AllSeasons(c *gin.Context) {
	series_id := c.Param("series_id")
	query := fmt.Sprintf(`SELECT * FROM seasons WHERE UPPER(series_id) = UPPER('%s')`, series_id)

	data, err := querySeasons(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func OneSeason(c *gin.Context) {
	series_id := c.Param("series_id")
	season_number, err := strconv.Atoi(c.Param("season_number"))
	if err != nil {
		c.Error(fmt.Errorf("non-integer input"))
		return
	}
	query := fmt.Sprintf(`SELECT * FROM seasons WHERE UPPER(series_id) = UPPER('%s') AND season_number = %d`, series_id, season_number)

	data, err := querySeasons(c, query)
	if err != nil {
		c.Error(err)
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

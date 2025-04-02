package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ken-ux/realhousewives-api/db"
	"github.com/ken-ux/realhousewives-api/defs"
)

func AllSeries(c *gin.Context) {
	query := `SELECT * FROM series`
	data, err := querySeries(c, query)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func OneSeries(c *gin.Context) {
	series_id := c.Param("series_id")
	query := fmt.Sprintf(`SELECT * FROM series WHERE UPPER(series_id) = UPPER('%s')`, series_id)

	data, err := querySeries(c, query)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func querySeries(c *gin.Context, query string) ([]defs.Series, error) {
	rows, err := db.Pool.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[defs.Series])
}

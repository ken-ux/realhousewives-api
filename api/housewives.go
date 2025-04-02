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

func AllHousewives(c *gin.Context) {
	query := `SELECT * FROM housewives`

	// Check for any query string parameters
	if len(c.Request.URL.Query()) != 0 {
		queryMap := c.Request.URL.Query()
		// Check for query string paramters that are valid for filtering data
		if queryMap["first_name"] != nil || queryMap["last_name"] != nil || queryMap["series"] != nil {
			query = createComplexQueryString(c)
		}
	}

	data, err := queryHousewives(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func createComplexQueryString(c *gin.Context) string {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	series := c.Query("series")

	query := `SELECT * FROM housewives WHERE`

	// Replace initial SQL string if a series is specified since it requires multiple JOINs
	if series != "" {
		query = fmt.Sprintf(`SELECT t1.* FROM housewives t1
				LEFT JOIN seasons_housewife t2 ON t1.housewife_id = t2.housewife_id
				LEFT JOIN seasons t3 ON t2.season_id = t3.season_id
				LEFT JOIN series t4 on t3.series_id = t4.series_id
				WHERE UPPER(t4.series_id) = UPPER('%s')`, series)
	}

	if first_name != "" {
		if series != "" {
			query += fmt.Sprintf(` AND UPPER(t1.first_name) = UPPER('%s')`, first_name)
		} else {
			query += fmt.Sprintf(` UPPER(first_name) = UPPER('%s')`, first_name)
		}
	}

	if last_name != "" {
		if series != "" {
			query += fmt.Sprintf(` AND UPPER(t1.last_name) = UPPER('%s')`, last_name)
		} else {
			if first_name != "" {
				query += ` AND`
			}
			query += fmt.Sprintf(` UPPER(last_name) = UPPER('%s')`, last_name)
		}
	}

	return query
}

func OneHousewife(c *gin.Context) {
	housewife_id, err := strconv.Atoi(c.Param("housewife_id"))
	if err != nil {
		c.Error(err)
		return
	}
	query := fmt.Sprintf(`SELECT * FROM housewives WHERE housewife_id = '%d'`, housewife_id)

	data, err := queryHousewives(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func queryHousewives(c *gin.Context, query string) ([]defs.Housewife, error) {
	rows, err := db.Pool.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[defs.Housewife])
}

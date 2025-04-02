package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ken-ux/realhousewives-api/db"
	"github.com/ken-ux/realhousewives-api/defs"
)

func AllQuotes(c *gin.Context) {
	query := `SELECT t2.housewife_id, t2.first_name, t2.last_name, t1.content, t1.tagline FROM quotes t1
			LEFT JOIN housewives t2 ON t1.housewife_id = t2.housewife_id`

	// Adjust query statement if query string parameters are specified
	if len(c.Request.URL.Query()) > 0 {
		err := parseQueryStrings(c, &query)
		if err != nil {
			c.Error(err)
			return
		}
	}

	data, err := queryQuotes(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func FilteredQuotes(c *gin.Context) {
	housewife_id, err := strconv.Atoi(c.Param("housewife_id"))
	if err != nil {
		c.Error(err)
		return
	}

	query := fmt.Sprintf(`SELECT t2.housewife_id, t2.first_name, t2.last_name, t1.content, t1.tagline FROM quotes t1
			LEFT JOIN housewives t2 ON t1.housewife_id = t2.housewife_id
			WHERE t2.housewife_id = %d`, housewife_id)

	// Adjust query statement if query string parameters are specified
	if len(c.Request.URL.Query()) > 0 {
		err := parseQueryStrings(c, &query)
		if err != nil {
			c.Error(err)
			return
		}
	}

	data, err := queryQuotes(c, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func parseQueryStrings(c *gin.Context, query *string) error {
	if c.Query("tagline") != "" {
		tagline, err := strconv.ParseBool(c.Query("tagline"))
		if err != nil {
			return errors.New("non-boolean value input for tagline parameter")
		}

		// Add SQL WHERE clause if pointer query doesn't contain it already
		if !strings.Contains(*query, "WHERE") {
			*query += fmt.Sprintf(` WHERE t1.tagline = %v`, tagline)
		} else {
			*query += fmt.Sprintf(` AND t1.tagline = %v`, tagline)
		}
	}

	if c.Query("random") != "" {
		random, err := strconv.ParseBool(c.Query("random"))
		if err != nil {
			return errors.New("non-boolean value input for random parameter")
		}

		if random {
			*query += ` ORDER BY RANDOM()
					LIMIT 1`
		}
	}
	return nil
}

func queryQuotes(c *gin.Context, query string) ([]defs.Quote, error) {
	rows, err := db.Pool.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[defs.Quote])
}

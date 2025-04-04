package docs

import (
	"fmt"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func Reference(c *gin.Context) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "./openapi.yml",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "The Real Housewives API",
		},
		DarkMode: true,
	})

	if err != nil {
		c.Error(err)
		return
	}

	fmt.Fprintln(c.Writer, htmlContent)
}

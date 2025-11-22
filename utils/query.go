package utils

import "github.com/gin-gonic/gin"

func GetSearchQuery(c *gin.Context) string {
	query := c.Query("q")
	if query == "" {
		query = c.Query("query")
	}
	return query
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
)

func searchTerm(c *gin.Context) {
	field := c.Param("field")
	term := c.Param("term")
	ids := indexer.GetIndexTerm(field, term)
	documents := indexer.GetDocuments(ids)
	if len(documents) != 0 {
		c.JSON(200, documents)
	} else {
		c.JSON(404, []string{})
	}
}

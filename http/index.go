package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
)

func indexList(c *gin.Context) {
	indeces := indexer.ListIndices()
	c.JSON(200, indeces)
}

func indexGet(c *gin.Context) {
	field := c.Param("field")
	index := indexer.GetIndex(field)
	if index != nil {
		c.JSON(200, index)
	} else {
		c.JSON(404, make(map[string]string))
	}
}

func indexGetWithTerm(c *gin.Context) {
	field := c.Param("field")
	term := c.Param("term")
	index := indexer.GetIndexTerm(field, term)
	if index != nil {
		c.JSON(200, index)
	} else {
		c.JSON(404, make(map[string]string))
	}
}

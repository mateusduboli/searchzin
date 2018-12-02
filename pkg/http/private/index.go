package private

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/pkg/index"
)

func IndexList(c *gin.Context) {
	indeces := index.ListIndices()
	c.JSON(200, indeces)
}

func IndexGet(c *gin.Context) {
	field := c.Param("field")
	index := index.GetIndex(field)
	if index != nil {
		c.JSON(200, index)
	} else {
		c.JSON(404, make(map[string]string))
	}
}

func IndexGetWithTerm(c *gin.Context) {
	field := c.Param("field")
	term := c.Param("term")
	index := index.GetIndexTerm(field, term)
	if index != nil {
		c.JSON(200, index)
	} else {
		c.JSON(404, make(map[string]string))
	}
}

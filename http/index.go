package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
)

func indexList(c *gin.Context) {
	indeces := indexer.ListIndices()
	c.JSON(200, indeces)
}

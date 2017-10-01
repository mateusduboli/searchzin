package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
)

func documentsSave(c *gin.Context) {
	var document map[string]interface{}
	c.BindJSON(&document)
	indexer.IndexDocument(indexer.NewDocument(document))
	c.JSON(200, document)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
)

func documentSave(c *gin.Context) {
	var document map[string]interface{}
	c.BindJSON(&document)
	indexer.IndexDocument(indexer.NewDocument(document))
	c.JSON(200, document)
}

func documentList(c *gin.Context) {
	documents := indexer.ListDocuments()
	c.JSON(200, documents)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
)

func documentSave(c *gin.Context) {
	var rawDocument map[string]interface{}
	c.BindJSON(&rawDocument)
	document := indexer.NewDocument(rawDocument)
	indexer.IndexDocument(document)
	indexer.IndexDocumentFields(document)
	c.JSON(200, document)
}

func documentList(c *gin.Context) {
	documents := indexer.ListDocuments()
	c.JSON(200, documents)
}

func documentGet(c *gin.Context) {
	id := c.Param("id")
	document, err := indexer.GetDocument(id)
	if err == nil {
		if id != "" {
			c.JSON(200, document)
		} else {
			c.JSON(404, map[string]string{})
		}
	} else {
		c.JSON(500, err)
	}
}

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/internal/document"
	"github.com/mateusduboli/searchzin/internal/index"
)

func DocumentSave(c *gin.Context) {
	var rawDocument map[string]interface{}
	c.BindJSON(&rawDocument)

	document := document.NewDocument(rawDocument)

	index.IndexDocument(document)
	index.IndexDocumentFields(document)
	c.JSON(200, document)
}

func DocumentDelete(c *gin.Context) {
	id := c.Param("id")
	has_deleted, err := index.DeleteDocument(id)
	if err == nil {
		if has_deleted {
			c.JSON(204, nil)
		} else {
			c.JSON(404, nil)
		}
	} else {
		c.JSON(500, err)
	}
}

func DocumentList(c *gin.Context) {
	documents := index.ListDocuments()
	c.JSON(200, documents)
}

func DocumentGet(c *gin.Context) {
	id := c.Param("id")
	document, err := index.GetDocument(id)
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

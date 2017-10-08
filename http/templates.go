package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
	"net/http"
)

func index(c *gin.Context) {
	documents := indexer.ListDocuments()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Documents": documents,
	})
}

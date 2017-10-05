package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/indexer"
	"log"
	"net/http"
)

func index(c *gin.Context) {
	documents := indexer.ListDocuments()
	log.Printf("Documents: [%s]", documents)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Documents": documents,
	})
}

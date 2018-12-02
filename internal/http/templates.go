package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/internal/index"
	"net/http"
)

func indexPage(c *gin.Context) {
	documents := index.ListDocuments()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Documents": documents,
	})
}

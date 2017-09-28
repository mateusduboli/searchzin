package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func templates(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
}

package http

import (
	"github.com/gin-gonic/gin"
	. "github.com/mateusduboli/searchzin/internal/http/internal"
	. "github.com/mateusduboli/searchzin/internal/http/v1"
)

func Searchzin(r *gin.Engine) {
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", indexPage)

	// JSON API
	api := r.Group("api")

	// v1 Version of the API
	v1 := api.Group("v1")

	// Health and metrics
	v1.GET("/ping", Ping)

	// Document API
	v1.POST("/documents", DocumentSave)
	v1.GET("/documents", DocumentList)
	v1.GET("/documents/:id", DocumentGet)
	v1.DELETE("/documents/:id", DocumentDelete)

	// Search api
	v1.GET("/search", SearchTerm)

	// Internal structure of the database
	internal := api.Group("internal")

	// Indexes
	internal.GET("/indexes", IndexList)
	internal.GET("/indexes/:field", IndexGet)
	internal.GET("/indexes/:field/:term", IndexGetWithTerm)
}

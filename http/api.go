package http

import "github.com/gin-gonic/gin"

func Api(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexPage)

	v1 := r.Group("api/v1")
	v1.GET("/ping", ping)
	v1.POST("/documents", documentSave)
	v1.GET("/documents", documentList)
	v1.GET("/documents/:id", documentGet)
	v1.DELETE("/documents/:id", documentDelete)
	v1.GET("/indexes", indexList)
	v1.GET("/indexes/:field", indexGet)
	v1.GET("/indexes/:field/:term", indexGetWithTerm)
	v1.GET("/search", searchTerm)
}

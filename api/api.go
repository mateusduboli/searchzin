package api

import "github.com/gin-gonic/gin"

func Api(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)

	v1 := r.Group("v1")
	v1.GET("/ping", ping)
}

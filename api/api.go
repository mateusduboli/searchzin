package api

import "github.com/gin-gonic/gin"

func Api(r *gin.Engine) {
	templates(r)
	ping(r)
}

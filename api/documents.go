package api

import "github.com/gin-gonic/gin"

func documentsSave(c *gin.Context) {
	var document map[string]interface{}
	c.BindJSON(&document)
	c.JSON(200, document)
}

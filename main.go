package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/api"
)

func main() {
	r := gin.Default()
	api.Api(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

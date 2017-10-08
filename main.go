package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/http"
)

func main() {
	r := gin.Default()
	http.Api(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

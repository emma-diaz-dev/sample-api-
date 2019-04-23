package main

import (
	"github.com/audio35444/sample-api/src/api/application"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	application.MapURLs(r)
	r.Run(":8080")
}

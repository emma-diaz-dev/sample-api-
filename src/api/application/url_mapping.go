package application

import (
	"github.com/audio35444/sample-api/src/api/controller"
	"github.com/gin-gonic/gin"
)

func MapURLs(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"status": "pong",
		})
	})
	r.POST("/entity", controller.EntityPost)
	r.GET("/examples", controller.ExampleGetAll)
	r.GET("/index/:index_name/new", controller.NewIndex)
	// r.GET("/examples/:example_id")
}

package controller

import (
	"net/http"

	"github.com/audio35444/sample-api/src/api/services"
	"github.com/gin-gonic/gin"
)

func ExampleGetAll(c *gin.Context) {
	e, err := services.GetExamples()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, e)
}

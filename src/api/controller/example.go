package controller

import (
	"net/http"

	"github.com/audio35444/sample-api/src/api/services"
	"github.com/gin-gonic/gin"
)

func GetBySite(c *gin.Context) {
	siteID := c.Param("site_id")
	res, err := services.GetSiteByID(siteID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func ExampleGetAll(c *gin.Context) {
	e, err := services.GetExamples()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, e)
}

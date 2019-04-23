package controller

import (
	"fmt"
	"net/http"

	"github.com/audio35444/sample-api/src/api/domain"
	"github.com/gin-gonic/gin"
)

func EntityPost(c *gin.Context) {
	entity := &domain.Entity{}
	err := c.BindJSON(entity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("[Name: %s] [Description: %s]", entity.Name, entity.Description)
	c.JSON(http.StatusOK, entity)
}

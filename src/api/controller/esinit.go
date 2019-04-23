package controller

import (
	"errors"
	"net/http"

	"github.com/audio35444/sample-api/src/api/services"
	"github.com/gin-gonic/gin"
)

func NewIndex(c *gin.Context) {
	indexName := c.Param("index_name")
	if indexName == "" {
		err := errors.New("index_name is empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	str, err := services.NewIndex(indexName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, str)
}

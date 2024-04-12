package routes

import (
	"collector/data"
	"collector/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMetricsRoute(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection name is required"})
		return
	}

	metrics, err := data.GetMetrics(collection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"metrics": metrics})
}

func PushMetricsRoute(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection name is required"})
		return
	}

	var form map[string]interface{}

	// Bind form data
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	bsonData, err := helpers.ToBsonD(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.PushMetrics(collection, bsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metrics pushed successfully"})
}

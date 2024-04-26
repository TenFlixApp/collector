package routes

import (
	"collector/data"
	"collector/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMetricsRoute(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection name is required"})
		return
	}

	var (
		metrics []bson.M
		err     error
		filter  = c.Query("filter")
	)
	if filter != "" {
		var parsedJson map[string]interface{}
		err = json.Unmarshal([]byte(filter), &parsedJson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid filter: %s", err.Error())})
			return
		}

		bsonFilter, err := helpers.ToBsonD(parsedJson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		metrics, err = data.GetFilteredMetrics(collection, bsonFilter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		metrics, err = data.GetMetrics(collection)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if metrics == nil {
		metrics = []bson.M{}
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

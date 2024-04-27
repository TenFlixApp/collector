package routes

import (
	"collector/data"
	"collector/helpers"
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

	metrics, err := data.GetMetrics(collection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func FilterMetricsRoute(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection name is required"})
		return
	}

	filter, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid pipeline: %s", err.Error())})
		return
	}

	var bsonFilter bson.D
	err = bson.UnmarshalExtJSON(filter, true, &bsonFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unable to unmarshal filter: %s", err.Error())})
		return
	}

	metrics, err := data.GetFilteredMetrics(collection, bsonFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func AggregateMetricsRoute(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection name is required"})
		return
	}

	pipeline, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid pipeline: %s", err.Error())})
		return
	}

	var bsonPipeline bson.A
	err = bson.UnmarshalExtJSON(pipeline, true, &bsonPipeline)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unable to unmarshal pipeline: %s", err.Error())})
		return
	}

	metrics, err := data.GetAggregatedMetrics(collection, bsonPipeline)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func PushMetricsRoute(c *gin.Context) {
	collection := c.Param("collection")
	if collection == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection name is required"})
		return
	}

	var form map[string]interface{}

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

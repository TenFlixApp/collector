package routes

import (
	"collector/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MetricsRoute(c *gin.Context) {
	var form map[string]interface{}

	// Bind form data
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	err := data.PushMetrics("metrics", form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

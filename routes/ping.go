package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingRoute(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

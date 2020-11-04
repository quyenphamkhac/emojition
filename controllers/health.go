package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController struct
type HealthController struct{}

// Status method: Get server status
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

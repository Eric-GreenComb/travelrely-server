package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index HeIndexalth
func Index(c *gin.Context) {
	c.String(http.StatusOK, "ticket chain")
}

// Health Health
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

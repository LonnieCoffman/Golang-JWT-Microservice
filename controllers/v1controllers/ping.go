package v1controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary ping example
// @Description do ping
// @Tags API Health
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

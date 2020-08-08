package v1controllers

import (
	"fmt"
	"net/http"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"

	"github.com/gin-gonic/gin"
)

// ClientLogin godoc
// @Summary Client Login
// @Tags Client: Authentication
// @Accept json
// @Produce json
// @Param  Body body swagger.adminLoginBody true "JSON request body"
// @Success 200 {object} swagger.adminLogin200
// @Failure 401 {object} swagger.adminLogin401
// @Failure 403 {object} swagger.adminLogin403
// @Failure 422 {object} swagger.adminLogin422
// @Router /login [post]
func ClientLogin(c *gin.Context) {
	client := models.Client{}
	clientSession := models.ClientSession{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	// Verify email and password in database
	if status, err := client.Authenticate(); err != nil {
		c.JSON(status, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Update clientSession fields
	clientSession.ClientID = client.ID
	clientSession.RemoteAddress = c.ClientIP()

	// Create and store access token and refresh token
	if err := client.CreateTokens(&clientSession); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Return JSON response containing tokens
	c.JSON(http.StatusOK, gin.H{
		"access_token":  clientSession.AccessToken,
		"message":       "Authenticated",
		"refresh_token": clientSession.RefreshToken,
		"success":       true,
	})
}

// ClientLogout godoc
// @Summary Client Logout
// @Tags Client: Authentication
// @Security bearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} swagger.adminLogout200
// @Router /logout [post]
func ClientLogout(c *gin.Context) {
	client := *(c.Keys["client"].(*models.Client))
	clientSession := *(c.Keys["clientSession"].(*models.ClientSession))

	fmt.Println(client)
	fmt.Println(clientSession)
	// admin.DestroyTokens(&adminSession)

	// // Return JSON response
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Logged out",
	// 	"success": true,
	// })
}

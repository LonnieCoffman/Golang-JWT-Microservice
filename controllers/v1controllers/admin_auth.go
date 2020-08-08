package v1controllers

import (
	"net/http"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"

	"github.com/gin-gonic/gin"
)

// AdminLogin godoc
// @Summary Administrator Login
// @Tags Admin: Authentication
// @Accept json
// @Produce json
// @Param  Body body swagger.adminLoginBody true "JSON request body"
// @Success 200 {object} swagger.adminLogin200
// @Failure 401 {object} swagger.adminLogin401
// @Failure 403 {object} swagger.adminLogin403
// @Failure 422 {object} swagger.adminLogin422
// @Router /admin/login [post]
func AdminLogin(c *gin.Context) {
	admin := models.Admin{}
	adminSession := models.AdminSession{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	// Verify email and password in database
	if status, err := admin.Authenticate(); err != nil {
		c.JSON(status, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Update adminSession fields
	adminSession.AdminID = admin.ID
	adminSession.RemoteAddress = c.ClientIP()

	// Create and store access token and refresh token
	if err := admin.CreateTokens(&adminSession); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Return JSON response containing tokens
	c.JSON(http.StatusOK, gin.H{
		"access_token":  adminSession.AccessToken,
		"message":       "Authenticated",
		"refresh_token": adminSession.RefreshToken,
		"success":       true,
	})
}

// AdminLogout godoc
// @Summary Administrator Logout
// @Tags Admin: Authentication
// @Security bearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} swagger.adminLogout200
// @Router /admin/logout [post]
func AdminLogout(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))
	adminSession := *(c.Keys["adminSession"].(*models.AdminSession))

	admin.DestroyTokens(&adminSession)

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out",
		"success": true,
	})
}

// AdminRefresh godoc
// @Summary Administrator Refresh Token
// @Tags Admin: Authentication
// @Accept json
// @Produce json
// @Param  Body body swagger.adminRefreshBody true "JSON request body"
// @Success 200 {object} swagger.adminRefresh200
// @Failure 401 {object} swagger.adminRefresh401
// @Failure 403 {object} swagger.adminRefresh403
// @Failure 422 {object} swagger.adminRefresh422
// @Router /admin/refresh [post]
func AdminRefresh(c *gin.Context) {
	admin := models.Admin{}
	adminSession := models.AdminSession{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&adminSession); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	if status, err := admin.VerifyRefresh(&adminSession); err != nil {
		c.JSON(status, gin.H{"success": false, "message": err.Error()})
		return
	}

	adminSession.RemoteAddress = c.ClientIP()

	// Create and store access token and refresh token
	if err := admin.CreateTokens(&adminSession); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  adminSession.AccessToken,
		"message":       "Token refreshed",
		"refresh_token": adminSession.RefreshToken,
		"success":       true,
	})

}

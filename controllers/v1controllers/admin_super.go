package v1controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"

	"github.com/gin-gonic/gin"
)

// AdminGetAll godoc
// @Summary Get all administrators
// @Tags Admin: SuperAdmin
// @Produce json
// @Security bearerAuth
// @Success 200 {object} swagger.adminGetAll200
// @Failure 401 {object} swagger.adminLogin401
// @Failure 403 {object} swagger.adminLogin403
// @Router /admins [get]
func AdminGetAll(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	admins, status, err := admin.GetAllAdmins()
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_results": len(admins),
		"success":       true,
		"message":       "All admins",
		"data":          admins,
	})
}

// AdminGetByID godoc
// @Summary Get administrator by ID
// @Tags Admin: SuperAdmin
// @Produce json
// @Security bearerAuth
// @Param id path int true "Admin ID"
// @Success 200 {object} swagger.adminGetByID200
// @Failure 404 {object} swagger.adminGetByID404
// @Failure 422 {object} swagger.adminGetByID422
// @Router /admins/{id} [get]
func AdminGetByID(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	// id, err := strconv.Atoi(c.Param("id"))
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "ID in invalid format", "success": false})
		return
	}

	data, status, err := admin.GetAdminByID(id)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"success": true,
		"message": "Returned admin",
	})
}

// AdminCreate godoc
// @Summary Create a new administrator
// @Tags Admin: SuperAdmin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param  Body body swagger.adminCreateAdminBody true "JSON request body"
// @Success 201 {object} swagger.adminCreateAdmin201
// @Failure 409 {object} swagger.adminCreateAdmin409
// @Failure 422 {object} swagger.adminCreateAdmin422
// @Router /admins [post]
func AdminCreate(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))
	data := models.Admin{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	status, err := admin.CreateAdmin(&data)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	// could not nullify time values from data struct so created new struct
	type result struct {
		ID        uint64 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Role      string `json:"role"`
		Status    string `json:"status"`
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Role:      data.Role,
			Status:    data.Status,
		},
		"success": true,
		"message": "Created new admin",
	})
}

// AdminEditAdmin godoc
// @Summary Edit an administrator
// @Tags Admin: SuperAdmin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param id path int true "Admin ID"
// @Param  Body body swagger.adminEditAdminBody true "JSON request body<br><br>* All fields are optional.<br>** If field left blank it will not be updated."
// @Success 200 {object} swagger.adminEditAdmin200
// @Failure 404 {object} swagger.adminEditAdmin404
// @Failure 409 {object} swagger.adminEditAdmin409
// @Failure 422 {object} swagger.adminEditAdmin422
// @Router /admins/{id} [put]
func AdminEditAdmin(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))
	data := models.Admin{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	// Verify ID paramater is an integer and add to data struct
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "ID in invalid format", "success": false})
		return
	}
	data.ID = id

	status, err := admin.EditAdmin(&data)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	// could not nullify time values from data struct so created new struct
	type result struct {
		ID        uint64 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Role      string `json:"role"`
		Status    string `json:"status"`
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Role:      data.Role,
			Status:    data.Status,
		},
		"success": true,
		"message": "Updated admin",
	})
}

// AdminDeleteAdmin godoc
// @Summary "Soft" delete an administrator
// @Tags Admin: SuperAdmin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param id path int true "Admin ID"
// @Success 200 {object} swagger.adminDeleteAdmin200
// @Failure 404 {object} swagger.adminDeleteAdmin404
// @Failure 422 {object} swagger.adminDeleteAdmin422
// @Router /admins/{id} [delete]
func AdminDeleteAdmin(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "ID in invalid format", "success": false})
		return
	}

	data, status, err := admin.DeleteAdmin(id)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Deleted admin with ID of %v", data.ID),
	})
}

// AdminDeleteClient godoc
// @Summary "Soft" delete a client
// @Tags Admin: SuperAdmin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param id path int true "Client ID"
// @Success 200 {object} swagger.adminDeleteClient200
// @Failure 404 {object} swagger.adminDeleteClient404
// @Failure 422 {object} swagger.adminDeleteClient422
// @Router /clients/{id} [delete]
func AdminDeleteClient(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "ID in invalid format", "success": false})
		return
	}

	data, status, err := admin.DeleteClient(id)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Deleted client with ID of %v", data.ID),
	})
}

package v1controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"

	"github.com/gin-gonic/gin"
)

// AdminGetSelf godoc
// @Summary Get own details
// @Tags Admin
// @Produce json
// @Security bearerAuth
// @Success 200 {object} swagger.adminGetSelf200
// @Router /admin [get]
func AdminGetSelf(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	// remove password from response
	admin.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"data":    admin,
		"success": true,
		"message": "Returned own details",
	})
}

// AdminEditSelf godoc
// @Summary Edit own details
// @Tags Admin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param  Body body swagger.adminEditSelfBody true "JSON request body<br><br>* All fields are optional.<br>** If field left blank it will not be updated."
// @Success 200 {object} swagger.adminEditAdmin200
// @Failure 409 {object} swagger.adminEditAdmin409
// @Failure 422 {object} swagger.adminEditAdmin422
// @Router /admin [put]
func AdminEditSelf(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))
	data := models.Admin{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	data.ID = admin.ID

	status, err := admin.AdminEditSelf(&data)
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
		"message": "Updated own details",
	})

	fmt.Println(admin)
}

// AdminCreateClient godoc
// @Summary Create a new client
// @Tags Admin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param  Body body swagger.adminCreateClientBody true "JSON request body<br><br>* Phone is optional<br>** Notes is optional"
// @Success 201 {object} swagger.adminCreateClient201
// @Failure 409 {object} swagger.adminCreateClient409
// @Failure 422 {object} swagger.adminCreateClient422
// @Router /clients [post]
func AdminCreateClient(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))
	data := models.Client{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	status, err := admin.CreateClient(&data)
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
		Phone     string `json:"phone"`
		Notes     string `json:"notes"`
		Status    string `json:"status"`
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result{
			ID:        data.ID,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Phone:     data.Phone,
			Notes:     data.Notes,
			Status:    data.Status,
		},
		"success": true,
		"message": "Created new client",
	})
}

// AdminGetAllClients godoc
// @Summary Get all clients
// @Tags Admin
// @Produce json
// @Security bearerAuth
// @Success 200 {object} swagger.adminGetAllClients200
// @Failure 404 {object} swagger.adminGetClientByID404
// @Router /clients [get]
func AdminGetAllClients(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	clients, status, err := admin.GetAllClients()
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_results": len(clients),
		"success":       true,
		"message":       "All clients",
		"data":          clients,
	})
}

// AdminGetClientByID godoc
// @Summary Get client by ID
// @Tags Admin
// @Produce json
// @Security bearerAuth
// @Param id path int true "Admin ID"
// @Success 200 {object} swagger.adminGetClientByID200
// @Failure 404 {object} swagger.adminGetClientByID404
// @Failure 422 {object} swagger.adminGetClientByID422
// @Router /clients/{id} [get]
func AdminGetClientByID(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))

	// id, err := strconv.Atoi(c.Param("id"))
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "ID in invalid format", "success": false})
		return
	}

	data, status, err := admin.GetClientByID(id)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"success": true,
		"message": "Returned client",
	})
}

// AdminEditClient godoc
// @Summary Edit a client
// @Tags Admin
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param id path int true "Client ID"
// @Param  Body body swagger.adminEditClientBody true "JSON request body<br><br>* All fields are optional.<br>** If field left blank it will not be updated."
// @Success 200 {object} swagger.adminEditClient200
// @Failure 404 {object} swagger.adminEditClient404
// @Failure 409 {object} swagger.adminEditClient409
// @Failure 422 {object} swagger.adminEditClient422
// @Router /clients/{id} [put]
func AdminEditClient(c *gin.Context) {
	admin := *(c.Keys["admin"].(*models.Admin))
	client := models.Client{}

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	// Verify ID paramater is an integer and add to data struct
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "ID in invalid format", "success": false})
		return
	}
	client.ID = id

	status, err := admin.EditClient(&client)
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
		Phone     string `json:"phone"`
		Notes     string `json:"notes"`
		Status    string `json:"status"`
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Email:     client.Email,
			Phone:     client.Phone,
			Notes:     client.Notes,
			Status:    client.Status,
		},
		"success": true,
		"message": "Updated client",
	})
}

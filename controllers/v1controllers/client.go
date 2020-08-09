package v1controllers

import (
	"fmt"
	"net/http"

	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"
	"github.com/gin-gonic/gin"
)

// ClientGetSelf godoc
// @Summary Get own details
// @Tags Client
// @Produce json
// @Security bearerAuth
// @Success 200 {object} swagger.clientGetSelf200
// @Router / [get]
func ClientGetSelf(c *gin.Context) {
	client := *(c.Keys["client"].(*models.Client))

	// we do not want to show the client all details
	type result struct {
		ID        uint64 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Email:     client.Email,
			Phone:     client.Phone,
		},
		"success": true,
		"message": "Returned own details",
	})
}

// ClientEditSelf godoc
// @Summary Edit own details
// @Tags Client
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param  Body body swagger.clientEditSelfBody true "JSON request body<br><br>* All fields are optional.<br>** If field left blank it will not be updated."
// @Success 200 {object} swagger.clientEditSelf200
// @Failure 409 {object} swagger.adminEditAdmin409
// @Failure 422 {object} swagger.adminEditAdmin422
// @Router / [put]
func ClientEditSelf(c *gin.Context) {
	client := *(c.Keys["client"].(*models.Client))

	// Verify JSON in correct format
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "message": "Invalid JSON provided"})
		return
	}

	status, err := client.ClientEditSelf()
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error(), "success": false})
		return
	}

	type result struct {
		ID        uint64 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Email:     client.Email,
			Phone:     client.Phone,
		},
		"success": true,
		"message": "Updated own details",
	})

	fmt.Println(client)
}

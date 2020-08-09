package models

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/utils/validation"
)

// EditSelf Godoc
func (client *Client) ClientEditSelf() (int, error) {

	// prevent modifying certain fields
	client.Notes = ""
	client.Status = ""

	// if email is provided verify
	if client.Email != "" {
		if !validation.EmailValid(client.Email) {
			return http.StatusUnprocessableEntity, fmt.Errorf("Email is in an invalid format")
		}
	}

	// if password is provided verify
	if client.Password != "" {
		if !validation.PasswordValid(client.Password) {
			return http.StatusUnprocessableEntity, fmt.Errorf("Password does not meet requirements")
		}
		client.Password = hashPassword(client.Password)
	}

	// Update
	result := config.Config.DB.Model(&client).Update(&client)
	if result.Error != nil {
		fmt.Println(result.Error)
		if strings.Contains(result.Error.Error(), "email") {
			return http.StatusConflict, fmt.Errorf("Email address already in use")
		}
		return http.StatusInternalServerError, result.Error
	}

	// Get updated data for response
	config.Config.DB.First(&client)

	return 0, nil
}

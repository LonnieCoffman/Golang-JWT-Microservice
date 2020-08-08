package models

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/utils/validation"
)

// EditSelf Godoc
func (admin *Admin) AdminEditSelf(data *Admin) (int, error) {

	// prevent modifying status and role
	data.Role = ""
	data.Status = ""

	// if email is provided verify
	if data.Email != "" {
		if !validation.EmailValid(data.Email) {
			return http.StatusUnprocessableEntity, fmt.Errorf("Email is in an invalid format")
		}
	}

	// if password is provided verify
	if data.Password != "" {
		if !validation.PasswordValid(data.Password) {
			return http.StatusUnprocessableEntity, fmt.Errorf("Password does not meet requirements")
		}
		data.Password = hashPassword(data.Password)
	}

	// Update
	result := config.Config.DB.Model(&data).Update(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
		if strings.Contains(result.Error.Error(), "email") {
			return http.StatusConflict, fmt.Errorf("Email address already in use")
		}
		return http.StatusInternalServerError, result.Error
	}

	// Get updated data for response
	config.Config.DB.First(&data)

	return 0, nil
}

// CreateClient Godoc
func (admin *Admin) CreateClient(data *Client) (int, error) {

	// we do not want these fields available here:
	data.ID = 0
	data.Status = "active"

	if data.FirstName == "" {
		return http.StatusUnprocessableEntity, fmt.Errorf("First name is required")
	}

	if data.LastName == "" {
		return http.StatusUnprocessableEntity, fmt.Errorf("Last name is required")
	}

	if data.Email == "" || !validation.EmailValid(data.Email) {
		return http.StatusUnprocessableEntity, fmt.Errorf("Email is in an invalid format")
	}

	if data.Password == "" || !validation.PasswordValid(data.Password) {
		return http.StatusUnprocessableEntity, fmt.Errorf("Password does not meet requirements")
	}

	// hash password
	data.Password = hashPassword(data.Password)

	if err := config.Config.DB.Create(&data).Error; err != nil {
		if strings.Contains(err.Error(), "email") {
			return http.StatusConflict, fmt.Errorf("Client with email already exists")
		}
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

// GetAllClients Godoc
func (admin *Admin) GetAllClients() ([]Client, int, error) {
	clients := []Client{}

	err := config.Config.DB.Select("id, first_name, last_name, email, phone, notes, status, created_at, updated_at").Find(&clients).Error
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("Clients not found")
	}

	return clients, 0, nil
}

// GetClientByID Godoc
func (admin *Admin) GetClientByID(id uint64) (Client, int, error) {
	client := Client{}

	err := config.Config.DB.Select("id, first_name, last_name, email, phone, notes, status, created_at, updated_at").Where("id = ?", id).Find(&client).Error
	if err != nil {
		return client, http.StatusNotFound, fmt.Errorf("Client not found")
	}

	return client, 0, nil
}

// EditClient Godoc
func (admin *Admin) EditClient(client *Client) (int, error) {

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

	// if status provided it can only be active or suspended
	if client.Status != "" {
		if client.Status != "active" && client.Status != "suspended" {
			return http.StatusUnprocessableEntity, fmt.Errorf("Status must contain a valid option")
		}
	}

	// Update
	result := config.Config.DB.Model(&client).Update(&client)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "email") {
			return http.StatusConflict, fmt.Errorf("Email address already in use")
		}
		return http.StatusInternalServerError, result.Error
	}
	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("Active admin not found")
	}

	// Get updated client for response
	config.Config.DB.First(&client)

	return 0, nil
}

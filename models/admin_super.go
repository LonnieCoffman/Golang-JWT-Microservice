package models

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/utils/validation"

	"golang.org/x/crypto/bcrypt"
)

// GetAllAdmins Godoc
func (admin *Admin) GetAllAdmins() ([]Admin, int, error) {
	admins := []Admin{}

	err := config.Config.DB.Select("id, first_name, last_name, email, role, status, created_at, updated_at").Find(&admins).Error
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("Admins not found")
	}

	return admins, 0, nil
}

// GetAdminByID Godoc
func (admin *Admin) GetAdminByID(id uint64) (Admin, int, error) {
	data := Admin{}

	err := config.Config.DB.Select("id, first_name, last_name, email, role, status, created_at, updated_at").Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, http.StatusNotFound, fmt.Errorf("Admin not found")
	}

	return data, 0, nil
}

// CreateAdmin Godoc
func (admin *Admin) CreateAdmin(data *Admin) (int, error) {

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

	if data.Role == "" || (data.Role != "admin" && data.Role != "superadmin") {
		return http.StatusUnprocessableEntity, fmt.Errorf("Role must contain a valid option")
	}

	// default to active
	data.Status = "active"

	// hash password
	data.Password = hashPassword(data.Password)

	if err := config.Config.DB.Create(&data).Error; err != nil {
		if strings.Contains(err.Error(), "email") {
			return http.StatusConflict, fmt.Errorf("Admin with email already exists")
		}
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

// DeleteAdmin Godoc
func (admin *Admin) DeleteAdmin(id uint64) (Admin, int, error) {
	data := Admin{}
	data.ID = id

	rows := config.Config.DB.Delete(&data).RowsAffected
	if rows == 0 {
		return data, http.StatusNotFound, fmt.Errorf("Active admin not found")
	}

	config.Config.DB.Where("admin_id = ?", id).Delete(AdminSession{})

	return data, 0, nil
}

// EditAdmin Godoc
func (admin *Admin) EditAdmin(data *Admin) (int, error) {

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

	// if role provided it must be either admin or superadmin
	if data.Role != "" {
		if data.Role != "admin" && data.Role != "superadmin" {
			return http.StatusUnprocessableEntity, fmt.Errorf("Role must contain a valid option")
		}
	}

	// if status provided it can only be active or suspended
	if data.Status != "" {
		if data.Status != "active" && data.Status != "suspended" {
			return http.StatusUnprocessableEntity, fmt.Errorf("Status must contain a valid option")
		}
	}

	// Update
	result := config.Config.DB.Model(&data).Update(&data)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "email") {
			return http.StatusConflict, fmt.Errorf("Email address already in use")
		}
		return http.StatusInternalServerError, result.Error
	}
	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("Active admin not found")
	}

	// Get updated data for response
	config.Config.DB.First(&data)

	return 0, nil
}

func hashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pass)
}

// DeleteClient Godoc
func (admin *Admin) DeleteClient(id uint64) (Client, int, error) {
	data := Client{}
	data.ID = id

	rows := config.Config.DB.Delete(&data).RowsAffected
	if rows == 0 {
		return data, http.StatusNotFound, fmt.Errorf("Active client not found")
	}

	config.Config.DB.Where("client_id = ?", id).Delete(ClientSession{})

	return data, 0, nil
}

package utils

import (
	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"

	"golang.org/x/crypto/bcrypt"
)

// Migrate creates all of the database tables
func Migrate() {
	config.Config.DB.AutoMigrate(&models.Admin{}, &models.AdminSession{}, &models.Client{}, &models.ClientSession{})
	config.Config.DB.Model(&models.AdminSession{}).AddForeignKey("admin_id", "admins(id)", "CASCADE", "CASCADE")
	config.Config.DB.Model(&models.ClientSession{}).AddForeignKey("client_id", "clients(id)", "CASCADE", "CASCADE")
}

// Seed populates the database with sample data
func Seed() {
	admin := &models.Admin{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "admin@test.com",
		Password:  hashPassword("password"),
		Role:      "admin",  // 1 = admin, 2 = superadmin
		Status:    "active", // 1 = active, 2 = suspended, 3 = deleted
	}
	superadmin := &models.Admin{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "superadmin@test.com",
		Password:  hashPassword("password"),
		Role:      "superadmin", // 1 = admin, 2 = superadmin
		Status:    "active",     // 1 = active, 2 = suspended, 3 = deleted
	}
	config.Config.DB.Save(admin)
	config.Config.DB.Save(superadmin)
}

func hashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pass)
}

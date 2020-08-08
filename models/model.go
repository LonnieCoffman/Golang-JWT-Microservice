package models

import (
	"time"
)

// // Model defined same os gorm.Model but column names and json tags added
// type Model struct {
// 	ID        uint64     `gorm:"primary_key;column:id" json:"id"`
// 	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
// 	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
// 	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
// }

// Admin model
type Admin struct {
	ID        uint64     `gorm:"primary_key;column:id" json:"id"`
	FirstName string     `gorm:"column:first_name" json:"first_name"`
	LastName  string     `gorm:"column:last_name" json:"last_name"`
	Email     string     `gorm:"column:email;UNIQUE" json:"email"`
	Password  string     `gorm:"column:password" json:"password,omitempty"`
	Role      string     `gorm:"column:role;type:enum('admin','superadmin')" json:"role"`
	Status    string     `gorm:"column:status;type:enum('active','suspended')" json:"status"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

// AdminSession model
type AdminSession struct {
	ID                     uint64    `gorm:"primary_key;column:id" json:"id"`
	AdminID                uint64    `gorm:"column:admin_id" json:"admin_id"`
	AccessToken            string    `gorm:"column:access_token;UNIQUE" json:"access_token"`
	AccessTokenExpiration  int64     `gorm:"column:access_token_expiration;type:int(12)" json:"access_token_expiration"`
	RefreshToken           string    `gorm:"column:refresh_token;UNIQUE" json:"refresh_token"`
	RefreshTokenExpiration int64     `gorm:"column:refresh_token_expiration;type:int(12)" json:"refresh_token_expiration"`
	RemoteAddress          string    `gorm:"column:remote_address;type:varchar(64)" json:"remote_address"`
	Revoked                bool      `gorm:"column:revoked;type:boolean" json:"revoked"`
	CreatedAt              time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Client model
type Client struct {
	ID        uint64     `gorm:"primary_key;column:id" json:"id"`
	FirstName string     `gorm:"column:first_name" json:"first_name"`
	LastName  string     `gorm:"column:last_name" json:"last_name"`
	Email     string     `gorm:"column:email;UNIQUE" json:"email"`
	Notes     string     `gorm:"column:notes;type:text" json:"notes"`
	Phone     string     `gorm:"column:phone" json:"phone"`
	Password  string     `gorm:"column:password" json:"password,omitempty"`
	Status    string     `gorm:"column:status;type:enum('active','suspended')" json:"status"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

// ClientSession model
type ClientSession struct {
	ID                     uint64    `gorm:"primary_key;column:id" json:"id"`
	ClientID               uint64    `gorm:"column:client_id" json:"client_id"`
	AccessToken            string    `gorm:"column:access_token;UNIQUE" json:"access_token"`
	AccessTokenExpiration  int64     `gorm:"column:access_token_expiration;type:int(12)" json:"access_token_expiration"`
	RefreshToken           string    `gorm:"column:refresh_token;UNIQUE" json:"refresh_token"`
	RefreshTokenExpiration int64     `gorm:"column:refresh_token_expiration;type:int(12)" json:"refresh_token_expiration"`
	RemoteAddress          string    `gorm:"column:remote_address;type:varchar(64)" json:"remote_address"`
	Revoked                bool      `gorm:"column:revoked;type:boolean" json:"revoked"`
	CreatedAt              time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at" json:"updated_at"`
}

package models

import (
	"fmt"
	"net/http"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate verifies a username and password combination exists in the database
func (admin *Admin) Authenticate() (int, error) {

	// email and password cannot be missing or blank
	if admin.Email == "" || admin.Password == "" {
		return http.StatusUnprocessableEntity, fmt.Errorf("Required fields either missing or empty")
	}

	// set provided password variable
	password := admin.Password

	// get admin from the database
	err := config.Config.DB.Model(Admin{}).Where("email = ?", admin.Email).Take(&admin).Error
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return http.StatusUnauthorized, fmt.Errorf("Please provide valid credentials")
	}

	// Has account been "soft" deleted
	if admin.Status == "deleted" {
		return http.StatusUnauthorized, fmt.Errorf("Please provide valid credentials")
	}

	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Please provide valid credentials")
	}

	// Has account been suspended
	if admin.Status == "suspended" {
		return http.StatusForbidden, fmt.Errorf("Account has been suspended")
	}

	return 0, nil
}

// CreateTokens creates and stores a new admin auth token and refresh token
func (admin *Admin) CreateTokens(adminSession *AdminSession) error {
	// create access token
	expireUnix := time.Now().Add(time.Minute * time.Duration(config.Config.ADMIN_TOKEN_EXPIRE)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	atClaims := token.Claims.(jwt.MapClaims)
	atClaims["sub"] = admin.ID
	atClaims["exp"] = expireUnix
	atClaims["role"] = admin.Role
	signedAtToken, err := token.SignedString([]byte(config.Config.ADMIN_TOKEN_SECRET))
	if err != nil {
		return err
	}
	adminSession.AccessToken = signedAtToken
	adminSession.AccessTokenExpiration = expireUnix

	// create refresh token
	expireUnix = time.Now().Add(time.Minute * time.Duration(config.Config.ADMIN_REFRESH_TOKEN_EXPIRE)).Unix()
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = admin.ID
	rtClaims["exp"] = expireUnix
	signedRtToken, err := refreshToken.SignedString([]byte(config.Config.ADMIN_REFRESH_SECRET))
	if err != nil {
		return err
	}
	adminSession.RefreshToken = signedRtToken
	adminSession.RefreshTokenExpiration = expireUnix

	// store tokens in database on new or update on refresh
	if adminSession.ID == 0 {
		config.Config.DB.Create(&adminSession)
	} else {
		config.Config.DB.Save(&adminSession)
	}

	return nil
}

// DestroyTokens deletes the tokens from the database.  Not returning an error here, since it would be meaningless to a user logging out
// as they would have already deleted thier stored token and would not be using it in the future.
func (admin *Admin) DestroyTokens(adminSession *AdminSession) {
	config.Config.DB.Delete(&adminSession)
	return
}

func (admin *Admin) VerifyRefresh(adminSession *AdminSession) (int, error) {
	// email and password cannot be missing or blank
	if adminSession.RefreshToken == "" {
		return http.StatusUnprocessableEntity, fmt.Errorf("Required fields either missing or empty")
	}

	// validate token
	_, err := jwt.Parse(adminSession.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.ADMIN_REFRESH_SECRET), nil
	})
	if err != nil {
		return http.StatusUnauthorized, err
	}

	err = config.Config.DB.Table("admin_sessions").Select("*").
		Joins("JOIN admins on admin_sessions.admin_id = admins.id").
		Where("refresh_token = ?", adminSession.RefreshToken).
		Scan(&admin).Scan(&adminSession).
		Error
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Invalid Token")
	}
	admin.ID = adminSession.AdminID // reinsert admin ID

	if admin.Status == "suspended" {
		return http.StatusForbidden, fmt.Errorf("Account has been suspended")
	}

	if admin.Status == "deleted" {
		return http.StatusUnauthorized, fmt.Errorf("Invalid token")
	}

	if adminSession.Revoked {
		return http.StatusUnauthorized, fmt.Errorf("Invalid Token")
	}

	return 0, nil
}

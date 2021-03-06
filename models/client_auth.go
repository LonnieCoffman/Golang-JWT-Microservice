package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate verifies a username and password combination exists in the database
func (client *Client) Authenticate() (int, error) {

	// email and password cannot be missing or blank
	if client.Email == "" || client.Password == "" {
		return http.StatusUnprocessableEntity, fmt.Errorf("Required fields either missing or empty")
	}

	// set provided password variable
	password := client.Password

	// get client from the database
	err := config.Config.DB.Model(Client{}).Where("email = ?", client.Email).Take(&client).Error
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return http.StatusUnauthorized, fmt.Errorf("Please provide valid credentials")
	}

	// Has account been "soft" deleted
	if client.Status == "deleted" {
		return http.StatusUnauthorized, fmt.Errorf("Please provide valid credentials")
	}

	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password)); err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Please provide valid credentials")
	}

	// Has account been suspended
	if client.Status == "suspended" {
		return http.StatusForbidden, fmt.Errorf("Account has been suspended")
	}

	return 0, nil
}

// CreateTokens creates and stores a new client auth token and refresh token
func (client *Client) CreateTokens(clientSession *ClientSession) error {
	// create access token
	expireUnix := time.Now().Add(time.Minute * time.Duration(config.Config.TOKEN_EXPIRE)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	atClaims := token.Claims.(jwt.MapClaims)
	atClaims["sub"] = client.ID
	atClaims["exp"] = expireUnix
	signedAtToken, err := token.SignedString([]byte(config.Config.TOKEN_SECRET))
	if err != nil {
		return err
	}
	clientSession.AccessToken = signedAtToken
	clientSession.AccessTokenExpiration = expireUnix

	// create refresh token
	expireUnix = time.Now().Add(time.Minute * time.Duration(config.Config.REFRESH_TOKEN_EXPIRE)).Unix()
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = client.ID
	rtClaims["exp"] = expireUnix
	signedRtToken, err := refreshToken.SignedString([]byte(config.Config.REFRESH_SECRET))
	if err != nil {
		return err
	}
	clientSession.RefreshToken = signedRtToken
	clientSession.RefreshTokenExpiration = expireUnix

	// store tokens in database on new or update on refresh
	if clientSession.ID == 0 {
		config.Config.DB.Create(&clientSession)
	} else {
		config.Config.DB.Save(&clientSession)
	}

	return nil
}

// DestroyClientTokens deletes the tokens from the database.  Not returning an error here, since it would be meaningless to a user logging out
// as they would have already deleted thier stored token and would not be using it in the future.
func (client *Client) DestroyClientTokens(clientSession *ClientSession) {
	config.Config.DB.Delete(&clientSession)
	return
}

// VerifyRefresh godoc
func (client *Client) VerifyRefresh(clientSession *ClientSession) (int, error) {
	// email and password cannot be missing or blank
	if clientSession.RefreshToken == "" {
		return http.StatusUnprocessableEntity, fmt.Errorf("Required fields either missing or empty")
	}

	// validate token
	_, err := jwt.Parse(clientSession.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.REFRESH_SECRET), nil
	})
	if err != nil {
		return http.StatusUnauthorized, err
	}

	err = config.Config.DB.Table("client_sessions").Select("*").
		Joins("JOIN clients on client_sessions.client_id = clients.id").
		Where("refresh_token = ?", clientSession.RefreshToken).
		Scan(&client).Scan(&clientSession).
		Error
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Invalid Token")
	}
	client.ID = clientSession.ClientID // reinsert client ID

	if client.Status == "suspended" {
		return http.StatusForbidden, fmt.Errorf("Account has been suspended")
	}

	if client.Status == "deleted" {
		return http.StatusUnauthorized, fmt.Errorf("Invalid token")
	}

	if clientSession.Revoked {
		return http.StatusUnauthorized, fmt.Errorf("Invalid Token")
	}

	return 0, nil
}

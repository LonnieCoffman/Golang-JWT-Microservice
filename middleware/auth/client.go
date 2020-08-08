package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	client        = models.Client{}
	clientSession = models.ClientSession{}
)

// Client godoc
func Client() gin.HandlerFunc {
	return func(c *gin.Context) {
		if status, err := validateClient(c); err != nil {
			c.JSON(status, gin.H{"success": false, "message": err.Error()})
			c.Abort()
		}
		c.Next()
	}
}

func validateClient(c *gin.Context) (int, error) {

	// Verify token
	jwtToken, err := VerifyClientToken(c.Request)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Please provide a valid token")
	}

	// Extract the client ID from the token
	if err := getClientID(jwtToken); err != nil {
		return http.StatusUnauthorized, err
	}

	err = config.Config.DB.Table("client_sessions").Select("*").
		Joins("JOIN clients on client_sessions.client_id = clients.id").
		Where("client_id = ? AND access_token = ?", client.ID, clientSession.AccessToken).
		Scan(&client).Scan(&clientSession).
		Error
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}
	client.ID = clientSession.ClientID // reinsert client ID

	if client.Status == "suspended" {
		return http.StatusForbidden, fmt.Errorf("Account has been suspended")
	}

	if client.Status == "deleted" {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	if clientSession.Revoked {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	c.Set("client", &client)
	c.Set("clientSession", &clientSession)

	return 0, nil
}

func getClientID(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		clientID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			return err
		}
		client.ID = clientID
	}
	return nil
}

// VerifyClientToken validates the token signing method and secret
func VerifyClientToken(r *http.Request) (*jwt.Token, error) {
	// extract the token from the header
	auth := r.Header.Get("Authorization")
	bearer := strings.Split(auth, "Bearer ")
	if len(bearer) != 2 {
		return nil, fmt.Errorf("token missing")
	}

	clientSession.AccessToken = bearer[1]

	// verify the token
	token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

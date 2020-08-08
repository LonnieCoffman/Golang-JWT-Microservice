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

type Auth struct {
	ID                uint64
	AccessTokenString string
	AccessToken       *jwt.Token
}

var (
	rolelvl = map[string]int{
		"superadmin": 2,
		"admin":      1,
	}
	authDetails  = Auth{}
	admin        = models.Admin{}
	adminSession = models.AdminSession{}
)

func SuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if status, err := validateAdminRole(c, "superadmin"); err != nil {
			c.JSON(status, gin.H{"success": false, "message": err.Error()})
			c.Abort()
		}
		c.Next()
	}
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if status, err := validateAdminRole(c, "admin"); err != nil {
			c.JSON(status, gin.H{"success": false, "message": err.Error()})
			c.Abort()
		}
		c.Next()
	}
}

func validateAdminRole(c *gin.Context, role string) (int, error) {

	// Verify token
	jwtToken, err := VerifyToken(c.Request)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Please provide a valid token")
	}

	// Extract the admin ID from the token
	if err := getAdminID(jwtToken); err != nil {
		return http.StatusUnauthorized, err
	}

	err = config.Config.DB.Table("admin_sessions").Select("*").
		Joins("JOIN admins on admin_sessions.admin_id = admins.id").
		Where("admin_id = ? AND access_token = ?", admin.ID, adminSession.AccessToken).
		Scan(&admin).Scan(&adminSession).
		Error
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}
	admin.ID = adminSession.AdminID // reinsert admin ID

	if rolelvl[admin.Role] < rolelvl[role] {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	if admin.Status == "suspended" {
		return http.StatusForbidden, fmt.Errorf("Account has been suspended")
	}

	if admin.Status == "deleted" {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	if adminSession.Revoked {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	c.Set("admin", &admin)
	c.Set("adminSession", &adminSession)

	return 0, nil
}

func getAdminID(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		adminID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			return err
		}
		admin.ID = adminID
	}
	return nil
}

// VerifyToken validates the token signing method and secret
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	// extract the token from the header
	auth := r.Header.Get("Authorization")
	bearer := strings.Split(auth, "Bearer ")
	if len(bearer) != 2 {
		return nil, fmt.Errorf("token missing")
	}

	adminSession.AccessToken = bearer[1]

	// verify the token
	token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.ADMIN_TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

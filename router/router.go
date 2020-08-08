package router

import (
	"github.com/LonnieCoffman/Golang-JWT-Microservice/controllers/v1controllers"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/middleware/auth"

	"github.com/gin-gonic/gin"
)

// SetupRoutes is the versioning entry point for routing
func SetupRoutes(api *gin.Engine) {
	v1 := api.Group("/v1")
	groupV1Routes(v1)
}

// groupV1Routes contains all of the grouped V1 routes
func groupV1Routes(v1 *gin.RouterGroup) {

	// status
	v1.GET("/ping", v1controllers.Ping)

	// groups
	admin := v1.Group("/admin")
	admins := v1.Group("/admins")
	clients := v1.Group("/clients")

	// admin - authentication routes
	admin.POST("/login", v1controllers.AdminLogin)
	admin.POST("/refresh", v1controllers.AdminRefresh)
	admin.POST("/logout", auth.Admin(), v1controllers.AdminLogout)

	// admin - super admin routes
	admins.GET("/", auth.SuperAdmin(), v1controllers.AdminGetAll)
	admins.GET("/:id", auth.SuperAdmin(), v1controllers.AdminGetByID)
	admins.POST("/", auth.SuperAdmin(), v1controllers.AdminCreate)
	admins.PUT("/:id", auth.SuperAdmin(), v1controllers.AdminEditAdmin)
	admins.DELETE("/:id", auth.SuperAdmin(), v1controllers.AdminDeleteAdmin)
	clients.DELETE("/:id", auth.SuperAdmin(), v1controllers.AdminDeleteClient)

	// admin - standard admin routes
	admin.GET("/", auth.Admin(), v1controllers.AdminGetSelf)
	admin.PUT("/", auth.Admin(), v1controllers.AdminEditSelf)
	clients.POST("/", auth.Admin(), v1controllers.AdminCreateClient)
	clients.GET("/", auth.Admin(), v1controllers.AdminGetAllClients)
	clients.GET("/:id", auth.Admin(), v1controllers.AdminGetClientByID)
	clients.PUT("/:id", auth.Admin(), v1controllers.AdminEditClient)

	// client - authentication routes
	v1.POST("/login", v1controllers.ClientLogin)
	//v1.POST("/refresh", v1controllers.ClientRefresh)
	v1.POST("/logout", auth.Client(), v1controllers.ClientLogout)
}

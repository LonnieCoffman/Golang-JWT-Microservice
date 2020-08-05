package router

import (
	v1controllers "test/controllers/v1"

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
}

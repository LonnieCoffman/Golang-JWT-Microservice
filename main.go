package main

import (
	"fmt"
	"log"

	"github.com/LonnieCoffman/Golang-JWT-Microservice/config"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/docs"
	"github.com/LonnieCoffman/Golang-JWT-Microservice/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Load configuration from file if it exists, ignoring any error since config.yml is only used locally
	config.LoadConfig(".")

	// intialize Gin with a logger and recovery
	api := gin.New()
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	// setup api routes
	router.SetupRoutes(api)

	// Swagger 2.0 Meta Information
	if config.Config.SWAGGER_ENABLE {
		docs.SwaggerInfo.Title = "Authentication Service"
		docs.SwaggerInfo.Description = "API providing authentication services for admins and clients"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.BasePath = "/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// connect to MySQL database
	config.Config.DB, config.Config.DBErr = gorm.Open("mysql", config.Config.DB_DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// Migrate data into database
	//utils.Migrate()
	//utils.Seed()

	log.Println("Successfully connected to database")
	defer config.Config.DB.Close()

	// start server
	err := api.Run(fmt.Sprintf("%v:%v", config.Config.SERVER_HOST, config.Config.SERVER_PORT))
	if err != nil {
		fmt.Println(err)
	}
}

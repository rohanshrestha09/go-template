package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/go-template/configs"
	docs "github.com/rohanshrestha09/go-template/docs"
	"github.com/rohanshrestha09/go-template/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {

	configs.LoadEnv()

	configs.InitializeFirebaseApp()

	configs.InitializeDatabase()
}

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization

func main() {
	server := gin.Default()

	server.SetTrustedProxies([]string{"localhost"})

	server.MaxMultipartMemory = 10 << 20

	router := server.Group("/api/v1")

	routes.SetupRouter(router)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":5000")

}

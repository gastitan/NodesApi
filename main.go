package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gotest.com/nodes-api/controller"
	"gotest.com/nodes-api/database"

	"gotest.com/nodes-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getNodes(ctx *gin.Context) {
	controller.FindAll(ctx)
}

func saveNode(ctx *gin.Context) {
	controller.Save(ctx)
}

func getNodeById(ctx *gin.Context) {
	controller.FindById(ctx)
}

func updateNode(ctx *gin.Context) {
	controller.Update(ctx)
}

func nearestNode(ctx *gin.Context) {
	controller.Nearest(ctx)
}

func main() {

	docs.SwaggerInfo.Title = "Nodes API"
	docs.SwaggerInfo.Description = "API of nodes"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/nodes", getNodes)
	router.GET("/nodes/:id", getNodeById)
	router.POST("/nodes", saveNode)
	router.PUT("/nodes/:id", updateNode)

	//check params lat, lng
	router.GET("/nodes/nearest", nearestNode)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + os.Getenv("PORT"))
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = database.OpenDatabase()
	if err != nil {
		log.Fatalf("Error connection to DB")
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k-yoshigai/learning-jwt-go/controllers"
	"github.com/k-yoshigai/learning-jwt-go/database"
	"github.com/k-yoshigai/learning-jwt-go/middlewares"
)

func main() {
	connectionsString := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	database.Connect(connectionsString)
	database.Migrate()
	router := initRouter()
	router.Run(":8088")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}

package main

import (
	"os"
	
	"schoolbackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	route := gin.New()
	route.Use(gin.Logger())
	routes.CommonRoutes(route)
	routes.ParentRoute(route)
	routes.StudentRoute(route)
	routes.TeacherRoute(route)
	route.Run(":" + port)

}
package routes

import (
	"schoolbackend/middleware"

	"github.com/gin-gonic/gin"
)

func StudentRoute(incomingRoute *gin.Engine) {
	incomingRoute.Use(middleware.AuthenticateStudent())
	studentRoute := incomingRoute.Group("/student")
	{
		studentRoute.GET("/student/mygrade",nil)
		studentRoute.GET("/student/myattendance",nil)
		studentRoute.POST("/student/assignment/:id",nil)
	}
	incomingRoute.GET("/club",nil)
	incomingRoute.GET("/club/:id",nil)
	incomingRoute.POST("/club/:id/apply",nil)
}
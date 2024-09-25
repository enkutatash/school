package routes

import (
	"schoolbackend/middleware"

	"github.com/gin-gonic/gin"
)

func ParentRoute(incomingRoute *gin.Engine) {
	parentRoute := incomingRoute.Group("/parent")
	parentRoute.Use(middleware.AuthenticateParent())
	{
		parentRoute.GET("/child_attendance", nil)
		parentRoute.GET("/child_discipline_board", nil)
		parentRoute.GET("/child_mark", nil)
		parentRoute.GET("/child_home_teacher", nil)
	}
}
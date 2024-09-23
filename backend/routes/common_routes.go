package routes

import (
	"schoolbackend/controller"
	"schoolbackend/usecase"

	"github.com/gin-gonic/gin"
)

var (
	appusecase usecase.CommonUsecases      = usecase.NewCommonUsecase()
	app        controller.CommonController = controller.NewCommonController(appusecase)
)

func CommonRoutes(incomingRoute *gin.Engine) {

	incomingRoute.GET("/club", nil)
	incomingRoute.GET("/club/:id", nil)
	incomingRoute.POST("/club/:id/apply", nil)

	incomingRoute.POST("/register/student", app.RegisterStudent)
	incomingRoute.POST("/register/teacher", app.RegisterTeacher)
	incomingRoute.POST("/login/student", app.LoginStudent)
	incomingRoute.POST("/login/teacher", app.LoginTeacher)
}

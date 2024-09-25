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
	// Club-related routes
incomingRoute.GET("/clubs", app.GetClubs)  // List of clubs
incomingRoute.GET("/club/:club_id", app.GetClubByID)  // Get specific club by ID
incomingRoute.POST("/club/:club_id/apply", app.ApplyForClub)  // Apply for a club

// Distinguishing application management with more specific routes
incomingRoute.PUT("/club/:club_id/application/:student_id/accept", app.AcceptClubRequest)  // Accept a student's application
incomingRoute.PUT("/club/:club_id/application/:student_id/reject", app.RejectClubRequest)  // Reject a student's application
incomingRoute.GET("/club/:club_id/applications", app.GetClubApplications)  // Get all applications for a club

// Registration and login routes
incomingRoute.POST("/register/student", app.RegisterStudent)
incomingRoute.POST("/register/teacher", app.RegisterTeacher)
incomingRoute.POST("/login/student", app.LoginStudent)
incomingRoute.POST("/login/teacher", app.LoginTeacher)

	
}

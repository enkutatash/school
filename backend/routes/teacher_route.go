package routes

import (
	"schoolbackend/middleware"

	"github.com/gin-gonic/gin"
)

func TeacherRoute(incomingRoute *gin.Engine) {
	incomingRoute.Use(middleware.AuthenticateTeacher())
	teacherRoute := incomingRoute.Group("/teacher")
	{
		teacherRoute.GET("/:grade/:section_id/grade", nil)
		teacherRoute.PUT("/:grade/:section_id/grade", nil)
		teacherRoute.GET("/:grade/:section_id/attendance", nil)
		teacherRoute.PUT("/:grade/:section_id/attendance", nil)
		teacherRoute.POST("/quiz", nil)
		teacherRoute.PUT("/quiz/:quiz_id/:question_id", nil)
		teacherRoute.POST("/quiz/:quiz_id/:student_submission_id", nil)
		teacherRoute.POST("/assignment", nil)
		teacherRoute.GET("/assignment_submissions", nil)
		// create a section
		teacherRoute.POST("/section", nil)
		// add student to section
		teacherRoute.POST("/section/:section_id/student", nil)
		// remove student from section
		teacherRoute.DELETE("/section/:section_id/student/:student_id", nil)
	}
	incomingRoute.GET("/club", nil)
	incomingRoute.GET("/club/:id", nil)
	incomingRoute.POST("/club/:id/apply", nil)
}
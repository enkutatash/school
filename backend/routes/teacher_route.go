package routes

import (
	"schoolbackend/controller"
	"schoolbackend/middleware"
	"schoolbackend/usecase"

	"github.com/gin-gonic/gin"
)

var (
	teacherusecase usecase.TeacherUseCase       = usecase.NewTeacherUsecase()
	teacherService controller.TeacherController = controller.NewTeacherController(teacherusecase)
)

func TeacherRoute(incomingRoute *gin.Engine) {
	teacherRoute := incomingRoute.Group("/teacher")
	teacherRoute.Use(middleware.AuthenticateTeacher())
	{
		teacherRoute.POST("/found_club", teacherService.FoundClub)
		teacherRoute.PUT("/:club_id/assign_lead", teacherService.AssignLead)

		teacherRoute.GET("/grade/:grade/section/:section_id/grade", nil)
		teacherRoute.PUT("/grade/:grade/section/:section_id/grade", nil)
		teacherRoute.GET("/grade/:grade/section/:section_id/attendance", nil)
		teacherRoute.PUT("/grade/:grade/section/:section_id/attendance", nil)
		
		teacherRoute.POST("/quiz", nil)
		teacherRoute.PUT("/quiz/:quiz_id/:question_id", nil)
		teacherRoute.POST("/quiz/:quiz_id/:student_submission_id", nil)
		teacherRoute.POST("/assignment", nil)
		teacherRoute.GET("/assignment_submissions", nil)
		// create a section
		teacherRoute.POST("/section", teacherService.CreateSection)
		// add courses 
		teacherRoute.PUT("/section/:section_id/courses", teacherService.AddCourses)
		// assign rep to section
		teacherRoute.PUT("/section/:section_id/student/rep", teacherService.AssignSectionRep)
		// add student to section
		teacherRoute.PUT("/section/:section_id/student/add", teacherService.AddStudentToSection)
		// remove student from section
		teacherRoute.PUT("/section/:section_id/student/remove", teacherService.RemoveStudentFromSection)
		// assign teacher to section
		teacherRoute.PUT("/section/:section_id/assign_teacher", teacherService.AssignTeacher)
		// remove teacher from section
		teacherRoute.PUT("/section/:section_id/unassign_teacher", teacherService.UnAssignTeacher)
	}
	// incomingRoute.GET("/club", nil)
	// incomingRoute.GET("/club/:id", nil)
	// incomingRoute.POST("/club/:id/apply", nil)
}

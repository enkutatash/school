package controller

import (
	"net/http"
	"schoolbackend/models"
	"schoolbackend/token"
	"schoolbackend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	TeacherService usecase.TeacherUseCase
)

type TeacherController interface {
	FoundClub(c *gin.Context)
	AssignLead(c *gin.Context)
	CreateSection(c *gin.Context)
	AssignSectionRep(c *gin.Context)
	AddStudentToSection(c *gin.Context)
	RemoveStudentFromSection(c *gin.Context)

	AddCourses(c *gin.Context)
	AssignTeacher(c *gin.Context)
	UnAssignTeacher(c *gin.Context)
}
type TeacherControllerSample struct{}

// AssignTeacher implements TeacherController.
func (t *TeacherControllerSample) AssignTeacher(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sectionID := c.Param("section_id")
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	if sectionID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	section, er := TeacherService.GetSectionByID(sectionID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	if section.HomeTeacherID != claim.(*token.SignedDetails).Uid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only home teacher can assign teacher to this section"})
		return
	}
	var RequestBody struct{
		Subject string `json:"subject"`
		TeacherId string `json:"teacher_id"`
	}

	if er := c.BindJSON(&RequestBody); er != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	} 

	er = TeacherService.AssignTeacher(RequestBody.Subject,RequestBody.TeacherId,sectionID)
	if er != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully assign the teacher"})
}

// RemoveTeacher implements TeacherController.
func (t *TeacherControllerSample) UnAssignTeacher(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sectionID := c.Param("section_id")
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	if sectionID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	section, er := TeacherService.GetSectionByID(sectionID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	if section.HomeTeacherID != claim.(*token.SignedDetails).Uid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only home teacher can assign teacher to this section"})
		return
	}
	var RequestBody struct{
		TeacherId string `json:"teacher_id"`
	}

	if er := c.BindJSON(&RequestBody); er != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	} 

	er = TeacherService.UnAssignTeacher(RequestBody.TeacherId,sectionID)
	if er != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return 
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully unassign the teacher"})
}

// AddCourses implements TeacherController.
func (t *TeacherControllerSample) AddCourses(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sectionID := c.Param("section_id")
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	if sectionID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	section, er := TeacherService.GetSectionByID(sectionID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	if section.HomeTeacherID != claim.(*token.SignedDetails).Uid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only home teacher can add courses to section"})
		return
	}

	var requestBody struct {
		CourseID []string `json:"courses"`
	}

	// Parse the incoming JSON
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	er = TeacherService.AddCoursesToSection(requestBody.CourseID,sectionID)
	if er != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully added courses to section"})





}

// AddStudentToSection implements TeacherController.
func (t *TeacherControllerSample) AddStudentToSection(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sectionID := c.Param("section_id")
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	if sectionID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	section, er := TeacherService.GetSectionByID(sectionID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	if section.HomeTeacherID != claim.(*token.SignedDetails).Uid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only home teacher can add student to section"})
		return
	}

	var requestBody struct {
		StudentID []string `json:"students_id"`
	}

	// Parse the incoming JSON
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}
	invalidStudent, er := TeacherService.AddStudentToSection(sectionID, requestBody.StudentID)

	if len(invalidStudent) > 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Some students are already assigned to a section",
			"assigned students": invalidStudent,
		})
		return
	}
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully added students to section"})

}

// RemoveStudentFromSection implements TeacherController.
func (t *TeacherControllerSample) RemoveStudentFromSection(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sectionID := c.Param("section_id")
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	if sectionID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	section, er := TeacherService.GetSectionByID(sectionID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	if section.HomeTeacherID != claim.(*token.SignedDetails).Uid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only home teacher can remove student from section"})
		return
	}

	var requestBody struct {
		StudentID string `json:"student_id"`
	}

	// Parse the incoming JSON
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}
	er = TeacherService.RemoveStudentFromSection(sectionID, requestBody.StudentID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully removed student from section"})
}

// AssignSectionRep implements TeacherController.
func (t *TeacherControllerSample) AssignSectionRep(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sectionID := c.Param("section_id")
	type StudentID struct {
		StudentID string `json:"student_id" binding:"required"`
	}
	var studentID StudentID
	if sectionID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	if err := c.BindJSON(&studentID); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	section, er := TeacherService.GetSectionByID(sectionID)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	if section.HomeTeacherID != claim.(*token.SignedDetails).Uid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only section teacher can assign section rep"})
		return
	}
	usecase_err := TeacherService.AssignSectionRep(sectionID, studentID.StudentID)
	if usecase_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": usecase_err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully assigned section rep"})

}

// CreateSection implements TeacherController.
func (t *TeacherControllerSample) CreateSection(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var section models.Section
	if err := c.BindJSON(&section); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(section)
	if validationErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	claim, err := c.Get("claims")
	if !err {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	section.HomeTeacherID = claim.(*token.SignedDetails).Uid
	section.Students = []string{}
	section.Subjects = []string{}
	section.Teachers = []models.TeacherMap{}
	section.SectionID = primitive.NewObjectID()

	er := TeacherService.CreateSection(section)
	if er != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully created section"})

}

// AssignLead implements TeacherController.
func (t *TeacherControllerSample) AssignLead(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	clubID := c.Param("club_id")
	type StudentID struct {
		StudentID string `json:"student_id" binding:"required"`
	}
	var studentID StudentID
	if clubID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "club_id is required"})
		return
	}
	if err := c.BindJSON(&studentID); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}
	clientToken := c.GetHeader("Authorization")
	splitToken := strings.Split(clientToken, "Bearer ")
	clientToken = splitToken[1]

	claim, _ := token.ValidateToken(clientToken)
	if claim.Role != "teacher" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only Teachers can assign lead"})
		return
	}

	teacherID := claim.Uid

	clubData, err := Service.GetClubByID(clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if clubData.ClubTeacherID != teacherID {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only club teacher can assign lead"})
		return
	}
	err = TeacherService.AssignLead(clubID, studentID.StudentID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully assigned lead"})
}

func (t *TeacherControllerSample) FoundClub(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var club models.Club
	err := c.BindJSON(&club)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clientToken := c.GetHeader("Authorization")
	splitToken := strings.Split(clientToken, "Bearer ")
	clientToken = splitToken[1]

	claim, _ := token.ValidateToken(clientToken)
	if claim.Role != "teacher" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "only Teachers can found club"})
		return
	}

	teacherID := claim.Uid
	club.ClubTeacherID = teacherID
	club.ClubID = primitive.NewObjectID()
	err = TeacherService.FoundClub(club)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully founded club"})

}

func NewTeacherController(service usecase.TeacherUseCase) TeacherController {
	TeacherService = service
	return &TeacherControllerSample{}
}

package controller

import (
	"net/http"
	"schoolbackend/models"
	"schoolbackend/token"
	"schoolbackend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	Service  usecase.CommonUsecases
)

type CommonController interface {
	RegisterStudent(c *gin.Context)
	RegisterTeacher(c *gin.Context)
	RegisterParent(c *gin.Context)
	LoginStudent(c *gin.Context)
	LoginTeacher(c *gin.Context)
	LoginParent(c *gin.Context)

	GetClubs(c *gin.Context)
	GetClubByID(c *gin.Context)
	ApplyForClub(c *gin.Context)
	AcceptClubRequest(c *gin.Context)
	RejectClubRequest(c *gin.Context)
	GetClubApplications(c *gin.Context)

	GetAllSections(c *gin.Context)
	GetSectionStudents(c *gin.Context)
	GetSectionTeachers(c *gin.Context)
}

type CommonControllerSample struct{}

// GetAllSections implements CommonController.
func (*CommonControllerSample) GetAllSections(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	sections, err := Service.GetAllSections()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, sections)
}
// GetSectionTeachers implements CommonController.
func (*CommonControllerSample) GetSectionTeachers(c *gin.Context) {
	section_id := c.Param("section_id")
	if section_id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	students, err := Service.GetSectionTeachers(section_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}


// GetSectionStudents implements CommonController.
func (*CommonControllerSample) GetSectionStudents(c *gin.Context) {
	section_id := c.Param("section_id")
	if section_id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "section_id is required"})
		return
	}
	students, err := Service.GetSectionStudents(section_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

// AcceptClubRequest implements CommonController.
func (*CommonControllerSample) AcceptClubRequest(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	clubID := c.Param("club_id")
	applicantID := c.Param("student_id")
	if clubID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "club_id is required"})
		return
	}
	if applicantID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}
	clientToken := c.GetHeader("Authorization")
	splitToken := strings.Split(clientToken, "Bearer ")
	clientToken = splitToken[1]

	claim, _ := token.ValidateToken(clientToken)

	studentID := claim.Uid
	clubData, err := Service.GetClubByID(clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if clubData.ClubStudentRepID != studentID {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Only the club leader can view applications"})
		return
	}
	err = Service.AcceptRequest(applicantID, clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully accepted request"})
}

// GetClubApplications implements CommonController.
func (*CommonControllerSample) GetClubApplications(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	clubID := c.Param("club_id")
	if clubID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "club_id is required"})
		return
	}
	clientToken := c.GetHeader("Authorization")
	splitToken := strings.Split(clientToken, "Bearer ")
	clientToken = splitToken[1]

	claim, _ := token.ValidateToken(clientToken)

	studentID := claim.Uid
	clubData, err := Service.GetClubByID(clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if clubData.ClubStudentRepID != studentID {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Only the club leader can view applications"})
		return
	}
	applicants, err := Service.GetClubApplications(clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, applicants)
}

// RejectClubRequest implements CommonController.
func (*CommonControllerSample) RejectClubRequest(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	clubID := c.Param("club_id")
	applicantID := c.Param("student_id")
	if clubID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "club_id is required"})
		return
	}
	if applicantID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}
	clientToken := c.GetHeader("Authorization")
	splitToken := strings.Split(clientToken, "Bearer ")
	clientToken = splitToken[1]

	claim, _ := token.ValidateToken(clientToken)

	studentID := claim.Uid
	clubData, err := Service.GetClubByID(clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if clubData.ClubStudentRepID != studentID {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Only the club leader can view applications"})
		return
	}
	err = Service.RejectRequest(applicantID, clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "rejected request"})
}

// LoginParent implements CommonController.
func (*CommonControllerSample) LoginParent(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}
	parent, err := Service.LoginParent(email, password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, parent)
}

// LoginStudent implements CommonController.
func (*CommonControllerSample) LoginStudent(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.Student
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if user.Email == nil || user.Password == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}
	student, err := Service.LoginStudent(*user.Email, *user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

// LoginTeacher implements CommonController.
func (*CommonControllerSample) LoginTeacher(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.Teacher
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if user.Email == nil || user.Password == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}
	teacher, err := Service.LoginTeacher(*user.Email, *user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, teacher)

}

// RegisterParent implements CommonController.
// check if they are valid parents
func (*CommonControllerSample) RegisterParent(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var parent models.Parent
	err := c.BindJSON(&parent)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Service.RegisterParent(parent)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully registered parent"})
}

// RegisterTeacher implements CommonController.
func (*CommonControllerSample) RegisterTeacher(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var teacher models.Teacher

	err := c.BindJSON(&teacher)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Service.RegisterTeacher(teacher)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully registered teacher"})
}

// RegisterStudent implements CommonController.
func (*CommonControllerSample) RegisterStudent(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "please send valid student data"})
		return
	}
	validationErr := validate.Struct(student)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	err = Service.RegisterStudent(student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully registered student"})
}

// ApplyForClub implements CommonController.
func (*CommonControllerSample) ApplyForClub(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	clubID := c.Param("club_id")
	if clubID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "club_id is required"})
		return
	}
	clientToken := c.GetHeader("Authorization")
	splitToken := strings.Split(clientToken, "Bearer ")
	clientToken = splitToken[1]

	claim, _ := token.ValidateToken(clientToken)
	if claim.Role != "student" && claim.Role != "teacher" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Only Students and Teachers can apply for club"})
		return
	}

	studentID := claim.Uid
	err := Service.ApplyClub(studentID, clubID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully applied for club"})

}

// GetClubByID implements CommonController.
func (*CommonControllerSample) GetClubByID(c *gin.Context) {

	clubID := c.Param("club_id")
	if clubID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "club_id is required"})
		return
	}
	club, err := Service.GetClubByID(clubID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, club)
}

// GetClubs implements CommonController.
func (*CommonControllerSample) GetClubs(c *gin.Context) {
	clubs, err := Service.GetClubs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, clubs)
}

func NewCommonController(service usecase.CommonUsecases) CommonController {
	Service = service
	return &CommonControllerSample{}
}

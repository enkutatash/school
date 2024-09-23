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
	Service usecase.CommonUsecases
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
}

type CommonControllerSample struct{}

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
	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
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
    if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":  err.Error()})
		return
	}
	err = Service.RegisterTeacher(teacher)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":  err.Error()})
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":  err.Error()})
		return
	}
	validationErr := validate.Struct(student)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	err = Service.RegisterStudent(student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":  err.Error()})
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
	if claim.Role != "student" || claim.Role == "teacher" {
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

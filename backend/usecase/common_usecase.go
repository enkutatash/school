package usecase

import (
	"errors"
	"log"
	"schoolbackend/database"
	"schoolbackend/models"
	"schoolbackend/token"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type CommonUsecases interface {
	RegisterStudent(student models.Student) error
	RegisterTeacher(teacher models.Teacher) error
	RegisterParent(parent models.Parent) error
	LoginStudent(email string, password string) (*models.Student, error)
	LoginTeacher(email string, password string) (*models.Teacher, error)
	LoginParent(email string, password string) (*models.Parent, error)
	GetStudentGrade(studentID string) (models.GradeReport, error)
	GetStudentAttendance(studentID string) (models.StudentAttendance, error)
	GetClubs() ([]models.Club, error)
	GetClubByID(clubID string) (*models.Club, error)

	//
	ApplyClub(studentID string, clubID string) error
	// should be club leader
	AcceptRequest(studentID string, clubID string) error
	RejectRequest(studentID string, clubID string) error
	GetClubApplications(clubID string) ([]models.Student, error)
	SendNotification(notification models.Notification) error
	
}
type UsecaseSample struct{}



// AcceptRequest implements CommonUsecases.
func (u *UsecaseSample) AcceptRequest(studentID string, clubID string) error {
	err := database.AcceptRequest(studentID, clubID)
	if err != nil {
		return err
	}
	return nil
}

// GetClubApplications implements CommonUsecases.
func (u *UsecaseSample) GetClubApplications(clubID string) ([]models.Student, error) {
	applicants, err := database.GetAllApplicant(clubID)
	if err != nil {
		return nil, err
	}
	return applicants, nil
}

// RejectRequest implements CommonUsecases.
func (u *UsecaseSample) RejectRequest(studentID string, clubID string) error {
	err := database.RejectRequest(studentID, clubID)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hashed)
}

func VerifyPassword(hashedPassword, password string) bool {
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return bcryptErr == nil
}

// LoginParent implements CommonUsecases.
func (u *UsecaseSample) LoginParent(email string, password string) (*models.Parent, error) {
	panic("unimplemented")
}

// LoginStudent implements CommonUsecases.
func (u *UsecaseSample) LoginStudent(email string, password string) (*models.Student, error) {

	foundUser, err := database.LoginStudent(email)
	if err != nil {
		return nil, errors.New("student not found")
	}

	if !VerifyPassword(*foundUser.Password, password) {
		return nil, errors.New("invalid password")
	}
	return foundUser, nil

}

// LoginTeacher implements CommonUsecases.
func (u *UsecaseSample) LoginTeacher(email string, password string) (*models.Teacher, error) {
	foundUser, err := database.LoginTeacher(email)
	if err != nil {
		return nil, errors.New("teacher not found")
	}

	if !VerifyPassword(*foundUser.Password, password) {
		return nil, errors.New("invalid password")
	}
	return foundUser, nil
}

// RegisterParent implements CommonUsecases.
func (u *UsecaseSample) RegisterParent(parent models.Parent) error {
	panic("unimplemented")
}

// RegisterStudent implements CommonUsecases.
func (u *UsecaseSample) RegisterStudent(student models.Student) error {
	valid := database.ValidStudent(*student.FirstName, *student.LastName)
	if !valid {
		return errors.New("no Student on this name")
	}

	email := student.Email

	exist := database.EmailExistStudent(*email)
	if exist {
		return errors.New("email already exist")
	}
	student.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	student.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	student.StudentID = primitive.NewObjectID()
	uid := student.StudentID.Hex()
	studenttoken, refreshToken, err := token.GenerateToken(*student.Email, *student.FirstName, uid, *student.LastName, "student")
	if err != nil {
		return err
	}
	password := HashPassword(*student.Password)
	student.Password = &password
	student.Token = &studenttoken
	student.RefreshToken = &refreshToken
	err = database.RegisterStudent(student)
	if err != nil {
		return err
	}
	return nil
}

// RegisterTeacher implements CommonUsecases.
func (u *UsecaseSample) RegisterTeacher(teacher models.Teacher) error {
	valid := database.ValidTeacher(*teacher.FirstName, *teacher.LastName)
	if !valid {
		return errors.New("no teacher on this name")
	}

	email := teacher.Email

	exist := database.EmailExistTeacher(*email)
	if exist {
		return errors.New("email already Exist")
	}
	teacher.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	teacher.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	teacher.TeacherID = primitive.NewObjectID()
	uid := teacher.TeacherID.Hex()
	teachertoken, refreshToken, err := token.GenerateToken(*teacher.Email, *teacher.FirstName, uid, *teacher.LastName, "teacher")
	if err != nil {
		return err
	}
	password := HashPassword(*teacher.Password)
	teacher.Password = &password
	teacher.Token = &teachertoken
	teacher.RefreshToken = &refreshToken
	err = database.RegisterTeacher(teacher)
	if err != nil {
		return err
	}
	return nil
}

// ApplyClub implements CommonUsecases.
func (u *UsecaseSample) ApplyClub(studentID string, clubID string) error {
	exist := database.CheckStudent(studentID)
	if !exist {
		return errors.New("please register first")
	}
	err := database.ApplyClub(studentID, clubID)
	if err != nil {
		return err
	}
	return nil
}

// GetClubByID implements CommonUsecases.
func (u *UsecaseSample) GetClubByID(clubID string) (*models.Club, error) {
	club, err := database.GetClubByID(clubID)
	if err != nil {
		return club, err
	}
	return club, nil
}

// GetClubs implements CommonUsecases.
func (u *UsecaseSample) GetClubs() ([]models.Club, error) {
	clubs, err := database.GetClubs()
	if err != nil {
		return nil, err
	}
	return clubs, nil
}

// GetStudentAttendance implements CommonUsecases.
func (u *UsecaseSample) GetStudentAttendance(studentID string) (models.StudentAttendance, error) {
	panic("unimplemented")
}

// GetStudentGrade implements CommonUsecases.
func (u *UsecaseSample) GetStudentGrade(studentID string) (models.GradeReport, error) {
	panic("unimplemented")
}

// SendNotification implements CommonUsecases.
func (u *UsecaseSample) SendNotification(notification models.Notification) error {
	panic("unimplemented")
}

func NewCommonUsecase() CommonUsecases {
	return &UsecaseSample{}
}

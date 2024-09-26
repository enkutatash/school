package database

import (
	"context"
	"errors"
	"schoolbackend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterStudent(student models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	student.Clubs = []string{}

	_, inserErr := RegisteredStudent.InsertOne(ctx, student)

	if inserErr != nil {
		return inserErr
	}
	var grade models.GradeReport
	grade.StudentID = student.StudentID.Hex()
	grade.Result = []models.SubjectAssessment{}
	_,inertErr := GradeReportData.InsertOne(ctx,grade)
	if inertErr != nil{
		return inertErr
	}
	return nil
}

func RegisterTeacher(teacher models.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, inserErr := RegisteredTeacher.InsertOne(ctx, teacher)

	if inserErr != nil {
		return inserErr
	}
	
	return nil
}


func LoginStudent(email string)(*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser models.Student
	err := RegisteredStudent.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
	if err != nil {
		return nil,errors.New("user not found")
	}
	return &foundUser,nil

}

func LoginTeacher(email string)(*models.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser models.Teacher
	err := RegisteredTeacher.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
	if err != nil {
		return nil,errors.New("user not found")
	}
	return &foundUser,nil

}



// section

func GetAllSections() ([]models.Section, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	projection := bson.M{
		"_id":  1,
		"grade":1,
		"section_name": 1,
		"teacher": 1,
		"student_rep":  1,
	}
	cursor, err := SectionsData.Find(ctx, bson.M{}, options.Find().SetProjection(projection))

	if err != nil {
		return nil, err
	}

	var sections []models.Section
	if err = cursor.All(ctx, &sections); err != nil {
		return nil, err
	}

	return sections, nil
}

func GetSectionStudents(sectionID string)([]models.Student,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sectionid , err := primitive.ObjectIDFromHex(sectionID)
	if err != nil{
		return nil,err	
	}
	var students []models.Student
	var section models.Section

	err = SectionsData.FindOne(ctx,bson.M{"_id":sectionid}).Decode(&section)
	if err != nil{
		return nil,err
	}


	for _,studentID := range section.Students{
		var student models.Student
		student, err = GetStudentByID(studentID)
		if err != nil{
			return nil,err
		}
		students = append(students,student)
	}
	return students,nil
	
}

func GetSectionTeachers(sectionID string)(*map[string]models.Teacher,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sectionid , err := primitive.ObjectIDFromHex(sectionID)
	if err != nil{
		return nil,err	
	}
	 teachers:=  make(map[string]models.Teacher)
	var section models.Section

	err = SectionsData.FindOne(ctx,bson.M{"_id":sectionid}).Decode(&section)
	if err != nil{
		return nil,err
	}


	for _,teacherID := range section.Teachers{
		var teacher models.Teacher
		teacher, err = GetTeacherByID(teacherID.TeacherID)
		if err != nil{
			return nil,err
		}
		teacher.Token = nil
		teacher.Password = nil
		teacher.RefreshToken = nil
		teacher.CreatedAt = time.Time{}
		teacher.UpdatedAt = time.Time{}


		subject, err := GetSubjectByID(teacherID.SubjectID)
		if err != nil{
			return nil,err
		}
		teachers[subject.SubjectName] = teacher
	}
	return &teachers,nil
	
}
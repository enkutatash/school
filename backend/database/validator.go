package database

import (
	"context"
	"errors"
	"schoolbackend/models"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
    StudentData        = GetData(Client, "studentslist")
    RegisteredStudent  = GetData(Client, "registeredstudents")
    TeacherData        = GetData(Client, "teacherslist")
	RegisteredTeacher  = GetData(Client, "registeredteachers")
	Clubs = GetData(Client, "clubs")
	SectionsData = GetData(Client, "sections")
	SubjectsData = GetData(Client,"subjects")
	GradeReportData = GetData(Client,"gradereport")
	
)

func EmailExistStudent(email string) bool {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    count, err := RegisteredStudent.CountDocuments(ctx, bson.M{"email": email})
    if err != nil {
		
		log.Println("Error counting documents:", err) 
        return false  
    }
	
    return count > 0
}

func ValidStudent(FirstName string,LastName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := StudentData.CountDocuments(ctx, bson.M{"first_name": FirstName,"last_name": LastName})
	if err != nil {
		log.Panic(err)
		return false  
	}

	return count > 0
}

func EmailExistTeacher(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := RegisteredTeacher.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		log.Panic(err)
		return false  
	}
	return count > 0
}

func ValidTeacher(FirstName string,LastName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := TeacherData.CountDocuments(ctx, bson.M{"first_name": FirstName,"last_name": LastName})
	if err != nil {
		log.Panic(err)
		return false  
	}

	return count > 0
}


func GetSectionByID(sectionID string)(models.Section,error){
	sectionid, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return models.Section{},err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var section models.Section
	err = SectionsData.FindOne(ctx, bson.M{"_id": sectionid}).Decode(&section)
	if err != nil {
		return models.Section{},err
	}
	return section,nil

}


func GetStudentByID(studentID string)(models.Student,error){
	studentid, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		return models.Student{},err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var student models.Student
	err = RegisteredStudent.FindOne(ctx, bson.M{"_id": studentid}).Decode(&student)
	if err != nil {
		return models.Student{},err
	}
	return student,nil
}

func GetTeacherByID(teacherID string)(models.Teacher,error){
	teacherid, err := primitive.ObjectIDFromHex(teacherID)
	if err != nil {
		return models.Teacher{},err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var teacher models.Teacher
	err = RegisteredTeacher.FindOne(ctx, bson.M{"_id": teacherid}).Decode(&teacher)
	if err != nil {
		return models.Teacher{},err
	}
	return teacher,nil
}
func ValidSubject(subjectID string) bool{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	subjectid ,err := primitive.ObjectIDFromHex(subjectID)
	if err != nil {
		return false
	}
	count, err := SubjectsData.CountDocuments(ctx, bson.M{"_id": subjectid})
	if err != nil {
		log.Panic(err)
		return false  
	}

	return count > 0
}

func GetSubjectByID(subjectID string) (*models.Subject,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	subjectid ,err := primitive.ObjectIDFromHex(subjectID)
	if err != nil {
		return nil,err
	}
	var subject models.Subject
	err = SubjectsData.FindOne(ctx, bson.M{"_id": subjectid}).Decode(&subject)
	if err != nil {
		
		return nil,err
	}

	return &subject,nil

}

func ValidTeacherToSection(sectionID string,teacherID string) (*string ,error){
	sectionid, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return nil,err
	}
	
	filter := bson.M{
		"_id": sectionid,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var section models.Section
	err = SectionsData.FindOne(ctx, filter).Decode(&section)
	if err != nil {
		return nil,err
	}
	for _, teacher := range section.Teachers {
		if teacher.TeacherID == teacherID {
			return &teacher.SubjectID,nil
		}
	}
	return nil, errors.New("teacher not found in section")

}
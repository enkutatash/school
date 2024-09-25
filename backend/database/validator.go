package database

import (
	"context"
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
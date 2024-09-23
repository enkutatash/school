package database

import (
	"context"
	
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
    StudentData        = GetData(Client, "studentslist")
    RegisteredStudent  = GetData(Client, "registeredstudents")
    TeacherData        = GetData(Client, "teacherslist")
	RegisteredTeacher  = GetData(Client, "registeredteachers")
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

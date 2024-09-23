package database

import (
	"context"
	"errors"
	"schoolbackend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func RegisterStudent(student models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, inserErr := RegisteredStudent.InsertOne(ctx, student)
	if inserErr != nil {
		return inserErr
	}
	return nil
}

func RegisterTeacher(teacher models.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, inserErr := TeacherData.InsertOne(ctx, teacher)
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

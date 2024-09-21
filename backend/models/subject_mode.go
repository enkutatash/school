package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subject struct {
	SubjectID primitive.ObjectID `json:"_id" bson:"_id"`
	SubjectName string `json:"subject_name" bson:"subject_name"`
	TeacherID []string `json:"teacher" bson:"teacher"`
	Grade int `json:"grade" bson:"grade"`
	DepartmentLeadID string `json:"department_lead" bson:"department_lead"`
}
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Section struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	SectionName   string             `json:"section_name" bson:"section_name"`
	Grade         int                `json:"grade" bson:"grade"`
	HomeTeacherID string             `json:"teacher" bson:"teacher"`
	StudentRepID  string             `json:"student_rep" bson:"student_rep"`
	Students 	[]string           `json:"students" bson:"students"`
}

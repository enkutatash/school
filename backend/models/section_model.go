package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Section struct {
	SectionID     primitive.ObjectID `json:"_id" bson:"_id"`
	SectionName   string             `json:"section_name" bson:"section_name" validate:"required,min=2,max=30"`
	Grade         int                `json:"grade" bson:"grade" validate:"required,min=9,max=12"`
	HomeTeacherID string             `json:"teacher" bson:"teacher"`
	StudentRepID  string             `json:"student_rep" bson:"student_rep"`
	Students      []string           `json:"students" bson:"students"`
	Teachers      []TeacherMap       `json:"teachers" bson:"teachers"`
	Subjects      []string           `json:"subjects" bson:"subjects"`
}

type TeacherMap struct {
	TeacherID string `json:"teacher_id" bson:"teacher_id"`
	SubjectID string `json:"subject_id" bson:"subject_id"`
}

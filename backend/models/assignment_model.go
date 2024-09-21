package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	AssignmentID   primitive.ObjectID `json:"_id" bson:"_id"`
	SubjectID      *string             `json:"subject" bson:"subject"`
	TeacherID 	*string             `json:"teacher" bson:"teacher"`
	Deadline       time.Time             `json:"deadline" bson:"deadline"`
	AssignmentFile []byte			  `json:"assignment_file" bson:"assignment_file"`
}
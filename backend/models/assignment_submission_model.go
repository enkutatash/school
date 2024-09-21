package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssignmentSubmission struct {
	SubmissionID primitive.ObjectID `json:"_id" bson:"_id"`
	AssignmentID *string             `json:"assignment" bson:"assignment"`
	StudentsID []string 		   `json:"students" bson:"students"`
	SubmissionFile []byte			  `json:"submission_file" bson:"submission_file"`
	SubmissionDescription string `json:"submission_description" bson:"submission_description"`
}
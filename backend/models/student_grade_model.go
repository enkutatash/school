package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GradeReport struct {
	StudentID string              `json:"student_id" bson:"student_id"`
	Result    []SubjectAssessment `json:"subjects_result" bson:"subjects_result"`
}

type SubjectAssessment struct {
	SubjectID string `json:"subject_id" bson:"subject_id"`
	Assessment []AssessmentType `json:"assessment" bson:"assessment"`
}

type AssessmentType struct {
	AssessmentTypeID primitive.ObjectID    `json:"assessment_type_id" bson:"assessment_type_id"`
	AssessmentName   string    `json:"assessment_name" bson:"assessment_name"`
	AssessmentWeight int       `json:"assessment_weight" bson:"assessment_weight"`
	AssessmentValue  float64   `json:"assessment_value" bson:"assessment_value"`
	AssessmentDate   time.Time `json:"assessment_date" bson:"assessment_date"`
}
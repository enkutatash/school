package models

type GradeReport struct {
	StudentID string `json:"_id" bson:"_id"`
	Subjects  map[string]float64     `json:"subjects_result" bson:"subjects_result"`

}
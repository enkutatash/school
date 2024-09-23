package models

type Question struct {
	QuestionID string `json:"_id" bson:"_id"`
	TypeOfQuestion string `json:"type_of_question" bson:"type_of_question"`
	Question string `json:"question" bson:"question"`
	Options []string `json:"options" bson:"options"`
}
type Quiz struct {
	QuizID string `json:"_id" bson:"_id"`
	SubjectID string `json:"subject_id" bson:"subject_id"`
	Questions []Question `json:"questions" bson:"questions"`
	TeacherID string `json:"teacher_id" bson:"teacher_id"`
	QuizDate string `json:"quiz_date" bson:"quiz_date"`
	QuizDuration string `json:"quiz_duration" bson:"quiz_duration"`
	


}
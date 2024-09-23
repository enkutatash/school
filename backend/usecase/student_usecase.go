package usecase

import "schoolbackend/models"

type StudentUseCase interface {
	GetAssignments(SubjectID string) (models.Assignment, error)
	SubmitAssignment(studentID []string, assignmentID string, assignmentFile []byte) error // by token
	TakeQuiz(quizID string, studentID string) error
	GetQuiz(quizID string) (models.Quiz, error)
}

type StudentUseCaseSample struct{}

// GetAssignments implements StudentUseCase.
func (s *StudentUseCaseSample) GetAssignments(SubjectID string) (models.Assignment, error) {
	panic("unimplemented")
}

// GetQuiz implements StudentUseCase.
func (s *StudentUseCaseSample) GetQuiz(quizID string) (models.Quiz, error) {
	panic("unimplemented")
}

// SubmitAssignment implements StudentUseCase.
func (s *StudentUseCaseSample) SubmitAssignment(studentID []string, assignmentID string, assignmentFile []byte) error {
	panic("unimplemented")
}

// TakeQuiz implements StudentUseCase.
func (s *StudentUseCaseSample) TakeQuiz(quizID string, studentID string) error {
	panic("unimplemented")
}

func NewStudentUsecase() StudentUseCase {
	return &StudentUseCaseSample{}
}

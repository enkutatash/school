package usecase

import (
	"schoolbackend/models"
	"time"
)

type TeacherUseCase interface {
	GetAssignments(SubjectID string, Grade int) (models.Assignment, error)
	SubmittedAssignment(assignmentID string, Grade int) (models.Assignment, error)
	GradeAssignment(assignmentID string, studentID string, grade int, Grade int) error
	PostAssignment(assignment models.Assignment, Grade int) error
	EditAssignment(assignmentID string, assignment models.Assignment, Grade int) error // check assignment poster is that teacher
	DeleteAssignment(assignmentID string, Grade int) error                             // check assignment poster is that teacher
	CreateSection(section models.Section, Grade int) error
	EditSection(sectionID string, section models.Section, Grade int) error
	DeleteSection(sectionID string, Grade int) error
	AddStudentToSection(sectionID string, studentID string, Grade int) error      // check if that is their section
	RemoveStudentFromSection(sectionID string, studentID string, Grade int) error // check if that is their section
	GetAttendence(SubjectID string, SectionID string, Grade int) (models.StudentAttendance, error)
	MarkAttendence(SubjectID string, SectionID string, studentID string, mark models.AttendanceMark, timeStamp *time.Time) error
	PostQuiz(quiz models.Quiz, Grade int, sectionID []string) error
	EditQuiz(quizID string, quiz models.Quiz, Grade int) error // check the poster
	DeleteQuiz(quizID string, Grade int) error                 // check the poster
	GradeQuiz(quizID string, studentID string, grade int) error
}

type TeacherUseCaseSample struct{}

// AddStudentToSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) AddStudentToSection(sectionID string, studentID string, Grade int) error {
	panic("unimplemented")
}

// CreateSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) CreateSection(section models.Section, Grade int) error {
	panic("unimplemented")
}

// DeleteAssignment implements TeacherUseCase.
func (t *TeacherUseCaseSample) DeleteAssignment(assignmentID string, Grade int) error {
	panic("unimplemented")
}

// DeleteQuiz implements TeacherUseCase.
func (t *TeacherUseCaseSample) DeleteQuiz(quizID string, Grade int) error {
	panic("unimplemented")
}

// DeleteSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) DeleteSection(sectionID string, Grade int) error {
	panic("unimplemented")
}

// EditAssignment implements TeacherUseCase.
func (t *TeacherUseCaseSample) EditAssignment(assignmentID string, assignment models.Assignment, Grade int) error {
	panic("unimplemented")
}

// EditQuiz implements TeacherUseCase.
func (t *TeacherUseCaseSample) EditQuiz(quizID string, quiz models.Quiz, Grade int) error {
	panic("unimplemented")
}

// EditSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) EditSection(sectionID string, section models.Section, Grade int) error {
	panic("unimplemented")
}

// GetAssignments implements TeacherUseCase.
func (t *TeacherUseCaseSample) GetAssignments(SubjectID string, Grade int) (models.Assignment, error) {
	panic("unimplemented")
}

// GetAttendence implements TeacherUseCase.
func (t *TeacherUseCaseSample) GetAttendence(SubjectID string, SectionID string, Grade int) (models.StudentAttendance, error) {
	panic("unimplemented")
}

// GradeAssignment implements TeacherUseCase.
func (t *TeacherUseCaseSample) GradeAssignment(assignmentID string, studentID string, grade int, Grade int) error {
	panic("unimplemented")
}

// GradeQuiz implements TeacherUseCase.
func (t *TeacherUseCaseSample) GradeQuiz(quizID string, studentID string, grade int) error {
	panic("unimplemented")
}

// MarkAttendence implements TeacherUseCase.
func (t *TeacherUseCaseSample) MarkAttendence(SubjectID string, SectionID string, studentID string, mark models.AttendanceMark, timeStamp *time.Time) error {
	panic("unimplemented")
}

// PostAssignment implements TeacherUseCase.
func (t *TeacherUseCaseSample) PostAssignment(assignment models.Assignment, Grade int) error {
	panic("unimplemented")
}

// PostQuiz implements TeacherUseCase.
func (t *TeacherUseCaseSample) PostQuiz(quiz models.Quiz, Grade int, sectionID []string) error {
	panic("unimplemented")
}

// RemoveStudentFromSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) RemoveStudentFromSection(sectionID string, studentID string, Grade int) error {
	panic("unimplemented")
}

// SubmittedAssignment implements TeacherUseCase.
func (t *TeacherUseCaseSample) SubmittedAssignment(assignmentID string, Grade int) (models.Assignment, error) {
	panic("unimplemented")
}

func NewTeacherUsecase() TeacherUseCase {
	return &TeacherUseCaseSample{}
}

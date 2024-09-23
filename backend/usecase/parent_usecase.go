package usecase

import "schoolbackend/models"

type ParentUseCase interface {
	GetStudentGrade(studentID string) (models.GradeReport, error)
	GetStudentAttendance(studentID string) (models.StudentAttendance, error)
	GetStudentHomeTeacher(studentID string) (models.Teacher, error)
}

type ParentusecaseSample struct{}

// GetStudentAttendance implements ParentUseCase.
func (p *ParentusecaseSample) GetStudentAttendance(studentID string) (models.StudentAttendance, error) {
	panic("unimplemented")
}

// GetStudentGrade implements ParentUseCase.
func (p *ParentusecaseSample) GetStudentGrade(studentID string) (models.GradeReport, error) {
	panic("unimplemented")
}

// GetStudentHomeTeacher implements ParentUseCase.
func (p *ParentusecaseSample) GetStudentHomeTeacher(studentID string) (models.Teacher, error) {
	panic("unimplemented")
}

func NewParentUsecase() ParentUseCase {
	return &ParentusecaseSample{}
}

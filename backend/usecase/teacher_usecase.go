package usecase

import (
	"errors"
	"schoolbackend/database"
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
	CreateSection(section models.Section) error
	GetSectionByID(sectionID string) (models.Section, error)
	AssignSectionRep(sectionID string, studentID string) error
	EditSection(sectionID string, section models.Section, Grade int) error
	DeleteSection(sectionID string, Grade int) error
	AddStudentToSection(sectionID string, studentID []string) ([]models.Student, error) // check if that is their section
	RemoveStudentFromSection(sectionID string, studentID string) error                  // check if that is their section
	GetAttendence(SubjectID string, SectionID string, Grade int) (models.StudentAttendance, error)
	MarkAttendence(SubjectID string, SectionID string, studentID string, mark models.AttendanceMark, timeStamp *time.Time) error
	PostQuiz(quiz models.Quiz, Grade int, sectionID []string) error
	EditQuiz(quizID string, quiz models.Quiz, Grade int) error // check the poster
	DeleteQuiz(quizID string, Grade int) error                 // check the poster
	GradeQuiz(quizID string, studentID string, grade int) error

	FoundClub(club models.Club) error
	AssignLead(clubID string, studentID string) error

	AssignTeacher(subjectID string, TeacherID string, sectionID string) error
	UnAssignTeacher(TeacherID string, sectionID string) error
	AddCoursesToSection(subjectsID []string, sectionID string) error

	// teacher of the section

	NewAssessment(assessment models.AssessmentType, sectionID string, teacherID string) error
	GetSectionResult(sectionID string, subjectID string) (*map[string][]models.AssessmentType, error)
	UpdateSectionResult(sectionID string, studentID string, assessment_id string, assessmentValue float64,teacherID string) error
}

type TeacherUseCaseSample struct{}

// UpdateSectionResult implements TeacherUseCase.
func (t *TeacherUseCaseSample) UpdateSectionResult(sectionID string, studentID string, assessment_id string, assessmentValue float64,teacherID string) error {
	validSection := database.CheckSection(sectionID)
	if !validSection {
		return errors.New("invalid section id")
	}

	validStudent := database.CheckStudent(studentID)
	if !validStudent {
		return errors.New("invalid student id")
	}

	validStudent = database.CheckSectionRep(sectionID,studentID)
	if !validStudent{
		return errors.New("student is not in this class")
	}

	subjectID, err := database.ValidTeacherToSection(sectionID, teacherID)
	if err != nil {
		return err
	}
	err = database.UpdateSectionResult( sectionID,studentID, assessment_id, assessmentValue,*subjectID)
	if err != nil {
		return err
	}
	return nil
}

// GetSectionResult implements TeacherUseCase.
func (t *TeacherUseCaseSample) GetSectionResult(sectionID string, teacherID string) (*map[string][]models.AssessmentType, error) {
	validSection := database.CheckSection(sectionID)
	if !validSection {
		return nil, errors.New("invalid section id")
	}

	subjectID, err := database.ValidTeacherToSection(sectionID, teacherID)
	if err != nil {
		return nil, err
	}
	report, err := database.GetStudentsGradeReport(sectionID, *subjectID)
	if err != nil {
		return nil, err
	}
	return &report, nil

}

// NewAssessment implements TeacherUseCase.
func (t *TeacherUseCaseSample) NewAssessment(assessment models.AssessmentType, sectionID string, teacherID string) error {
	validSection := database.CheckSection(sectionID)
	if !validSection {
		return errors.New("invalid section id")
	}

	subjectID, err := database.ValidTeacherToSection(sectionID, teacherID)
	if err != nil {
		return err
	}

	err = database.NewAssessment(assessment, sectionID, *subjectID)
	if err != nil {
		return err
	}
	return nil

}

func (t *TeacherUseCaseSample) AddCoursesToSection(subjectsID []string, sectionID string) error {
	for _, id := range subjectsID {
		valid := database.ValidSubject(id)
		if !valid {
			return errors.New("invalid Subject id")
		}

		err := database.AddCourse(id, sectionID)

		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TeacherUseCaseSample) UnAssignTeacher(TeacherID string, sectionID string) error {
	checkTeacher := database.CheckTeacher(TeacherID)
	if !checkTeacher {
		return errors.New("invalid teacher id")
	}
	checkSection := database.CheckSection(sectionID)
	if !checkSection {
		return errors.New("invalid section id")
	}
	err := database.UnAssignTeacher(sectionID, TeacherID)
	if err != nil {
		return err
	}

	return nil
}

// AssignTeacher implements TeacherUseCase.
func (t *TeacherUseCaseSample) AssignTeacher(subjectID string, TeacherID string, sectionID string) error {
	checkSubject := database.ValidSubject(subjectID)
	if !checkSubject {
		return errors.New("invalid subject id")
	}
	checkTeacher := database.CheckTeacher(TeacherID)
	if !checkTeacher {
		return errors.New("invalid teacher id")
	}
	sectionData, err := database.GetSectionByID(sectionID)
	if err != nil {
		return errors.New("invalid section id")
	}
	for _, subject := range sectionData.Subjects {
		if subject == subjectID {
			err := database.AssignTeacher(subjectID, TeacherID, sectionID)
			// add if the id of the subject is same with the teacher next not implemented
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("subject not in the section")

}

// GetSectionByID implements TeacherUseCase.
func (t *TeacherUseCaseSample) GetSectionByID(sectionID string) (models.Section, error) {
	return database.GetSectionByID(sectionID)
}

// AssignSectionRep implements TeacherUseCase.
func (t *TeacherUseCaseSample) AssignSectionRep(sectionID string, studentID string) error {
	validStudent := database.CheckStudent(studentID)
	if !validStudent {
		return errors.New("student does not exist")
	}
	validRep := database.CheckSectionRep(sectionID, studentID)
	if !validRep {
		return errors.New("section rep should be from that section")
	}
	err := database.AssignSectionRep(sectionID, studentID)
	if err != nil {
		return err
	}
	return nil
}

// AssignLead implements TeacherUseCase.
func (t *TeacherUseCaseSample) AssignLead(clubID string, studentID string) error {
	validStudent := database.CheckStudent(studentID)
	if !validStudent {
		return errors.New("student does not exist")
	}
	err := database.AssignLead(clubID, studentID)
	if err != nil {
		return err
	}
	return nil
}

// FoundClub implements TeacherUseCase.
func (t *TeacherUseCaseSample) FoundClub(club models.Club) error {
	exist := database.CheckClub(club.ClubName)
	if exist {
		return errors.New("club already exist")
	}
	err := database.FoundClub(club)
	if err != nil {
		return err
	}
	return nil
}

// AddStudentToSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) AddStudentToSection(sectionID string, studentIDs []string) ([]models.Student, error) {
	var invalidStudents []models.Student

	for _, id := range studentIDs {
		validStudent := database.CheckStudent(id)
		if !validStudent {
			return nil, errors.New("student does not exist: " + id)
		}

		student, err := database.GetStudentByID(id)
		if err != nil {
			return nil, err
		}

		if student.SectionID != nil {
			// Add to the list of invalid students
			invalidStudents = append(invalidStudents, student)
		} else {
			err = database.AddStudentToSection(sectionID, id)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(invalidStudents) > 0 {
		return invalidStudents, errors.New("some students are already assigned to a section")
	}

	return nil, nil
}

// RemoveStudentFromSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) RemoveStudentFromSection(sectionID string, studentID string) error {
	validStudent := database.CheckStudent(studentID)
	if !validStudent {
		return errors.New("student does not exist: " + studentID)
	}

	student, err := database.GetStudentByID(studentID)
	if err != nil {
		return err
	}

	if *student.SectionID != sectionID {

		return errors.New("student is not assigned to this section")
	} else {
		err = database.RemoveStudentFromSection(sectionID, studentID)
		if err != nil {
			return err
		}
		return nil
	}
}

// CreateSection implements TeacherUseCase.
func (t *TeacherUseCaseSample) CreateSection(section models.Section) error {
	exist := database.CheckTeacher(section.HomeTeacherID)
	if exist {
		return errors.New("one teacher can be home teacher of only one section")
	}
	sectionID := section.SectionID.Hex()
	exist = database.CheckSection(sectionID)
	if exist {
		return errors.New("section already exist")
	}
	err := database.CreateSection(section)
	if err != nil {
		return errors.New("error in creating section")
	}
	return nil
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

// SubmittedAssignment implements TeacherUseCase.
func (t *TeacherUseCaseSample) SubmittedAssignment(assignmentID string, Grade int) (models.Assignment, error) {
	panic("unimplemented")
}

func NewTeacherUsecase() TeacherUseCase {
	return &TeacherUseCaseSample{}
}

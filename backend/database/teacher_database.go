package database

import (
	"context"
	"errors"
	"fmt"

	"schoolbackend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UnAssignTeacher(sectionID string, teacherID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sectionId, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return err
	}

	var section models.Section
	err = SectionsData.FindOne(ctx, bson.M{"_id": sectionId}).Decode(&section)
	if err != nil {
		return err // Handle error if section not found
	}

	// Create a new slice to hold the updated teachers
	var updatedTeachers []models.TeacherMap

	for _, teacher := range section.Teachers {
		if teacher.TeacherID != teacherID {
			// Add to the updated list if it is not the teacher to be removed
			updatedTeachers = append(updatedTeachers, teacher)
		}
	}

	// Update the section's teachers in the database
	_, err = SectionsData.UpdateOne(ctx, bson.M{"_id": sectionId}, bson.M{"$set": bson.M{"teachers": updatedTeachers}})
	return err
}


func AssignTeacher(subjectID string,teacherID string, sectionID string)error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// var section models.Section
	sectionId,err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return err
	}

	newTeacher := models.TeacherMap{
		TeacherID: teacherID,
		SubjectID: subjectID,
	}
	var section models.Section
	err = SectionsData.FindOne(ctx, bson.M{"_id": sectionId}).Decode(&section)
	if err != nil {
		return err // Handle error if section not found
	}

	// Check if the subject ID already exists
	var existingTeacher *models.TeacherMap
	for i, teacher := range section.Teachers {
		if teacher.SubjectID == newTeacher.SubjectID {
			existingTeacher = &section.Teachers[i]
			break
		}
	}

	if existingTeacher != nil {
		// If the teacher with the same SubjectID exists, update the TeacherID
		existingTeacher.TeacherID = newTeacher.TeacherID
	} else {
		// If it doesn't exist, push the new teacher to the array
		section.Teachers = append(section.Teachers, newTeacher)
	}

	// Update the section document
	_, err = SectionsData.UpdateOne(ctx, bson.M{"_id": sectionId}, bson.M{"$set": bson.M{"teachers": section.Teachers}})
	return err
	
}

func CheckSectionRep(sectionID string, studentID string) bool {
	studentid, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		return false
	}
	sectionid, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := SectionsData.CountDocuments(ctx, bson.M{"_id": sectionid, "students": studentid})
	if err != nil {
		return false
	}
	return count > 0
}

func AssignSectionRep(sectionID string, studentID string) error {
	sectionid, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = SectionsData.UpdateOne(ctx, bson.M{"_id": sectionid}, bson.M{"$set": bson.M{"student_rep": studentID}})
	if err != nil {
		return err
	}
	return nil
}

func CreateSection(section models.Section) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := SectionsData.InsertOne(ctx, section)
	if err != nil {
		return err
	}
	return nil
}

func CheckTeacher(teacherID string) bool {
	teacherId, err := primitive.ObjectIDFromHex(teacherID)
	if err != nil {
		return true

	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	count, err := RegisteredTeacher.CountDocuments(ctx, bson.M{"_id": teacherId})
	if err != nil {
		return true
	}
	return count > 0
}

func CheckSection(section string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sectionID,err := primitive.ObjectIDFromHex(section)
	if err != nil {
		return true
	}
	count, err := SectionsData.CountDocuments(ctx, bson.M{"_id": sectionID})
	if err != nil {
		return true
	}
	return count > 0
}

func AddStudentToSection(sectionID string, studentID string) error {
	studentid, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		return err
	}
	sectionid, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = SectionsData.UpdateOne(ctx, bson.M{"_id": sectionid}, bson.M{"$push": bson.M{"students": studentid}})
	if err != nil {
		return err
	}
	_,err = RegisteredStudent.UpdateOne(ctx,bson.M{"_id":studentid},bson.M{"$set":bson.M{"section":sectionid}})
	if err != nil {
		return err
	}
	return nil
}


func RemoveStudentFromSection(sectionID string, studentID string) error {
	studentid, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		return err
	}
	sectionid, err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var section models.Section
	_, err = SectionsData.UpdateOne(ctx, bson.M{"_id": sectionid}, bson.M{"$pull": bson.M{"students": studentid}})
	if err != nil {
		return err
	}

	_,err = RegisteredStudent.UpdateOne(ctx,bson.M{"_id":studentid},bson.M{"$set":bson.M{"section":nil}})
	if err != nil {
		return err
	}
	err = SectionsData.FindOne(ctx,bson.M{"_id":sectionid}).Decode(&section)
	if err != nil {
		return err
	}
	if section.StudentRepID == studentID {
		_,err = SectionsData.UpdateOne(ctx,bson.M{"_id":sectionid},bson.M{"$set":bson.M{"student_rep":nil}})
		if err != nil {
			return err
		}

	}

	return nil
}

func AddCourse(subjectID string,sectionID string) error{
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var section models.Section
	sectionid ,err := primitive.ObjectIDFromHex(sectionID)
	if err != nil {
		return err
	}
	
	err = SectionsData.FindOne(ctx, bson.M{"_id": sectionid}).Decode(&section)
	if err != nil {
		return err
	}

	// Check if the subject already exists in the Subjects slice
	for _, subject := range section.Subjects {
		if subject == subjectID {
			return errors.New("subject already exists in the section")
		}
	}
	
	_, err = SectionsData.UpdateOne(ctx, bson.M{"_id": sectionid}, bson.M{"$push": bson.M{"subjects": subjectID}})
	if err != nil {
		return err
	}
	
	return nil
}

func NewAssessment(assessment models.AssessmentType, sectionID string, subjectID string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    
    sectionid, err := primitive.ObjectIDFromHex(sectionID)
    if err != nil {
        return err
    }

 
    var section models.Section
    err = SectionsData.FindOne(ctx, bson.M{"_id": sectionid}).Decode(&section)
    if err != nil {
        return err
    }

    
    subjectFound := false

   
    for i, subject := range section.Assessments {
        if subject.SubjectID == subjectID {
            section.Assessments[i].Assessment = append(section.Assessments[i].Assessment, assessment)
            subjectFound = true
            break
        }
    }

    if !subjectFound {
        var newSubjectAssessment models.SubjectAssessment
        newSubjectAssessment.SubjectID = subjectID
        newSubjectAssessment.Assessment = append(newSubjectAssessment.Assessment, assessment)

       
        section.Assessments = append(section.Assessments, newSubjectAssessment)
    }

    _, err = SectionsData.UpdateOne(
        ctx,
        bson.M{"_id": sectionid},
        bson.M{"$set": bson.M{"assessments": section.Assessments}},
    )
    if err != nil {
        return err
    }

    return nil
}

func GetStudentsGradeReport(sectionID string,subjectID string)(map[string][]models.AssessmentType,error){
    fmt.Println("on db")
	students ,err := GetSectionStudents(sectionID)
	if err != nil {
		return nil,err
	}
	fmt.Println("get students")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gradeReport := make(map[string][]models.AssessmentType)
	for _,student := range students {
		var result models.GradeReport
		err := GradeReportData.FindOne(ctx,bson.M{"student_id":student.StudentID.Hex()}).Decode(&result)
		if err != nil {
			return nil,err
		}
		gradeReport[student.StudentID.Hex()] = make([]models.AssessmentType, 0)
		for _,subject := range result.Result {
			if subject.SubjectID == subjectID {
				gradeReport[student.StudentID.Hex()] = subject.Assessment
				break
		}

	}
}
return gradeReport,nil

}
func UpdateSectionResult(sectionID string, studentID string, assessmentID string, assessmentValue float64, subjectID string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Convert IDs to ObjectID
    studentIDObj, err := primitive.ObjectIDFromHex(studentID)
    if err != nil {
        return err
    }
    assessmentIDObj, err := primitive.ObjectIDFromHex(assessmentID)
    if err != nil {
        return err
    }

    // Find the Grade Report for the student
    var result models.GradeReport
    err = GradeReportData.FindOne(ctx, bson.M{"student_id": studentID}).Decode(&result)
    if err != nil {
        // If no report exists for the student, we need to create one
        if err == mongo.ErrNoDocuments {
            // Create new subject assessment
            newAssessment := models.AssessmentType{
                AssessmentTypeID: assessmentIDObj,
                AssessmentValue:  assessmentValue,
            }

            // Create a new SubjectAssessment
            newSubject := models.SubjectAssessment{
                SubjectID: subjectID,
                Assessment: []models.AssessmentType{newAssessment},
            }

            // Create new GradeReport
            newGradeReport := models.GradeReport{
                StudentID: studentID,
                Result:    []models.SubjectAssessment{newSubject},
            }

            // Insert the new GradeReport into the database
            _, err := GradeReportData.InsertOne(ctx, newGradeReport)
            if err != nil {
                return err
            }
            return nil // Successfully created new report
        }
        return err
    }

    // Flag to check if the assessment already exists
    assessmentFound := false

    // Loop through the subjects to find the one we need to update
    for i, subject := range result.Result {
        if subject.SubjectID == subjectID {
            for j, assessment := range subject.Assessment {
                if assessment.AssessmentTypeID == assessmentIDObj {
                    // Update the assessment value if it already exists
                    result.Result[i].Assessment[j].AssessmentValue = assessmentValue
                    assessmentFound = true
                    break
                }
            }
            // If assessment not found, create a new one
            if !assessmentFound {
                newAssessment := models.AssessmentType{
                    AssessmentTypeID: assessmentIDObj,
                    AssessmentValue:  assessmentValue,
                }
                result.Result[i].Assessment = append(result.Result[i].Assessment, newAssessment)
            }

            // Update the GradeReport in the database
            _, err = GradeReportData.UpdateOne(ctx,
                bson.M{"student_id": studentIDObj.Hex()},
                bson.M{"$set": bson.M{"subjects_result": result.Result}})
            if err != nil {
                return err
            }
            return nil // Successfully updated or added
        }
    }

    // If we reach here, the subject was not found, so we create a new subject entry
	fmt.Println("subject not found")
    newAssessment := models.AssessmentType{
        AssessmentTypeID: assessmentIDObj,
        AssessmentValue:  assessmentValue,
    }

    newSubject := models.SubjectAssessment{
        SubjectID: subjectID,
        Assessment: []models.AssessmentType{newAssessment},
    }

    // Add the new subject to the existing GradeReport
    result.Result = append(result.Result, newSubject)

    // Update the GradeReport in the database
    _, err = GradeReportData.UpdateOne(ctx,
        bson.M{"student_id": studentIDObj.Hex()},
        bson.M{"$set": bson.M{"subjects_result": result.Result}})
    if err != nil {
		fmt.Println("error",err.Error())
        return err
    }

    return nil // Successfully added new subject
}



func GetAssessmentByID(sectionID string, assessmentID string) (*models.AssessmentType, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Convert sectionID to ObjectID
    sectionid, err := primitive.ObjectIDFromHex(sectionID)
    if err != nil {
        return nil, err
    }

    // Find the section document
    var section models.Section
    err = SectionsData.FindOne(ctx, bson.M{"_id": sectionid}).Decode(&section)
    if err != nil {
        return nil, err
    }

    // Loop through the assessments to find the specific assessment
    for _, subjectAssessment := range section.Assessments {
        for _, assessment := range subjectAssessment.Assessment {
            // Assuming assessment has an ID field (update this to match your AssessmentType struct)
            if assessment.AssessmentTypeID.Hex() == assessmentID { // Change ID to the correct field name
                return &assessment, nil
            }
        }
    }

    // If the assessment is not found, return an error
    return nil, errors.New("assessment not found")
}

package database

import (
	"context"
	"errors"
	"schoolbackend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
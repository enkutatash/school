package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Club struct {
	ClubID primitive.ObjectID `json:"_id" bson:"_id"`
	ClubName string `json:"club_name" bson:"club_name"`
	ClubDescription string `json:"club_description" bson:"club_description"`
	ClubTeacherID string `json:"teacher" bson:"teacher"`
	ClubStudentRepID string `json:"student_rep" bson:"student_rep"`
	Location string `json:"location" bson:"location"`
	Members []string `json:"members" bson:"members"`
	
}
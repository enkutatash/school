package database

import (
	"context"
	"fmt"

	"schoolbackend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckClub(clubName string) bool{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := Clubs.CountDocuments(ctx, bson.M{"club_name": clubName})
	if err != nil {
		return false
	}
	return count > 0
}

func FoundClub(club models.Club) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	club.Members = []string{}
	club.ApplicationRequests = []string{}
	_, err := Clubs.InsertOne(ctx, club)
	if err != nil {
		return err
	}
	return nil
}

func GetClubs()([]models.Club,error){
	var clubs []models.Club
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	projection := bson.M{
		"club_name":        1, 
		"club_description": 1, 
		"location":         1, 
	}
	cursor, err := Clubs.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil,err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var club models.Club
		cursor.Decode(&club)
		clubs = append(clubs, club)
	}
	return clubs,nil
}

func GetClubByID(clubID string)(*models.Club,error){
	var club models.Club
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clubid, err := primitive.ObjectIDFromHex(clubID)
    if err != nil {
        return nil , err
    }
	projection := bson.M{
		"club_name":        1,
		"club_description": 1,
		"location":         1,
		"members":          1,
		"student_rep":1,
		"teacher":1,
	}
	err = Clubs.FindOne(ctx, bson.M{"_id": clubid}, options.FindOne().SetProjection(projection)).Decode(&club)
	if err != nil {
		return nil,err
	}
	return &club,nil
}

func CheckStudent(studentID string) bool{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(studentID)
    if err != nil {
        return false
    }
	count, err := RegisteredStudent.CountDocuments(ctx, bson.M{"_id": id})
	if err != nil {
		return false
	}
	return count > 0
}

func ApplyClub(studentID string,clubID string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clubid, err := primitive.ObjectIDFromHex(clubID)
    if err != nil {
        return err
    }
	filter := bson.M{"_id": clubid}
	update := bson.M{"$push":bson.M{"application_requests":studentID}}
	_, err = Clubs.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func AcceptRequest(studentID string,clubID string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clubid, err := primitive.ObjectIDFromHex(clubID)
    if err != nil {
        return err
    }
	filter := bson.M{"_id": clubid}
	update := bson.M{"$push":bson.M{"members":studentID}}
	_, err = Clubs.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	update2 := bson.M{"$pull":bson.M{"application_requests":studentID}}
	_, err = Clubs.UpdateOne(ctx, filter, update2)
	if err != nil {
		return err
	}
	return nil
}

func RejectRequest(studentID string,clubID string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clubid, err := primitive.ObjectIDFromHex(clubID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": clubid}
	update2 := bson.M{"$pull":bson.M{"application_requests":studentID}}
	_, err = Clubs.UpdateOne(ctx, filter, update2)
	if err != nil {
		return err
	}
	return nil
}

func GetAllApplicant(clubID string)([]models.Student,error){
	var students []models.Student
   var club models.Club
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()
   clubid, err := primitive.ObjectIDFromHex(clubID)
   if err != nil {
	   return nil,err
   }
   err = Clubs.FindOne(ctx, bson.M{"_id": clubid}).Decode(&club)
   if err != nil {
	   fmt.Println("find club")
	   return nil,err
   }
   for _,studentID := range club.ApplicationRequests{
	   var student models.Student
	   studID , err := primitive.ObjectIDFromHex(studentID)
	   if err != nil {
		   return nil,err
	   }
	   err = RegisteredStudent.FindOne(ctx, bson.M{"_id": studID}).Decode(&student)
	   if err != nil {
		   return nil,err
	   }
	   students = append(students,student)
   }
   return students,nil

}


func AssignLead(clubID string , studentID string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clubid, err := primitive.ObjectIDFromHex(clubID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": clubid}
	update := bson.M{"$set":bson.M{"student_rep":studentID}}
	_, err = Clubs.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
package models

type Notification struct {
	NotificationID string `json:"_id" bson:"_id"`
	PersonID       string `json:"person_id" bson:"person_id"`
	Message 	  string `json:"message" bson:"message"`
}
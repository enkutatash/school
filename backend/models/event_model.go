package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	EventID          primitive.ObjectID `json:"_id" bson:"_id"`
	EventName        string             `json:"event_name" bson:"event_name"`
	EventDescription string             `json:"event_description" bson:"event_description"`
	EventDate        time.Time          `json:"event_date" bson:"event_date"`
	EventLocation    string             `json:"event_location" bson:"event_location"`
	OrganizerID      string             `json:"organizer" bson:"organizer"`
}
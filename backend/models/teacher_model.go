package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	TeacherID    primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName    *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=30"`
	LastName     *string            `json:"last_name" bson:"last_name"  validate:"required,min=2,max=30"`
	Email        *string            `json:"email" bson:"email"  validate:"required,email"`
	Password     *string            `json:"password" bson:"password"  validate:"required,min=6"`
	Phone        *string            `json:"phone" bson:"phone"  validate:"required"`
	Token        *string            `json:"token" bson:"token"`
	RefreshToken *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	SectionID    *string            `json:"section" bson:"section"`
	ClubID       *string            `json:"club" bson:"club"`
	SubjectID    *string            `json:"subject" bson:"subject"`
	Role         string             `json:"role" bson:"role"`
}

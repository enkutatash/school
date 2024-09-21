package token

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"schoolbackend/database"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	Role      string
	jwt.StandardClaims
}

var userData *mongo.Collection = database.GetData(database.Client, "user")
var Secret_key = os.Getenv("SECRET_KEY")

func GenerateToken(email string, firstName string, uid string, lastname string,Role string) (signedToken string, refreshToken string, err error) {
	// Claims for the main token
	claim := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastname,
		Uid:       uid,
		Role:      Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	refreshClaim := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 168).Unix(), // Refresh token expires in 7 days
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(Secret_key))
	if err != nil {
		return "", "", fmt.Errorf("error in generating token: %v", err)
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim).SignedString([]byte(Secret_key))
	if err != nil {
		return "", "", fmt.Errorf("error in generating refresh token: %v", err)
	}

	return token, refreshToken, nil
}

func ValidateToken(signedtoken string) (claim *SignedDetails, msg string) {

	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret_key), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claim, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Invalid token"

	}
	if claim.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token has expired"

	}

	return claim, msg
}

func UpdateAllToken(signedtoken string, signedrefreshtoken string, uid string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedtoken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedrefreshtoken})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updated_at})

	upsert := true
	filter := bson.M{"user_id": uid}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userData.UpdateOne(ctx, filter, bson.D{{
		Key: "$set", Value: updateObj,
	}}, &opt)

	defer cancel()
	if err != nil {
		log.Panic(err)
	}

}

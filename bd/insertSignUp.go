package bd

import (
	"context"
	"time"

	"github.com/josnunezg/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertSignUp es the function to create an user in the BD
func InsertSignUp(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor-example")
	col := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

package bd

import (
	"context"
	"time"

	"github.com/josnunezg/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CgeckingExistingUser check if user exist
func CheckingExistingUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor-example")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var user models.User

	err := col.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()

	if err != nil {
		return user, false, ID
	}

	return user, true, ID
}

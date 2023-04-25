package db

import (
	"context"
	"time"

	"github.com/agomezguru/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* Verify if user was created before
 * Return object User if exist (models.User)
 * true if exist (bool)
 * User ID if exist (string)
 */
func UserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	db := MongoCN.Database("twitter")
	collection := db.Collection("users")

	condition := bson.M{"email":email} // Format native of MogoDB

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}  
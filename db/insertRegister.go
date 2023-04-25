package db

import (
	"context"
	"time"

	"github.com/agomezguru/twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Insert register inside MongoDB
 * Return ID of register inserted in json (string)
 * Insert successful (bool)
 * Error code
 */
func InsertRegister(u models.User) (string, bool, error) {
	// Here wait for response no more than 15 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	
	// Cancel timeout insde context
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("users")
	
	u.Password, _ = EncryptData(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	} 

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

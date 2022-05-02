package repositories

import (
	"context"
	"filterisasi/models"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertFiltered(ctx context.Context, database *mongo.Database, ppdbOptions []models.PpdbOption) {

	newValue := make([]interface{}, len(ppdbOptions))
	for _, v := range ppdbOptions {
		newValue = append(newValue, v)
	}

	/*
		objectId1, err := primitive.ObjectIDFromHex("608f7e3819a57c0012556c41")
		objectId2, err := primitive.ObjectIDFromHex("608f7e3819a57c0012556c42")
		objectId3, err := primitive.ObjectIDFromHex("608f7e3819a57c0012556c43")
		if err != nil {
			log.Println("Invalid id")
		}*/

	persons := []interface{}{}

	for i := 0; i < len(ppdbOptions); i++ {
		for _, v := range ppdbOptions[i].PpdbRegistration {
			persons = append(persons, v)
		}
	}

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	/*users := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}*/
	// insert the bson object slice using InsertMany()
	results, err := database.Collection("ppdb_filtereds").InsertMany(ctx, persons)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)
}

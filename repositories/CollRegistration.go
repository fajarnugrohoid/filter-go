package repositories

import (
	"context"
	"filterisasi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetRegistrations(ctx context.Context, database *mongo.Database, firstChoice primitive.ObjectID) []models.PpdbRegistration {

	//var optId = [1]primitive.ObjectID{firstChoice}
	criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "smk"}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"score", 1}})

	csr, err := database.Collection("ppdb_registrations").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.PpdbRegistration, 0)
	for csr.Next(ctx) {
		var row models.PpdbRegistration
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return result
}

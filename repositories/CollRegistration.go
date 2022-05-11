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
	//criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "sma", "status": "fit"}
	criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "sma"}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"distance1", 1}})

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

		tmp := models.PpdbRegistration{
			Id:                 row.Id,
			Name:               row.Name,
			OptionType:         row.OptionType,
			FirstChoiceOption:  row.FirstChoiceOption,
			SecondChoiceOption: row.SecondChoiceOption,
			ThirdChoiceOption:  row.ThirdChoiceOption,
			Score:              row.Score,
			Distance:           row.Distance1,
			Distance1:          row.Distance1,
			Distance2:          row.Distance2,
			Distance3:          row.Distance3,
			BirthDate:          row.BirthDate,
			AcceptedStatus:     0,
			AcceptedIndex:      0, //perlu di update idx berapa untuk firstchoice
			AcceptedChoiceId:   row.FirstChoiceOption,
		}

		result = append(result, tmp)
	}
	return result
}

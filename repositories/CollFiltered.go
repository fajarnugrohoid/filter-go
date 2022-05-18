package repositories

import (
	"context"
	"filterisasi/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InsertFiltered(ctx context.Context, database *mongo.Database, ppdbOptions []*models.PpdbOption, option_type string) {

	DeleteFilteredByOptionType(ctx, database, option_type)

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
		fmt.Println("optFiltered:", ppdbOptions[i].Name, " registrations:", len(ppdbOptions[i].PpdbRegistration))
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

func DeleteFilteredByOptionType(ctx context.Context, database *mongo.Database, option_type string) {
	//persons := []interface{}{}

	f := bson.M{"option_type": bson.M{"$eq": option_type}}

	/*
		for i := 0; i < len(ppdbOptions); i++ {
			for _, v := range ppdbOptions[i].PpdbRegistration {
				persons = append(persons, v.AcceptedStatus)
			}
		}*/
	_, err := database.Collection("ppdb_filtereds").DeleteMany(ctx, f)
	if err != nil {
		panic(err)
	}
}

func GetFiltereds(ctx context.Context, database *mongo.Database, optionType string) []models.PpdbFiltered {

	//var optId = [1]primitive.ObjectID{firstChoice}
	//criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "sma", "status": "fit"}
	criteria := bson.M{"option_type": optionType, "registration_level": "sma"}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"distance1", 1}})

	csr, err := database.Collection("ppdb_filtered").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.PpdbFiltered, 0)
	for csr.Next(ctx) {
		var row models.PpdbFiltered

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := models.PpdbFiltered{
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
			AcceptedStatus:     row.AcceptedStatus,
			AcceptedIndex:      row.AcceptedIndex,
		}

		result = append(result, tmp)
	}
	return result
}

func GetFilteredsByOpt(ctx context.Context, database *mongo.Database, optionType string, optId primitive.ObjectID) []models.PpdbFiltered {

	//var optId = [1]primitive.ObjectID{firstChoice}
	//criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "sma", "status": "fit"}
	criteria := bson.M{"option_type": optionType, "accepted_choice_id": optId}
	findOptions := options.Find()

	csr, err := database.Collection("ppdb_filtereds").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.PpdbFiltered, 0)
	for csr.Next(ctx) {
		var row models.PpdbFiltered

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := models.PpdbFiltered{
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
			AcceptedStatus:     row.AcceptedStatus,
			AcceptedIndex:      row.AcceptedIndex,
		}

		result = append(result, tmp)
	}
	return result
}

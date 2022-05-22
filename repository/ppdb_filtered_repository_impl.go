package repository

import (
	"context"
	"filterisasi/models/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type PpdbFilteredRepositoryImpl struct {
}

func NewPpdbFilteredRepository() PpdbFilteredRepository {
	return &PpdbFilteredRepositoryImpl{}
}

func (filtered PpdbFilteredRepositoryImpl) Save(ctx context.Context, database *mongo.Database, ppdbOptions []*domain.PpdbOption, optionType string) (*mongo.InsertManyResult, error) {
	//TODO implement me
	filtered.DeleteByOptionType(ctx, database, optionType)

	newValue := make([]interface{}, len(ppdbOptions))
	for _, v := range ppdbOptions {
		newValue = append(newValue, v)
	}

	persons := []interface{}{}

	for i := 0; i < len(ppdbOptions); i++ {
		fmt.Println("optFiltered:", ppdbOptions[i].Name, " registrations:", len(ppdbOptions[i].PpdbRegistration))
		for _, v := range ppdbOptions[i].PpdbRegistration {
			persons = append(persons, v)
		}
	}

	// insert the bson object slice using InsertMany()
	results, err := database.Collection("ppdb_filtereds").InsertMany(ctx, persons)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	//fmt.Println(results.InsertedIDs)
	return results, err
}

func (p PpdbFilteredRepositoryImpl) DeleteByOptionType(ctx context.Context, database *mongo.Database, option_type string) {
	//TODO implement me
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

func (filtered PpdbFilteredRepositoryImpl) GetByOpt(ctx context.Context, database *mongo.Database, optionType string, optId primitive.ObjectID) []domain.PpdbFiltered {
	//TODO implement me
	//var optId = [1]primitive.ObjectID{firstChoice}
	//criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "sma", "status": "fit"}
	criteria := bson.M{"option_type": optionType, "accepted_choice_id": optId}
	findOptions := options.Find()

	csr, err := database.Collection("ppdb_filtereds").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]domain.PpdbFiltered, 0)
	for csr.Next(ctx) {
		var row domain.PpdbFiltered

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := domain.PpdbFiltered{
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

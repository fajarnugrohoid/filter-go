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

func InsertStatistic(ctx context.Context, database *mongo.Database, ppdbStatistics []domain.PpdbStatistic, option_type string) {

	DeleteStatisticByOptionType(ctx, database, option_type)

	persons := []interface{}{}

	for i := 0; i < len(ppdbStatistics); i++ {
		persons = append(persons, ppdbStatistics[i])
	}

	// insert the bson object slice using InsertMany()
	results, err := database.Collection("ppdb_statistic").InsertMany(ctx, persons)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)
}

func DeleteStatisticByOptionType(ctx context.Context, database *mongo.Database, option_type string) {

	f := bson.M{"option_type": bson.M{"$eq": option_type}}

	_, err := database.Collection("ppdb_statistic").DeleteMany(ctx, f)
	if err != nil {
		panic(err)
	}
}

func GetAllStatistic(ctx context.Context, database *mongo.Database, optionType string) []domain.PpdbStatistic {

	criteria := bson.M{"option_type": optionType}
	findOptions := options.Find()

	csr, err := database.Collection("ppdb_statistic").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]domain.PpdbStatistic, 0)
	for csr.Next(ctx) {
		var row domain.PpdbStatistic

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := domain.PpdbStatistic{
			Id:         row.Id,
			Name:       row.Name,
			OptionType: row.OptionType,
			Pg:         row.Pg,
		}

		result = append(result, tmp)
	}
	return result
}

func GetStatisticById(ctx context.Context, database *mongo.Database, optionType string, id primitive.ObjectID) []domain.PpdbStatistic {

	criteria := bson.M{"option_type": optionType, "_id": id}
	findOptions := options.Find()

	csr, err := database.Collection("ppdb_statistic").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]domain.PpdbStatistic, 0)
	for csr.Next(ctx) {
		var row domain.PpdbStatistic

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := domain.PpdbStatistic{
			Id:         row.Id,
			Name:       row.Name,
			OptionType: row.OptionType,
			Pg:         row.Pg,
		}

		result = append(result, tmp)
	}
	return result
}

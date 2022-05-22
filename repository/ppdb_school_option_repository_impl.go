package repository

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type PpdbSchoolOptionRepositoryImpl struct {
}

func NewPpdbSchoolOptionRepositoy() PpdbSchoolOptionRepository {
	return &PpdbSchoolOptionRepositoryImpl{}
}

func (p PpdbSchoolOptionRepositoryImpl) GetSchoolOptionByLevel(ctx context.Context, database *mongo.Database, level string) []domain.PpdbOption {
	//TODO implement me
	//var obj1, _ = primitive.ObjectIDFromHex("608f879478c5383cc367ce62")

	registrationsCollection := database.Collection("ppdb_options")
	var optionsType = [7]string{"abk", "kondisi-tertentu", "ketm", "perpindahan", "prestasi-rapor", "prestasi", "zonasi"}

	objectId1, err := primitive.ObjectIDFromHex("608f85bf78c5383cc367c2f8")
	objectId2, err := primitive.ObjectIDFromHex("608f85be78c5383cc367c2f1")
	objectId6, err := primitive.ObjectIDFromHex("608f85bc78c5383cc367c2dd")
	if err != nil {
		log.Println("Invalid id")
	}

	var schoolIds = [3]primitive.ObjectID{objectId1, objectId2, objectId6}
	matchStage := bson.D{{"$match", bson.M{
		"type": bson.M{"$in": optionsType},
		"$and": []bson.M{bson.M{
			"school": bson.M{"$in": schoolIds},
		},
		},
	}}}
	//groupStage := bson.M{{"$group", bson.M{{"_id", "$podcast"}, {"total", bson.M{{"$sum", "$duration"}}}}}}

	pipeline := []bson.M{
		bson.M{"$match": bson.M{
			"$expr": bson.M{
				"$and": []bson.M{
					{"$eq": []string{"$_id", "$$school"}},
					{"$eq": []string{"$level", level}},
				},
			},
		}},
	}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "ppdb_schools"},
		{"let", bson.D{{"school", "$school"}}},
		{"pipeline", pipeline},
		{"as", "ppdb_schools"}}}}
	unwindStage := bson.D{{"$unwind", "$ppdb_schools"}}
	sortByName := bson.D{{"$sort", bson.D{{"name", 1}}}}
	sortByType := bson.D{{"$sort", bson.D{{"type", 1}}}}
	//allowDisk := bson.D{{"allow", true}}
	//fields := bson.D{{"$project", bson.D{{"name", 1}}}}

	showInfoCursor, err := registrationsCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, unwindStage, sortByName, sortByType,
	}, options.Aggregate().SetAllowDiskUse(true))

	if err != nil {
		panic(err)
	}

	//var showsWithInfo []bson.M
	var showsWithInfo []domain.PpdbOption

	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	defer showInfoCursor.Close(ctx)
	return showsWithInfo
}

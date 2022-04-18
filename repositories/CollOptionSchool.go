package repositories

import (
	"context"
	"filterisasi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

/*
	db.ppdb_options.aggregate([
		{ "$match" :
			{
				"type" :
					{ "$in" : ["abk", "ketm"]},
					"$and" : [{
						"school" : {
							"$in" : [ ObjectId("608f879478c5383cc367ce62")]
							}
					}]
				}
		},
				{
					"$lookup" :
					{
						"from" : "ppdb_schools",
						"let" : { "school" : "$school"},
						"pipeline" : [{ "$match" : { "$expr" : { "$eq" : ["$_id", "$$school"]}}}],
						"as" : "ppdb_schools"
					}
				},
				{ "$unwind" : "$ppdb_schools"},
				{ "$sort" : { "name" : 1}},
				{ "$sort" : { "type" : 1}},
				{ "$project" : { "_class" : 0}}
	]);
*/

func GetSchoolAndOption(ctx context.Context, database *mongo.Database) []models.PpdbOption {

	//var obj1, _ = primitive.ObjectIDFromHex("608f879478c5383cc367ce62")

	registrationsCollection := database.Collection("ppdb_options")

	objectId, err := primitive.ObjectIDFromHex("608f866978c5383cc367c75f")
	if err != nil {
		log.Println("Invalid id")
	}

	var optionsType = [2]string{"rapor"}
	var schoolIds = [1]primitive.ObjectID{objectId}
	matchStage := bson.D{{"$match", bson.M{
		"type": bson.M{"$in": optionsType},
		"$and": []bson.M{bson.M{
			"school": bson.M{"$in": schoolIds},
		},
		},
	}}}
	//groupStage := bson.M{{"$group", bson.M{{"_id", "$podcast"}, {"total", bson.M{{"$sum", "$duration"}}}}}}

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"$expr": bson.M{"$eq": []string{"$_id", "$$school"}}}},
	}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "ppdb_schools"},
		{"let", bson.D{{"school", "$school"}}},
		{"pipeline", pipeline},
		{"as", "ppdb_schools"}}}}
	unwindStage := bson.D{{"$unwind", "$ppdb_schools"}}
	sortByName := bson.D{{"$sort", bson.D{{"name", 1}}}}
	sortByType := bson.D{{"$sort", bson.D{{"type", 1}}}}
	//fields := bson.D{{"$project", bson.D{{"name", 1}}}}

	showInfoCursor, err := registrationsCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, unwindStage, sortByName, sortByType,
	})
	/*pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
	    [
			{ "$match" :
				{
					"type" :
						{ "$in" : ["ketm"]},
						"$and" : [{
							"school" : {
								"$in" : schoolIds
								}
						}]
					}
			},
					{
						"$lookup" :
						{
							"from" : "ppdb_schools",
							"let" : { "school" : "$school"},
							"pipeline" : [{ "$match" : { "$expr" : { "$eq" : ["$_id", "$$school"]}}}],
							"as" : "ppdb_schools"
						}
					},
					{ "$unwind" : "$ppdb_schools"},
					{ "$sort" : { "name" : 1}},
					{ "$sort" : { "type" : 1}},
					{ "$project" : { "_class" : 0}}
		]
		`)), true, &pipeline)
	showInfoCursor, err := registrationsCollection.Aggregate(ctx,
		pipeline,
	)
	*/

	if err != nil {
		panic(err)
	}
	//var showsWithInfo []bson.M
	var showsWithInfo []models.PpdbOption
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	defer showInfoCursor.Close(ctx)

	return showsWithInfo
}

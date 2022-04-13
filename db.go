package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

/*
// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
} */

// insertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

// insertMany is a user defined method, used to insert
// documents into collection returns result of
// InsertMany and error if any.
func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	// select database and collection ith Client.Database
	// method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertMany accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}

type Marks struct {
	rollNo   int // movie title
	maths    int // movie release year
	science  int // movie rating
	computer int // movie rating
}

func displayMarks(marks []*Marks) {
	fmt.Println("Marks : ")
	fmt.Println(len(marks))
	for _, m := range marks {
		fmt.Printf("\t- %s (%d) R:%.1f\n", m.rollNo)
	}
	fmt.Println()
}

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

type ppdbOption struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

func getSchoolAndOption(ctx context.Context, registrationsCollection *mongo.Collection) []ppdbOption {

	//var obj1, _ = primitive.ObjectIDFromHex("608f879478c5383cc367ce62")
	objectId, err := primitive.ObjectIDFromHex("608f879478c5383cc367ce62")
	if err != nil {
		log.Println("Invalid id")
	}

	var optionsType = [2]string{"abk", "ketm"}
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
	fields := bson.D{{"$project", bson.D{{"name", 1}}}}

	showInfoCursor, err := registrationsCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, unwindStage, sortByName, sortByType, fields,
	})
	if err != nil {
		panic(err)
	}
	//var showsWithInfo []bson.M
	var showsWithInfo []ppdbOption
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	return showsWithInfo
}

func main() {

	/*
		// get Client, Context, CancelFunc and err from connect method.
		client, ctx, cancel, err := connect("mongodb://localhost:27017")
		if err != nil {
			panic(err)
		}

		// Release resource when main function is returned.
		defer close(client, ctx, cancel) */

	// Create  a object of type interface to  store
	// the bson values, that  we are inserting into database.
	/*var document interface{}


	document = bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and aresult of
	// insertina single document into the collection.
	insertOneResult, err := insertOne(client, ctx, "gfg",
		"marks", document)

	// handle the error
	if err != nil {
		panic(err)
	}

	// print the insertion id of the document,
	// if it is inserted.
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	// Now will be inserting multiple documents into
	// the collection. create  a object of type slice
	// of interface to store multiple  documents
	var documents []interface{}

	// Storing into interface list.
	documents = []interface{}{
		bson.D{
			{"rollNo", 153},
			{"maths", 65},
			{"science", 59},
			{"computer", 55},
		},
		bson.D{
			{"rollNo", 162},
			{"maths", 86},
			{"science", 80},
			{"computer", 69},
		},
	}

	// insertMany insert a list of documents into
	// the collection. insertMany accepts client,
	// context, database name collection name
	// and slice of interface. returns error
	// if any and result of multi document insertion.
	insertManyResult, err := insertMany(client, ctx, "gfg",
		"marks", documents)

	// handle the error
	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertMany")

	// print the insertion ids of the multiple
	// documents, if they are inserted.
	for id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	} */

	// create a filter an option of type interface,
	// that stores bjson objects.
	/*
		var filter, option interface{}

		// filter  gets all document,
		// with maths field greater that 70
		filter = bson.D{
			{"maths", bson.D{{"$gt", 70}}},
		}

		//  option remove id field from all documents
		option = bson.D{{"_id", 0}}

		// call the query method with client, context,
		// database name, collection  name, filter and option
		// This method returns momngo.cursor and error if any.
		cursor, err := query(client, ctx, "gfg",
			"marks", filter, option)
		// handle the errors.
		if err != nil {
			panic(err)
		}

		var results []bson.D

		// to get bson object  from cursor,
		// returns error if any.
		if err := cursor.All(ctx, &results); err != nil {

			// handle the error
			panic(err)
		}

		//var m bson.M
		var s []*Marks

		// convert m to s
		//bsonBytes, _ := bson.Marshal(results)
		//bson.Unmarshal(bsonBytes, &s)

		data, err := bson.Marshal(results)
		if err != nil {
			panic(err)
		}
		err = bson.Unmarshal(data, &s)
		if err != nil {
			panic(err)
		}

		// printing the result of query.
		fmt.Println("Query Result")
		for _, doc := range results {
			fmt.Println(doc)
		}

		displayMarks(s) */

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")
	registrationsCollection := database.Collection("ppdb_options")

	//var schoolOption []bson.M
	var schoolOption []ppdbOption
	schoolOption = getSchoolAndOption(ctx, registrationsCollection)
	if err != nil {
		panic(err)
	}
	fmt.Println(schoolOption)

}

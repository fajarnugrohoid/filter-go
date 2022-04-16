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

type ppdbSchool struct {
	Type  string `bson:"type,omitempty"`
	Level string `bson:"level,omitempty"`
	code  int    `bson:"code,omitempty"`
}

type PpdbOption struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name,omitempty"`
	Type             string             `bson:"type,omitempty"`
	Quota            int                `bson:"quota,omitempty"`
	QuotaOld         int                `bson:"quota_old,omitempty"`
	TotalQuota       int                `bson:"total_quota,omitempty"`
	SchoolId         primitive.ObjectID `bson:"school,omitempty"`
	PpdbSchool       ppdbSchool         `bson:"ppdb_schools,omitempty"`
	ppdbRegistration []PpdbRegistration
}

type PpdbRegistration struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name,omitempty"`
	FirstChoiceOption  primitive.ObjectID `bson:"first_choice_option,omitempty"`
	SecondChoiceOption primitive.ObjectID `bson:"second_choice_option,omitempty"`
	ThirdChoiceOption  primitive.ObjectID `bson:"third_choice_option,omitempty"`
	Score              float32            `bson:"score,omitempty"`
}

func (ppdbOption PpdbOption) addItem(options []PpdbOption) []PpdbOption {
	return append(options, ppdbOption)
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

func getSchoolAndOption(ctx context.Context, database *mongo.Database) []PpdbOption {

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
	var showsWithInfo []PpdbOption
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	defer showInfoCursor.Close(ctx)

	return showsWithInfo
}

func find(ctx context.Context, database *mongo.Database, firstChoice primitive.ObjectID) []PpdbRegistration {

	//var optId = [1]primitive.ObjectID{firstChoice}
	criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "smk"}
	csr, err := database.Collection("ppdb_registrations").Find(ctx, criteria)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]PpdbRegistration, 0)
	for csr.Next(ctx) {
		var row PpdbRegistration
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return result
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

	//var schoolOption []bson.M
	var schoolOption []PpdbOption
	schoolOption = getSchoolAndOption(ctx, database)
	if err != nil {
		panic(err)
	}
	//fmt.Println(schoolOption)

	ppdbOptions := make([]PpdbOption, 0)

	for _, opt := range schoolOption {

		fmt.Printf(opt.Id.String())
		fmt.Printf(opt.PpdbSchool.Level)
		fmt.Printf(opt.PpdbSchool.Type)

		var studentRegistrations []PpdbRegistration
		studentRegistrations = find(ctx, database, opt.Id)
		/*for _, std := range studentRegistrations {
			fmt.Println(std.Name)
		}*/
		tmp := PpdbOption{Id: opt.Id, Name: opt.Name, ppdbRegistration: studentRegistrations}
		ppdbOptions = append(ppdbOptions, tmp)
		/*	ppdbOptions[i].Name = opt.Name
			ppdbOptions[i].Quota = opt.Quota
			ppdbOptions[i].ppdbRegistration = studentRegistrations
		*/
	}

	/*objectId, err := primitive.ObjectIDFromHex("60b5e513977fa9bd4ca13853")
	if err != nil {
		log.Println("Invalid id")
	}

		var studentRegistrations []ppdbRegistration
		studentRegistrations = find(ctx, database, objectId)
		for _, std := range studentRegistrations {
			fmt.Println(std.Name)
		} */

	fmt.Println("len:", len(ppdbOptions))
	for _, opt := range ppdbOptions {
		fmt.Println(opt.Id, " - ", opt.Name, " \n ")
		for _, std := range opt.ppdbRegistration {
			fmt.Println(">", std.Name)
		}
	}

}

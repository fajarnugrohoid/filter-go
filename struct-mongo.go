package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DummyStruct struct {
	User string  `bson:"user" json:"user"`
	Foo  FooType `bson:"foo" json:"foo"`
}

type FooType struct {
	BarA int `bson:"barA" json:"barA"`
	BarB int `bson:"bar_b" json:"bar_b"`
}

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
}

// query method returns a cursor and error.
func queryTest(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.

	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}

func fetchExpensiveItems(origin string, minPrice float64) ([]Item, error) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("gfg").C("catalog")
	defer session.Close()

	//create the aggregator pipeline that will fetch just the needed data from MongoDB, and nothing more
	pipe := c.Pipe([]bson.M{
		{"$match": bson.M{
			"brands": bson.M{
				"$elemMatch": bson.M{
					"items.origin": bson.M{"$eq": origin},
					"items.price":  bson.M{"$gte": minPrice},
				},
			},
		}},
		{"$project": bson.M{"_id": 0, "brands": 1}},
		{"$addFields": bson.M{
			"brands": bson.M{
				"$filter": bson.M{
					"input": bson.M{
						"$map": bson.M{
							"input": "$brands",
							"as":    "b",
							"in": bson.M{
								"items": bson.M{
									"$filter": bson.M{
										"input": "$$b.items",
										"as":    "i",
										"cond": bson.M{
											"$and": []interface{}{
												bson.M{"$eq": []interface{}{"$$i.origin", origin}},
												bson.M{"$gte": []interface{}{"$$i.price", minPrice}},
											},
										},
									},
								},
							},
						},
					},
					"as":   "b",
					"cond": bson.M{"$gt": []interface{}{bson.M{"$size": "$$b.items"}, 0}},
				},
			},
		},
		}})

	//execute the aggregation query
	var resp []bson.M
	err = pipe.All(&resp)
	if err != nil {
		return nil, err
	}

	//traverse the bson Map returned by the aggregation and extract the items
	var itemsFound []Item
	for _, catalogMap := range resp {
		brands := catalogMap["brands"].([]interface{})
		for _, b := range brands {
			brandsMap := b.(bson.M)
			items := brandsMap["items"].([]interface{})
			for _, b := range items {
				itemsMap := b.(bson.M)
				data, _ := json.Marshal(itemsMap)
				var item Item
				if err := json.Unmarshal(data, &item); err != nil {
					return nil, err
				}
				itemsFound = append(itemsFound, item)
			}
		}
	}

	return itemsFound, err
}

func Test(results []bson.M, err error) ([]Item, error) {
	//traverse the bson Map returned by the aggregation and extract the items
	var itemsFound []Item
	for _, catalogMap := range results {
		brands := catalogMap["brands"].([]interface{})
		for _, b := range brands {
			brandsMap := b.(bson.M)
			items := brandsMap["items"].([]interface{})
			for _, b := range items {
				itemsMap := b.(bson.M)
				data, _ := json.Marshal(itemsMap)
				var item Item
				if err := json.Unmarshal(data, &item); err != nil {
					return nil, err
				}
				itemsFound = append(itemsFound, item)
			}
		}
	}
	return itemsFound, err
}

type Item struct {
	Name   string  `bson:"name" json:"name"`
	Origin string  `bson:"origin" json:"origin"`
	Price  float64 `bson:"price" json:"price"`
}

func main() {
	/*test := DummyStruct{}
	test.User = "test"
	test.Foo.BarA = 123
	test.Foo.BarB = 321
	b, err := json.Marshal(test)
	if err != nil {
		fmt.Println("error marshaling test struct", err)
		return
	}
	fmt.Println("test data\n", string(b))

	items, err := fetchExpensiveItems("Italy", 200)
	if err != nil {
		fmt.Printf("Failed with error: %v", err)
	}

	fmt.Printf("Items matching criteria: %+v", items) */

	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is returned.
	defer close(client, ctx, cancel)

	// select database and collection.
	col := client.Database("gfg").Collection("catalog")

	pipe := []bson.M{
		{"$group": bson.M{
			"_id":   "",
			"sum":   bson.M{"$sum": "$countData"},
			"count": bson.M{"$sum": 1},
		}},
	}
	cursor, err := col.Aggregate(ctx, pipe)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	if err := cursor.Close(ctx); err != nil {
		panic(err)
	}

	itemx, err := Test(results, err)
	if err != nil {
		fmt.Printf("Failed with error: %v", err)
	}

	fmt.Printf("Items matching criteria: %+v", itemx)

	//fmt.Println(results)

}

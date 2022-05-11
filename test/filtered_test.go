package test

import (
	"context"
	"filterisasi/models"
	"filterisasi/repositories"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestStudentFiltered(t *testing.T) {

	ctx := context.Background()
	//ctx, _ := context.WithTimeout(context.Background(), 200*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	var students []models.PpdbFiltered
	var statistic []models.PpdbStatistic

	statistic = repositories.GetStatistic(ctx, database, "ketm")
	fmt.Println("statistic:", len(statistic))
	for i := 0; i < len(statistic); i++ {
		students = repositories.GetFilteredsByOpt(ctx, database, "ketm", statistic[i].Id)
		fmt.Println("students:", len(students))
		for j := 0; j < len(students); j++ {
			fmt.Println(students[j].Name, " ", students[j].Distance)
			if students[j].Distance > statistic[i].Pg {
				fmt.Println(">>", students[j].Name)
			}
		}
	}
}

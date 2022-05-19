package test

import (
	"context"
	"filterisasi/collection"
	"filterisasi/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestByOptionFiltered(t *testing.T) {

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

	statistic = collection.GetAllStatistic(ctx, database, "ketm")
	fmt.Println("statistic:", len(statistic))
	for i := 0; i < len(statistic); i++ {
		students = collection.GetFilteredsByOpt(ctx, database, "ketm", statistic[i].Id)
		fmt.Println("students:", len(students))
		for j := 0; j < len(students); j++ {
			fmt.Println(students[j].Name, " ", students[j].Distance)
			if students[j].Distance > statistic[i].Pg {
				fmt.Println(">>", students[j].Name)
			}
		}
	}
}

func TestByStudentFiltered(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	var optType string
	optType = "kondisi-tertentu"

	var students []models.PpdbFiltered
	var statistic []models.PpdbStatistic
	var allStatistic []models.PpdbStatistic
	allStatistic = collection.GetAllStatistic(ctx, database, optType)
	for x := 0; x < len(allStatistic); x++ {
		fmt.Println("optName:", allStatistic[x].Name)
		students = collection.GetFilteredsByOpt(ctx, database, optType, allStatistic[x].Id)
		fmt.Println("students:", len(students))
		for j := 0; j < len(students); j++ {
			fmt.Println(students[j].Name, " ", students[j].Distance1, " accStatus:", students[j].AcceptedStatus)
			if students[j].AcceptedStatus != 0 {
				if students[j].AcceptedStatus == 1 {
					statistic = collection.GetStatisticById(ctx, database, optType, students[j].FirstChoiceOption)
					for i := 0; i < len(statistic); i++ {
						fmt.Println(students[j].Distance1, " < ", statistic[i].Pg)
						assert.Greater(t, students[j].Distance1, statistic[i].Pg) //jika jarak siswa lebih besar dari pg statistic, maka memang benar harus terlempar ke pilihan 2,3, buangan
						if students[j].Distance1 < statistic[i].Pg {
							fmt.Println(">>error")
						}
					}
				}
				if students[j].AcceptedStatus == 2 {
					statistic = collection.GetStatisticById(ctx, database, optType, students[j].FirstChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance1 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Greater(t, students[j].Distance1, statistic[i].Pg)

					}
					statistic = collection.GetStatisticById(ctx, database, optType, students[j].SecondChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance2 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Greater(t, students[j].Distance2, statistic[i].Pg)

					}
				}

				if students[j].AcceptedStatus == 3 {
					statistic = collection.GetStatisticById(ctx, database, optType, students[j].FirstChoiceOption)
					for i := 0; i < len(statistic); i++ {
						fmt.Println(students[j].Name, "-", students[j].Distance1, " < ", statistic[i].Pg)
						if students[j].Distance1 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Greater(t, students[j].Distance1, statistic[i].Pg)
					}
					statistic = collection.GetStatisticById(ctx, database, optType, students[j].SecondChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance2 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Greater(t, students[j].Distance2, statistic[i].Pg)
					}
					statistic = collection.GetStatisticById(ctx, database, optType, students[j].ThirdChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance3 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Greater(t, students[j].Distance3, statistic[i].Pg)
					}
				}
			}
		}
	}
}

func TestCountStudentFiltered(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	var optType, level string
	optType = "ketm"
	level = "sma"

	database := client.Database("ppdb21")
	var filtereds []models.PpdbFiltered
	var statistic []models.PpdbStatistic
	var studentRegistrations []models.PpdbRegistration
	var totalFiltered int
	var totalRegistrations int
	statistic = collection.GetAllStatistic(ctx, database, optType)
	for i := 0; i < len(statistic); i++ {
		filtereds = collection.GetFilteredsByOpt(ctx, database, optType, statistic[i].Id)
		studentRegistrations = collection.GetRegistrations(ctx, database, level, statistic[i].Id)
		fmt.Println("students:", len(filtereds), "==", len(studentRegistrations))
		totalFiltered += len(filtereds)
		totalRegistrations += len(studentRegistrations)
	}
	fmt.Println("=========total===============")
	fmt.Println(totalFiltered, "==", totalRegistrations)
	assert.Equal(t, totalFiltered, totalRegistrations)

}

func TestData(t *testing.T) {
	assert.Less(t, 1, 2)
	assert.Less(t, 4, 2)
	assert.Less(t, 2, 5)
}

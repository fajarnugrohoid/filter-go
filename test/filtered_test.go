package test

import (
	"context"
	"filterisasi/models"
	"filterisasi/repositories"
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

	statistic = repositories.GetAllStatistic(ctx, database, "ketm")
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

func TestByStudentFiltered(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	var students []models.PpdbFiltered
	var statistic []models.PpdbStatistic
	var allStatistic []models.PpdbStatistic
	allStatistic = repositories.GetAllStatistic(ctx, database, "ketm")
	for x := 0; x < len(allStatistic); x++ {
		fmt.Println("optName:", allStatistic[x].Name)
		students = repositories.GetFilteredsByOpt(ctx, database, "ketm", allStatistic[x].Id)
		fmt.Println("students:", len(students))
		for j := 0; j < len(students); j++ {
			fmt.Println(students[j].Name, " ", students[j].Distance1, " accStatus:", students[j].AcceptedStatus)
			if students[j].AcceptedStatus != 0 {
				if students[j].AcceptedStatus == 1 {
					statistic = repositories.GetStatisticById(ctx, database, "ketm", students[j].FirstChoiceOption)
					for i := 0; i < len(statistic); i++ {
						fmt.Println(students[j].Distance1, " < ", statistic[i].Pg)
						assert.Less(t, students[j].Distance1, statistic[i].Pg)
						if students[j].Distance1 < statistic[i].Pg {
							fmt.Println(">>error")
						}
					}
				}
				if students[j].AcceptedStatus == 2 {
					statistic = repositories.GetStatisticById(ctx, database, "ketm", students[j].FirstChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance1 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Less(t, students[j].Distance1, statistic[i].Pg)

					}
					statistic = repositories.GetStatisticById(ctx, database, "ketm", students[j].SecondChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance2 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Less(t, students[j].Distance2, statistic[i].Pg)

					}
				}

				if students[j].AcceptedStatus == 3 {
					statistic = repositories.GetStatisticById(ctx, database, "ketm", students[j].FirstChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance1 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Less(t, students[j].Distance1, statistic[i].Pg)
					}
					statistic = repositories.GetStatisticById(ctx, database, "ketm", students[j].SecondChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance2 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Less(t, students[j].Distance2, statistic[i].Pg)
					}
					statistic = repositories.GetStatisticById(ctx, database, "ketm", students[j].ThirdChoiceOption)
					for i := 0; i < len(statistic); i++ {
						if students[j].Distance3 < statistic[i].Pg {
							fmt.Println(">>error")
						}
						assert.Less(t, students[j].Distance3, statistic[i].Pg)
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

	database := client.Database("ppdb21")
	var filtereds []models.PpdbFiltered
	var statistic []models.PpdbStatistic
	var studentRegistrations []models.PpdbRegistration
	var totalFiltered int
	var totalRegistrations int
	statistic = repositories.GetAllStatistic(ctx, database, "ketm")
	for i := 0; i < len(statistic); i++ {
		filtereds = repositories.GetFilteredsByOpt(ctx, database, "ketm", statistic[i].Id)
		studentRegistrations = repositories.GetRegistrations(ctx, database, statistic[i].Id)
		fmt.Println("students:", len(filtereds), "==", len(studentRegistrations))
		totalFiltered += len(filtereds)
		totalRegistrations += len(studentRegistrations)
	}
	fmt.Println("=========total===============")
	fmt.Println(totalFiltered, "==", totalRegistrations)
	assert.Equal(t, totalFiltered, totalRegistrations)

}

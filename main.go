package main

import (
	"context"
	"filterisasi/models"
	"filterisasi/repositories"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()

	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("STACK")
	fmt.Println(val)
	if os.Getenv("LOGGING") == "file" {
		logger.SetOutput(file)
	}
	logger.Info("hello logging")

	ctx := context.Background()
	//ctx, _ := context.WithTimeout(context.Background(), 200*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("URL")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	var schoolOption []models.PpdbOption

	schoolOption = repositories.GetSchoolAndOption(ctx, database)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(schoolOption))

	ppdbOptions := make([]models.PpdbOption, 0)

	for _, opt := range schoolOption {

		fmt.Printf(opt.Id.String())

		var studentRegistrations []models.PpdbRegistration
		studentRegistrations = repositories.GetRegistrations(ctx, database, opt.Id)
		//for _, std := range studentRegistrations {
		//	fmt.Println(std.Name)
		//}
		tmp := models.PpdbOption{
			Id:               opt.Id,
			Name:             opt.Name,
			Quota:            opt.Quota,
			Filtered:         0,
			PpdbRegistration: studentRegistrations,
		}
		ppdbOptions = append(ppdbOptions, tmp)
		///	ppdbOptions[i].Name = opt.Name
		//	ppdbOptions[i].Quota = opt.Quota
		//	ppdbOptions[i].ppdbRegistration = studentRegistrations
	}
	TmpId, err := primitive.ObjectIDFromHex("000000000000000000000000")
	tmp := models.PpdbOption{
		Id:               TmpId,
		Name:             "Temporary",
		Quota:            0,
		Filtered:         1,
		PpdbRegistration: nil,
	}
	ppdbOptions = append(ppdbOptions, tmp)
	repositories.InsertFiltered(ctx, database, ppdbOptions)

	/*objectId, err := primitive.ObjectIDFromHex("60b5e513977fa9bd4ca13853")
	if err != nil {
		log.Println("Invalid id")
	}
	var studentRegistrations []ppdbRegistration
	studentRegistrations = find(ctx, database, objectId)
	for _, std := range studentRegistrations {
		fmt.Println(std.Name)
	} */

	/*
		fmt.Println("len:", len(ppdbOptions))


		ppdbOptions = utility.ProcessFilter(ppdbOptions, false)

		fmt.Println("===========================res==============================")
		for _, opt := range ppdbOptions {
			fmt.Println(opt.Id, " - ", opt.Name, " : q: ", opt.Quota, " len.std:", len(opt.PpdbRegistration), " \n ")
			for i, std := range opt.PpdbRegistration {
				fmt.Println(">ori:", i, ":", std.Name, " - acc:", std.AcceptedStatus, " score: ", std.Score)
			}
		}*/
	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
}

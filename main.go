package main

import (
	"context"
	"filterisasi/controller"
	"filterisasi/helper"
	"filterisasi/logic"
	"filterisasi/models/domain"
	"filterisasi/repository"
	"filterisasi/service"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func NewDatabase(ctx context.Context, url string) (*mongo.Database, context.Context) {

	fmt.Println("ctx:", ctx)
	fmt.Println("url:", url)
	//ctx, _ := context.WithTimeout(context.Background(), 200*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	return database, ctx
}

func main() {

	start := time.Now()

	//logger := logrus.New()

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	/*
		filename := "logs/log_" + arg + "_" + formatted + ".log"
		file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_WRONLY, 0666)
	*/
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	/*
		val := os.Getenv("STACK")
		fmt.Println(val)
		if os.Getenv("LOGGING") == "file" {
			logger.SetOutput(file)
		}
		logger.Info("hello logging") */

	helper.SetLogArg(arg)

	logger := helper.InstanceLogger(helper.GetLogArg())

	ctx := context.Background()
	url := os.Getenv("URL")
	//database, ctx := NewDatabase(ctx, os.Getenv("URL"))

	fmt.Println("ctx:", ctx)
	fmt.Println("url:", url)
	//ctx, _ := context.WithTimeout(context.Background(), 200*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	var schoolOption []domain.PpdbOption
	var optionTypes map[string][]*domain.PpdbOption
	optionTypes = map[string][]*domain.PpdbOption{}

	fmt.Println("database:", database)
	schoolOption = repository.GetSchoolAndOption(ctx, database)
	if err != nil {
		panic(err)
	}

	logger.Info(len(schoolOption))

	ppdbRegistrationRepository := repository.NewPpdbRegistrationRepositoy()
	ppdbRegistrationService := service.NewPpdbRegistrationService(ppdbRegistrationRepository, database)
	initialController := controller.NewInitialController(ppdbRegistrationService)
	optionTypes = initialController.InitData(ctx, optionTypes, schoolOption)

	logger.Info("len abk:", len(optionTypes["abk"]))
	logger.Info("len ketm:", len(optionTypes["ketm"]))
	for i, opt := range optionTypes["ketm"] {
		logger.Info(i, "-", opt.Id, " - ", opt.Name, " - q: ", opt.Quota, " - p:", len(opt.PpdbRegistration))
		/*for i, std := range opt.PpdbRegistration {
			fmt.Println("", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
				" AcceptedIndex: ", std.AcceptedIndex)
		}*/
	}
	fmt.Println("len kondisi-tertentu:", len(optionTypes["kondisi-tertentu"]))
	for i, opt := range optionTypes["kondisi-tertentu"] {
		logger.Info(i, "-", opt.Id, " - ", opt.Name, " - q: ", opt.Quota, " - p:", len(opt.PpdbRegistration))
		/*for i, std := range opt.PpdbRegistration {
			fmt.Println("", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
				" AcceptedIndex: ", std.AcceptedIndex)
		}*/
	}

	optionTypes = logic.DoFilterSenior(optionTypes, logger)
	/*
		fmt.Println("===========================res-end==============================")
		for _, opt := range optionTypes["ketm"] {
			fmt.Println(opt.Id, " - ", opt.Name, " : q: ", opt.Quota, " len.std:", len(opt.PpdbRegistration), "")
			for i, std := range opt.PpdbRegistration {
				fmt.Println(">ori:", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1)
			}
			fmt.Println("\n")
		}
	*/

	for i := 0; i < len(optionTypes["kondisi-tertentu"]); i++ {
		logger.Info(i, "-", optionTypes["kondisi-tertentu"][i].Id, " - ", optionTypes["kondisi-tertentu"][i].Name,
			" : q: ", optionTypes["kondisi-tertentu"][i].Quota,
			" : p: ", len(optionTypes["kondisi-tertentu"][i].PpdbRegistration),
			" - needQuota:", optionTypes["kondisi-tertentu"][i].NeedQuota,
			" - AddQuota:", optionTypes["kondisi-tertentu"][i].AddQuota,
		)
	}

	controller.UpdateFilteredStatistic(ctx, database, optionTypes, logger)

	timeElapsed := time.Since(start)
	logger.Info("The `for` loop took %s", timeElapsed)
}

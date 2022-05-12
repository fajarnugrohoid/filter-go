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

	file, _ := os.OpenFile("application.log", os.O_RDWR|os.O_CREATE|os.O_WRONLY, 0666)

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

	var optionTypes map[string][]*models.PpdbOption
	optionTypes = map[string][]*models.PpdbOption{}

	//ppdbOptions := make([]models.PpdbOption, 0)

	for _, opt := range schoolOption {

		fmt.Printf(opt.Id.String())

		var studentRegistrations []models.PpdbRegistration
		//var studentHistory []models.PpdbRegistration
		studentRegistrations = repositories.GetRegistrations(ctx, database, opt.Id)
		//for _, std := range studentRegistrations {
		//	fmt.Println(std.Name)
		//}

		studentHistories := make([]models.PpdbRegistration, len(studentRegistrations), cap(studentRegistrations))
		copy(studentHistories, studentRegistrations)

		tmpOpt := &models.PpdbOption{
			Id:                  opt.Id,
			Name:                opt.Name,
			Quota:               opt.Quota,
			Type:                opt.Type,
			AddQuota:            0,
			Filtered:            0,
			UpdateQuota:         true,
			NeedQuotaFirstOpt:   0,
			PpdbRegistration:    studentRegistrations,
			RegistrationHistory: studentHistories,
			HistoryShifting:     make([]models.PpdbRegistration, 0),
		}
		tmpOptKetm := &models.PpdbOption{
			Id:                  opt.Id,
			Name:                opt.Name,
			Quota:               opt.Quota,
			Type:                opt.Type,
			AddQuota:            0,
			Filtered:            0,
			UpdateQuota:         true,
			NeedQuotaFirstOpt:   0,
			PpdbRegistration:    studentRegistrations,
			RegistrationHistory: studentHistories,
			HistoryShifting:     make([]models.PpdbRegistration, 0),
		}
		tmpKondisiTertentu := &models.PpdbOption{
			Id:                  opt.Id,
			Name:                opt.Name,
			Quota:               opt.Quota,
			Type:                opt.Type,
			AddQuota:            0,
			Filtered:            0,
			UpdateQuota:         true,
			NeedQuotaFirstOpt:   0,
			PpdbRegistration:    studentRegistrations,
			RegistrationHistory: studentHistories,
			HistoryShifting:     make([]models.PpdbRegistration, 0),
		}
		//ppdbOptions = append(ppdbOptions, tmpOpt)
		//optionTypes["abk"] = append(optionTypes["abk"], tmpOpt)
		//optionTypes["kondisi-tertentu"] = append(optionTypes["kondisi-tertentu"], tmpOpt)
		//optionTypes["ketm"] = append(optionTypes["ketm"], tmpOpt)

		switch opt.Type {
		case "abk":
			optionTypes["abk"] = append(optionTypes["abk"], tmpOpt)
			break
		case "kondisi-tertentu":
			optionTypes["kondisi-tertentu"] = append(optionTypes["kondisi-tertentu"], tmpKondisiTertentu)
			break
		case "ketm":
			optionTypes["ketm"] = append(optionTypes["ketm"], tmpOptKetm)
			break
		}

		///	ppdbOptions[i].Name = opt.Name
		//	ppdbOptions[i].Quota = opt.Quota
		//	ppdbOptions[i].ppdbRegistration = studentRegistrations
	}

	TmpIdAbk, err := primitive.ObjectIDFromHex("000000000000000000000001")
	TmpIdKetm, err := primitive.ObjectIDFromHex("000000000000000000000002")
	TmpIdKondisiTertentu, err := primitive.ObjectIDFromHex("000000000000000000000003")

	tmpAbk := &models.PpdbOption{
		Id:                  TmpIdAbk,
		Name:                "TemporaryAbk",
		Type:                "abk",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpKetm := &models.PpdbOption{
		Id:                  TmpIdKetm,
		Name:                "TemporaryKetm",
		Type:                "ketm",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpKondisiTertentu := &models.PpdbOption{
		Id:                  TmpIdKondisiTertentu,
		Name:                "TemporaryKondisiTertentu",
		Type:                "kondisi-tertentu",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	optionTypes["abk"] = append(optionTypes["abk"], tmpAbk)
	optionTypes["ketm"] = append(optionTypes["ketm"], tmpKetm)
	optionTypes["kondisi-tertentu"] = append(optionTypes["kondisi-tertentu"], tmpKondisiTertentu)

	/*objectId, err := primitive.ObjectIDFromHex("60b5e513977fa9bd4ca13853")
	if err != nil {
		log.Println("Invalid id")
	}
	var studentRegistrations []ppdbRegistration
	studentRegistrations = find(ctx, database, objectId)
	for _, std := range studentRegistrations {
		fmt.Println(std.Name)
	} */

	fmt.Println("len abk:", len(optionTypes["abk"]))
	fmt.Println("len ketm:", len(optionTypes["ketm"]))
	for i, opt := range optionTypes["ketm"] {
		fmt.Println(i, "-", opt.Id, " - ", opt.Name, " - q: ", opt.Quota, " - p:", len(opt.PpdbRegistration))
		/*for i, std := range opt.PpdbRegistration {
			fmt.Println("", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
				" AcceptedIndex: ", std.AcceptedIndex)
		}*/
	}
	fmt.Println("len kondisi-tertentu:", len(optionTypes["kondisi-tertentu"]))
	for i, opt := range optionTypes["kondisi-tertentu"] {
		fmt.Println(i, "-", opt.Id, " - ", opt.Name, " - q: ", opt.Quota, " - p:", len(opt.PpdbRegistration))
		/*for i, std := range opt.PpdbRegistration {
			fmt.Println("", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
				" AcceptedIndex: ", std.AcceptedIndex)
		}*/
	}

	optionTypes = models.DoFilter(optionTypes)
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

	ppdbStatistics := make([]models.PpdbStatistic, 0)
	for i := 0; i < len(optionTypes["ketm"]); i++ {
		fmt.Println(i, "-", optionTypes["ketm"][i].Id, " - ", optionTypes["ketm"][i].Name,
			" : q: ", optionTypes["ketm"][i].Quota,
			" : p: ", len(optionTypes["ketm"][i].PpdbRegistration),
			" - needQuota:", optionTypes["ketm"][i].NeedQuotaFirstOpt,
			" - AddQuota:", optionTypes["ketm"][i].AddQuota,
		)
		for i, std := range optionTypes["ketm"][i].PpdbRegistration {
			fmt.Println(">", i, ":", std.Name,
				" - acc:", std.AcceptedStatus,
				" - accId:", std.AcceptedChoiceId,
				" distance: ", std.Distance, " Birth:", std.BirthDate)
		}
		for i, std := range optionTypes["ketm"][i].RegistrationHistory {
			fmt.Println("hist>", i, ":", std.Name, " - acc:", std.AcceptedIndex)
		}
		for i, std := range optionTypes["ketm"][i].HistoryShifting {
			fmt.Println("shift>", i, ":", std.Name, " - acc:", std.AcceptedIndex)
		}

		var pg float64
		if len(optionTypes["ketm"][i].PpdbRegistration) > 0 {
			pg = optionTypes["ketm"][i].PpdbRegistration[len(optionTypes["ketm"][i].PpdbRegistration)-1].Distance
		} else {
			pg = 0
		}
		tmpStatistic := models.PpdbStatistic{
			Id:         optionTypes["ketm"][i].Id,
			Name:       optionTypes["ketm"][i].Name,
			OptionType: optionTypes["ketm"][i].Type,
			Quota:      optionTypes["ketm"][i].Quota,
			SchoolId:   optionTypes["ketm"][i].SchoolId,
			Pg:         pg,
		}
		ppdbStatistics = append(ppdbStatistics, tmpStatistic)
	}
	for i := 0; i < len(optionTypes["kondisi-tertentu"]); i++ {
		fmt.Println(i, "-", optionTypes["kondisi-tertentu"][i].Id, " - ", optionTypes["kondisi-tertentu"][i].Name,
			" : q: ", optionTypes["kondisi-tertentu"][i].Quota,
			" : p: ", len(optionTypes["kondisi-tertentu"][i].PpdbRegistration),
			" - needQuota:", optionTypes["kondisi-tertentu"][i].NeedQuotaFirstOpt,
			" - AddQuota:", optionTypes["kondisi-tertentu"][i].AddQuota,
		)
	}

	repositories.InsertFiltered(ctx, database, optionTypes["ketm"], "ketm")
	repositories.InsertFiltered(ctx, database, optionTypes["kondisi-tertentu"], "kondisi-tertentu")
	repositories.InsertStatistic(ctx, database, ppdbStatistics, "ketm")
	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
}

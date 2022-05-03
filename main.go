package main

import (
	"context"
	"filterisasi/models"
	"filterisasi/repositories"
	"filterisasi/utility"
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

	var optionTypes map[string][]*models.PpdbOption
	optionTypes = map[string][]*models.PpdbOption{}

	//ppdbOptions := make([]models.PpdbOption, 0)

	for _, opt := range schoolOption {

		fmt.Printf(opt.Id.String())

		var studentRegistrations []models.PpdbRegistration
		studentRegistrations = repositories.GetRegistrations(ctx, database, opt.Id)
		//for _, std := range studentRegistrations {
		//	fmt.Println(std.Name)
		//}
		tmpOpt := &models.PpdbOption{
			Id:               opt.Id,
			Name:             opt.Name,
			Quota:            opt.Quota,
			Type:             opt.Type,
			Filtered:         0,
			PpdbRegistration: studentRegistrations,
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
			optionTypes["kondisi-tertentu"] = append(optionTypes["kondisi-tertentu"], tmpOpt)
			break
		case "ketm":
			optionTypes["ketm"] = append(optionTypes["ketm"], tmpOpt)
			break
		}

		///	ppdbOptions[i].Name = opt.Name
		//	ppdbOptions[i].Quota = opt.Quota
		//	ppdbOptions[i].ppdbRegistration = studentRegistrations
	}

	TmpId, err := primitive.ObjectIDFromHex("000000000000000000000000")
	tmp := &models.PpdbOption{
		Id:               TmpId,
		Name:             "Temporary",
		Quota:            0,
		Filtered:         1,
		PpdbRegistration: nil,
	}
	optionTypes["ketm"] = append(optionTypes["ketm"], tmp)

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
	fmt.Println("len kondisi-tertentu:", len(optionTypes["kondisi-tertentu"]))
	fmt.Println("len ketm:", len(optionTypes["ketm"]))
	for i, opt := range optionTypes["ketm"] {
		fmt.Println(i, "-", opt.Id, " - ", opt.Name, " : q: ")
	}
	optionTypes["ketm"] = utility.Filter2OptionsShareQuota(optionTypes)

	//optionTypes["ketm"] = utility.ProcessFilter(optionTypes["ketm"], false)

	fmt.Println("===========================res==============================")
	for _, opt := range optionTypes["ketm"] {
		fmt.Println(opt.Id, " - ", opt.Name, " : q: ", opt.Quota, " len.std:", len(opt.PpdbRegistration), "")
		for i, std := range opt.PpdbRegistration {
			fmt.Println(">ori:", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1)
		}
		fmt.Println("\n")
	}

	//repositories.InsertFiltered(ctx, database, ppdbOptions)
	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
}

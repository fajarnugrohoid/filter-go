package main

import (
	"context"
	"filterisasi/models"
	"filterisasi/repositories"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Movie struct {
	Title string  // movie title
	Year  int     // movie release year
	Rate  float32 // movie rating
}

// byYear sorts all movies by release year
type byYear []*Movie

func (m byYear) Len() int           { return len(m) }
func (m byYear) Less(i, j int) bool { return m[i].Year < m[j].Year }
func (m byYear) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

// byTitle sorts all movies by title
type byTitle []*Movie

func (m byTitle) Len() int           { return len(m) }
func (m byTitle) Less(i, j int) bool { return m[i].Title < m[j].Title }
func (m byTitle) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

// byRate sorts all movies by rate
type byRate []*Movie

func (m byRate) Len() int           { return len(m) }
func (m byRate) Less(i, j int) bool { return m[i].Rate < m[j].Rate }
func (m byRate) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func indexOf(element primitive.ObjectID, data []models.PpdbOption) int {
	for k, v := range data {
		if element == v.Id {
			return k
		}
	}
	return -1 //not found.
}

func main() {
	/*
		start := time.Now()

		movies := []*Movie{
			&Movie{"The 400 Blows", 1959, 8.1},
			&Movie{"La Haine", 1995, 8.1},
			&Movie{"The Godfather", 1972, 9.2},
			&Movie{"The Godfather: Part II", 1974, 9},
			&Movie{"Mafioso", 1962, 7.7}}

		displayMovies("Movies (unsorted)", movies)

		sort.Sort(byYear(movies))
		displayMovies("Movies sorted by year", movies)

		sort.Sort(byTitle(movies))
		displayMovies("Movies sorted by title", movies)

		sort.Sort(sort.Reverse(byRate(movies)))
		displayMovies("Movies sorted by rate", movies)

		timeElapsed := time.Since(start)
		fmt.Printf("The `for` loop took %s", timeElapsed)*/

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("ppdb21")

	//var schoolOption []bson.M
	var schoolOption []models.PpdbOption
	schoolOption = repositories.GetSchoolAndOption(ctx, database)
	if err != nil {
		panic(err)
	}
	//fmt.Println(schoolOption)

	ppdbOptions := make([]models.PpdbOption, 0)

	for _, opt := range schoolOption {

		//fmt.Printf(opt.Id.String())
		//fmt.Printf(opt.PpdbSchool.Level)
		//fmt.Printf(opt.PpdbSchool.Type)

		var studentRegistrations []models.PpdbRegistration
		studentRegistrations = repositories.GetRegistrations(ctx, database, opt.Id)
		/*for _, std := range studentRegistrations {
			fmt.Println(std.Name)
		}*/
		tmp := models.PpdbOption{
			Id:               opt.Id,
			Name:             opt.Name,
			Quota:            opt.Quota,
			PpdbRegistration: studentRegistrations,
		}
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
	for i := 0; i < len(ppdbOptions); i++ {
		fmt.Println(ppdbOptions[i].Id, " - ", ppdbOptions[i].Name, " : q: ", ppdbOptions[i].Quota, " len.std:", len(ppdbOptions[i].PpdbRegistration), " \n ")
		for j := 0; j < len(ppdbOptions[i].PpdbRegistration); j++ {
			fmt.Println(">ori:", j, ":", ppdbOptions[i].PpdbRegistration[j].Name)

			if ppdbOptions[i].Quota > len(ppdbOptions[i].PpdbRegistration) {
				idx := indexOf(ppdbOptions[i].PpdbRegistration[j].SecondChoiceOption, ppdbOptions)
				if idx != -1 {
					ppdbOptions[idx].AddStd(ppdbOptions[i].PpdbRegistration[j])
					ppdbOptions[idx].RemoveStd(j)
				}
			}

		}
	}

	fmt.Println("===========================res==============================")
	for _, opt := range ppdbOptions {
		fmt.Println(opt.Id, " - ", opt.Name, " : q: ", opt.Quota, " len.std:", len(opt.PpdbRegistration), " \n ")
		for i, std := range opt.PpdbRegistration {
			fmt.Println(">ori:", i, ":", std.Name)
		}
	}
}

/*
func displayMovies(header string, movies []*Movie) {
	fmt.Println(header)
	for _, m := range movies {
		fmt.Printf("\t- %s (%d) R:%.1f\n", m.Title, m.Year, m.Rate)
	}
	fmt.Println()
}*/

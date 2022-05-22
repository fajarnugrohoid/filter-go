package repository

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type PpdbRegistrationRepositoyImpl struct {
}

func NewPpdbRegistrationRepositoy() PpdbRegistrationRepositoy {
	return &PpdbRegistrationRepositoyImpl{}
}

func (p *PpdbRegistrationRepositoyImpl) Save(ctx context.Context, database *mongo.Database, category domain.PpdbRegistration) domain.PpdbRegistration {
	//TODO implement me
	panic("implement me")
}

func (p *PpdbRegistrationRepositoyImpl) Update(ctx context.Context, database *mongo.Database, category domain.PpdbRegistration) domain.PpdbRegistration {
	//TODO implement me
	panic("implement me")
}

func (p *PpdbRegistrationRepositoyImpl) Delete(ctx context.Context, database *mongo.Database, category domain.PpdbRegistration) {
	//TODO implement me
	panic("implement me")
}

func (p *PpdbRegistrationRepositoyImpl) GetByFirstChoiceLevel(ctx context.Context, database *mongo.Database, level string, firstChoice primitive.ObjectID) []domain.PpdbRegistration {
	//TODO implement me
	criteria := bson.M{"first_choice_option": firstChoice, "registration_level": level}
	findOptions := options.Find()

	findOptions.SetSort(bson.D{{"distance1", 1}})

	csr, err := database.Collection("ppdb_registrations").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]domain.PpdbRegistration, 0)
	for csr.Next(ctx) {
		var row domain.PpdbRegistration

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := domain.PpdbRegistration{
			Id:                 row.Id,
			Name:               row.Name,
			OptionType:         row.OptionType,
			FirstChoiceOption:  row.FirstChoiceOption,
			SecondChoiceOption: row.SecondChoiceOption,
			ThirdChoiceOption:  row.ThirdChoiceOption,
			Score:              row.Score,
			Distance:           row.Distance1,
			Distance1:          row.Distance1,
			Distance2:          row.Distance2,
			Distance3:          row.Distance3,
			BirthDate:          row.BirthDate,
			AcceptedStatus:     0,
			AcceptedIndex:      0, //perlu di update idx berapa untuk firstchoice
			AcceptedChoiceId:   row.FirstChoiceOption,
		}

		result = append(result, tmp)
	}
	return result
}

func (p *PpdbRegistrationRepositoyImpl) GetAll(ctx context.Context, database *mongo.Database) []domain.PpdbRegistration {
	//TODO implement me
	//TODO implement me
	//var optId = [1]primitive.ObjectID{firstChoice}
	//criteria := bson.M{"first_choice_option": firstChoice, "registration_level": "sma", "status": "fit"}
	criteria := bson.M{}
	findOptions := options.Find()

	findOptions.SetSort(bson.D{{"distance1", 1}})

	csr, err := database.Collection("ppdb_registrations").Find(ctx, criteria, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]domain.PpdbRegistration, 0)
	for csr.Next(ctx) {
		var row domain.PpdbRegistration

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		tmp := domain.PpdbRegistration{
			Id:                 row.Id,
			Name:               row.Name,
			OptionType:         row.OptionType,
			FirstChoiceOption:  row.FirstChoiceOption,
			SecondChoiceOption: row.SecondChoiceOption,
			ThirdChoiceOption:  row.ThirdChoiceOption,
			Score:              row.Score,
			Distance:           row.Distance1,
			Distance1:          row.Distance1,
			Distance2:          row.Distance2,
			Distance3:          row.Distance3,
			BirthDate:          row.BirthDate,
			AcceptedStatus:     0,
			AcceptedIndex:      0, //perlu di update idx berapa untuk firstchoice
			AcceptedChoiceId:   row.FirstChoiceOption,
		}

		result = append(result, tmp)
	}
	return result
}

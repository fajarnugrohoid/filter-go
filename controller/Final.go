package controller

import (
	"context"
	"filterisasi/models/domain"
	"filterisasi/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateFiltered(ctx context.Context, database *mongo.Database, optionTypes map[string][]*domain.PpdbOption) {
	repository.InsertFiltered(ctx, database, optionTypes["abk"], "abk")
	repository.InsertFiltered(ctx, database, optionTypes["ketm"], "ketm")
	repository.InsertFiltered(ctx, database, optionTypes["kondisi-tertentu"], "kondisi-tertentu")
}
func UpdateAllStatistic(ctx context.Context, database *mongo.Database, optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger) {
	UpdateStatisticByOpt(ctx, database, optionTypes, "abk", logger)
	UpdateStatisticByOpt(ctx, database, optionTypes, "ketm", logger)
	UpdateStatisticByOpt(ctx, database, optionTypes, "kondisi-tertentu", logger)
}

func UpdateStatisticByOpt(ctx context.Context, database *mongo.Database, optionTypes map[string][]*domain.PpdbOption, optType string, logger *logrus.Logger) {
	ppdbStatistics := make([]domain.PpdbStatistic, 0)
	for i := 0; i < len(optionTypes[optType]); i++ {
		logger.Info(i, "-", optionTypes[optType][i].Id, " - ", optionTypes[optType][i].Name,
			" : q: ", optionTypes[optType][i].Quota,
			" : p: ", len(optionTypes[optType][i].PpdbRegistration),
			" - needQuota:", optionTypes[optType][i].NeedQuota,
			" - AddQuota:", optionTypes[optType][i].AddQuota,
		)
		for i, std := range optionTypes[optType][i].PpdbRegistration {
			logger.Info(">", i, ":", std.Name,
				" - acc:", std.AcceptedStatus,
				" - accId:", std.AcceptedChoiceId,
				" distance: ", std.Distance, " Birth:", std.BirthDate)
		}
		/*
			for i, std := range optionTypes["ketm"][i].RegistrationHistory {
				fmt.Println("hist>", i, ":", std.Name, " - acc:", std.AcceptedIndex)
			}
			for i, std := range optionTypes["ketm"][i].HistoryShifting {
				fmt.Println("shift>", i, ":", std.Name, " - acc:", std.AcceptedIndex)
			}*/

		var pg float64
		if len(optionTypes[optType][i].PpdbRegistration) > 0 {
			pg = optionTypes[optType][i].PpdbRegistration[len(optionTypes[optType][i].PpdbRegistration)-1].Distance
		} else {
			pg = 0
		}
		tmpStatistic := domain.PpdbStatistic{
			Id:         optionTypes[optType][i].Id,
			Name:       optionTypes[optType][i].Name,
			OptionType: optionTypes[optType][i].Type,
			Quota:      optionTypes[optType][i].Quota,
			SchoolId:   optionTypes[optType][i].SchoolId,
			Pg:         pg,
		}
		ppdbStatistics = append(ppdbStatistics, tmpStatistic)
	}
	repository.InsertStatistic(ctx, database, ppdbStatistics, optType)
}

func UpdateFilteredStatistic(ctx context.Context, database *mongo.Database, optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger) {
	UpdateFiltered(ctx, database, optionTypes)
	UpdateAllStatistic(ctx, database, optionTypes, logger)
}

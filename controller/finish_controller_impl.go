package controller

import (
	"context"
	"filterisasi/models/domain"
	"filterisasi/service"
	"github.com/sirupsen/logrus"
)

type FinishControllerImpl struct {
	PpdbFilteredService service.PpdbFilteredService
}

func NewFinishController(ppdbFilteredService service.PpdbFilteredService) FinishController {
	return &FinishControllerImpl{PpdbFilteredService: ppdbFilteredService}
}

func (controller FinishControllerImpl) UpdateFiltered(ctx context.Context, optionTypes map[string][]*domain.PpdbOption) {
	//TODO implement me
	controller.PpdbFilteredService.Save(ctx, optionTypes["abk"], "abk")
	controller.PpdbFilteredService.Save(ctx, optionTypes["ketm"], "ketm")
	controller.PpdbFilteredService.Save(ctx, optionTypes["kondisi-tertentu"], "kondisi-tertentu")
}

func (controller FinishControllerImpl) UpdateAllStatistic(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger) {
	//TODO implement me
	//UpdateStatisticByOpt(ctx, database, optionTypes, "abk", logger)
	//UpdateStatisticByOpt(ctx, database, optionTypes, "ketm", logger)
	//UpdateStatisticByOpt(ctx, database, optionTypes, "kondisi-tertentu", logger)
}

func (controller FinishControllerImpl) UpdateStatisticByOpt(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, optType string, logger *logrus.Logger) {
	//TODO implement me
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
	//repository.InsertStatistic(ctx, database, ppdbStatistics, optType)
}

func (controller FinishControllerImpl) UpdateFilteredStatistic(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger) {
	//TODO implement me
	controller.UpdateFiltered(ctx, optionTypes)
	controller.UpdateAllStatistic(ctx, optionTypes, logger)
}

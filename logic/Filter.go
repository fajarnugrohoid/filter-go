package logic

import (
	"filterisasi/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

/*
func (ppdbOptions *models.PpdbOptionList) ProcessFilter(status bool) []models.PpdbOption {

	for i := 0; i < len(ppdbOptions); i++ {
		if ppdbOptions[i].Filtered == 0 {
			sort.Sort(models.ByScore(ppdbOptions[i].PpdbRegistration))
			fmt.Println(ppdbOptions[i].Id, " - ", ppdbOptions[i].Name,
				" len.std:", len(ppdbOptions[i].PpdbRegistration),
				" : q: ", ppdbOptions[i].Quota, " \n ")

			if len(ppdbOptions[i].PpdbRegistration) > ppdbOptions[i].Quota {

				for j := ppdbOptions[i].Quota; j < len(ppdbOptions[i].PpdbRegistration); j++ {
					idx := -1
					if ppdbOptions[i].PpdbRegistration[j].AcceptedStatus == 0 {
						ppdbOptions[i].PpdbRegistration[j].AcceptedStatus = 1
						idx := FindIndex(ppdbOptions[i].PpdbRegistration[j].SecondChoiceOption, ppdbOptions)
						fmt.Println(">ori:", j, ":", ppdbOptions[i].PpdbRegistration[j].Name, "-", ppdbOptions[i].PpdbRegistration[j].SecondChoiceOption, " - ", idx)
					} else if ppdbOptions[i].PpdbRegistration[j].AcceptedStatus == 1 {
						ppdbOptions[i].PpdbRegistration[j].AcceptedStatus = 2
						idx := FindIndex(ppdbOptions[i].PpdbRegistration[j].ThirdChoiceOption, ppdbOptions)
						fmt.Println(">ori:", j, ":", ppdbOptions[i].PpdbRegistration[j].Name, "-", ppdbOptions[i].PpdbRegistration[j].SecondChoiceOption, " - ", idx)
					}
					if idx == -1 {
						ppdbOptions[i].PpdbRegistration[j].AcceptedStatus = 3
						ppdbOptions[len(ppdbOptions)-1].AddStd(ppdbOptions[i].PpdbRegistration[j])
						ppdbOptions[i].RemoveStd(j)
						j--
					} else {

						ppdbOptions[idx].AddStd(ppdbOptions[i].PpdbRegistration[j])
						ppdbOptions[i].RemoveStd(j)
						j--
						ppdbOptions[idx].Filtered = 0
						status = true
					}
				}

			}
			ppdbOptions[i].Filtered = 1
		}
	}

	if status == true {
		return ProcessFilter(ppdbOptions, false)
	}
	return ppdbOptions
} */
func DoFilter(optionTypes map[string][]*models.PpdbOption, logger *logrus.Logger) map[string][]*models.PpdbOption {
	optionTypes["abk"] = Filter1Options(optionTypes, "abk", 0, logger)
	for i := 0; i < len(optionTypes["abk"]); i++ {
		fmt.Println("", optionTypes["abk"][i].Name, "-Q:", optionTypes["abk"][i].Quota, "R:", len(optionTypes["abk"][i].PpdbRegistration))
	}
	SendQuota(optionTypes, "abk", "ketm", logger)
	//os.Exit(0)
	optionTypes = Doing2OptionsShareQuota(optionTypes, logger)
	return optionTypes
}

func Doing2OptionsShareQuota(optionTypes map[string][]*models.PpdbOption, logger *logrus.Logger) map[string][]*models.PpdbOption {
	optionTypes["ketm"] = Filter2OptionsShareQuota(optionTypes, "ketm", 0, logger)
	optionTypes["kondisi-tertentu"] = Filter2OptionsShareQuota(optionTypes, "kondisi-tertentu", 0, logger)

	//share quota
	var reFilterKondisiTertentu, reFilterKetm bool

	optionTypes, reFilterKetm = CheckQuota(optionTypes, "ketm", "kondisi-tertentu", false, logger)
	optionTypes, reFilterKondisiTertentu = CheckQuota(optionTypes, "kondisi-tertentu", "ketm", false, logger)

	optionTypes["ketm"] = Filter2OptionsShareQuota(optionTypes, "ketm", 1, logger)
	optionTypes["kondisi-tertentu"] = Filter2OptionsShareQuota(optionTypes, "kondisi-tertentu", 1, logger)

	if reFilterKondisiTertentu == true || reFilterKetm == true {
		return DoFilter(optionTypes, logger)
	}

	return optionTypes
}

func Filter2OptionsShareQuota(optionTypes map[string][]*models.PpdbOption, optType string, loop int, logger *logrus.Logger) []*models.PpdbOption {

	runtime.GOMAXPROCS(2)
	var messages = make(chan []*models.PpdbOption)
	logger.Info("Filter2OptionsShareQuota")

	var getFiltered = func(objs chan []*models.PpdbOption, option []*models.PpdbOption) {
		logger.Info("bef getFiltered:")
		ppdbOptions := ProcessFilter(option, true, loop, logger)
		logger.Info("aft getFiltered")
		messages <- ppdbOptions
	}
	logger.Info("end getFiltered")
	if optType == "ketm" {
		go getFiltered(messages, optionTypes["ketm"])
		logger.Info("go done getFiltered")

	} else if optType == "kondisi-tertentu" {
		go getFiltered(messages, optionTypes["kondisi-tertentu"])
		logger.Info("go done getFiltered")

	}
	logger.Info("go done getFiltered")
	data := <-messages // read from channel a
	close(messages)

	logger.Info("close messages")

	//fmt.Println(repositories)

	return data
}

func Filter1Options(optionTypes map[string][]*models.PpdbOption, optType string, loop int, logger *logrus.Logger) []*models.PpdbOption {

	runtime.GOMAXPROCS(1)
	var messages = make(chan []*models.PpdbOption)
	logger.Info("Filter1Options")

	var getFiltered = func(objs chan []*models.PpdbOption, option []*models.PpdbOption) {
		logger.Info("bef getFiltered:")
		ppdbOptions := ProcessFilter(option, true, loop, logger)
		logger.Info("aft getFiltered")
		messages <- ppdbOptions
	}
	logger.Info("end getFiltered")
	if optType == "abk" {
		go getFiltered(messages, optionTypes["abk"])
		logger.Info("go done getFiltered abk")
	}
	logger.Info("go done getFiltered")
	data := <-messages // read from channel a
	close(messages)

	logger.Info("close messages")

	return data
}

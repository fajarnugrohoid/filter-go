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
func DoFilterSenior(optionTypes map[string][]*models.PpdbOption, logger *logrus.Logger) map[string][]*models.PpdbOption {
	optionTypes["abk"] = Filter1Options(optionTypes, "abk", 0, logger) //no.1
	for i := 0; i < len(optionTypes["abk"]); i++ {
		fmt.Println("", optionTypes["abk"][i].Name, "-Q:", optionTypes["abk"][i].Quota, "R:", len(optionTypes["abk"][i].PpdbRegistration))
	}
	SendAllQuota(optionTypes, "abk", "ketm", logger) //no.1
	//os.Exit(0)
	optionTypes = Doing2OptionsShareQuota(optionTypes, "ketm", "kondisi-tertentu", logger) //no.1

	optionTypes = Doing2OptionsShareQuota(optionTypes, "anak-guru", "perpindahan", logger) //no.2

	optionTypes, _ = CheckQuota(optionTypes, "anak-guru", "kondisi-tertentu", false, logger)   //no.4
	optionTypes, _ = CheckQuota(optionTypes, "anak-guru", "ketm", false, logger)               //no.4
	optionTypes, _ = CheckQuota(optionTypes, "perpindahan", "kondisi-tertentu", false, logger) //no.4
	optionTypes, _ = CheckQuota(optionTypes, "perpindahan", "ketm", false, logger)
	optionTypes["anak-guru"] = Filter1Options(optionTypes, "anak-guru", 0, logger)     //no.4
	optionTypes["perpindahan"] = Filter1Options(optionTypes, "perpindahan", 0, logger) //no.4

	SendAllQuota(optionTypes, "ketm", "prestasi-rapor", logger)             //no.5
	SendAllQuota(optionTypes, "kondisi-tertentu", "prestasi-rapor", logger) //no.5

	optionTypes = Doing2OptionsShareQuota(optionTypes, "prestasi-rapor", "perlombaan", logger) //no.3

	//tinggal semua sisa dikirim ke tahap 2 (zonasi)

	return optionTypes
}

func Doing2OptionsShareQuota(optionTypes map[string][]*models.PpdbOption, optType1 string, optType2 string, logger *logrus.Logger) map[string][]*models.PpdbOption {
	optionTypes[optType1] = RunFilter(optionTypes, optType1, 0, logger)
	optionTypes[optType2] = RunFilter(optionTypes, optType2, 0, logger)

	for i := 0; i < len(optionTypes[optType1]); i++ {
		fmt.Println("", optionTypes[optType1][i].Name, "-Q:", optionTypes[optType1][i].Quota, "R:", len(optionTypes[optType1][i].PpdbRegistration))
	}
	for i := 0; i < len(optionTypes[optType2]); i++ {
		fmt.Println("", optionTypes[optType2][i].Name, "-Q:", optionTypes[optType2][i].Quota, "R:", len(optionTypes[optType2][i].PpdbRegistration))
	}

	//share quota
	var reFilterKondisiTertentu, reFilterKetm bool

	optionTypes, reFilterKetm = CheckQuota(optionTypes, optType1, optType2, false, logger)
	optionTypes, reFilterKondisiTertentu = CheckQuota(optionTypes, optType2, optType1, false, logger)

	optionTypes[optType1] = RunFilter(optionTypes, optType1, 1, logger)
	optionTypes[optType2] = RunFilter(optionTypes, optType2, 1, logger)

	if reFilterKondisiTertentu == true || reFilterKetm == true {
		return Doing2OptionsShareQuota(optionTypes, optType1, optType2, logger)

	}

	return optionTypes
}

func RunFilter(optionTypes map[string][]*models.PpdbOption, optType string, loop int, logger *logrus.Logger) []*models.PpdbOption {

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

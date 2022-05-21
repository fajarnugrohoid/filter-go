package logic

import (
	"filterisasi/models/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

func Doing2OptionsShareQuota(optionTypes map[string][]*domain.PpdbOption, optType1 string, optType2 string, logger *logrus.Logger) map[string][]*domain.PpdbOption {
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

	optionTypes, reFilterKetm = NeedSuplyQuota1Option(optionTypes, optType1, optType2, false, logger)
	optionTypes, reFilterKondisiTertentu = NeedSuplyQuota1Option(optionTypes, optType2, optType1, false, logger)

	optionTypes[optType1] = RunFilter(optionTypes, optType1, 1, logger)
	optionTypes[optType2] = RunFilter(optionTypes, optType2, 1, logger)

	if reFilterKondisiTertentu == true || reFilterKetm == true {
		return Doing2OptionsShareQuota(optionTypes, optType1, optType2, logger)

	}

	return optionTypes
}

func Filter3OptionsShareQuota(optionTypes map[string][]*domain.PpdbOption, optType1 string, optType2 string, optType3 string, logger *logrus.Logger) map[string][]*domain.PpdbOption {
	optionTypes[optType1] = RunFilter(optionTypes, optType1, 0, logger)
	optionTypes[optType2] = RunFilter(optionTypes, optType2, 0, logger)
	optionTypes[optType3] = RunFilter(optionTypes, optType3, 0, logger)

	for i := 0; i < len(optionTypes[optType1]); i++ {
		fmt.Println("", optionTypes[optType1][i].Name, "-Q:", optionTypes[optType1][i].Quota, "R:", len(optionTypes[optType1][i].PpdbRegistration))
	}
	for i := 0; i < len(optionTypes[optType2]); i++ {
		fmt.Println("", optionTypes[optType2][i].Name, "-Q:", optionTypes[optType2][i].Quota, "R:", len(optionTypes[optType2][i].PpdbRegistration))
	}

	//share quota
	var reFilterFirstOpt, reFilterSecondOpt, reFilterThirdOpt bool

	optionTypes, reFilterFirstOpt = NeedSupplyQuota2Options(optionTypes, optType1, optType2, optType3, false, logger)
	optionTypes, reFilterSecondOpt = NeedSupplyQuota2Options(optionTypes, optType2, optType1, optType3, false, logger)
	optionTypes, reFilterThirdOpt = NeedSupplyQuota2Options(optionTypes, optType3, optType1, optType2, false, logger)

	optionTypes[optType1] = RunFilter(optionTypes, optType1, 1, logger)
	optionTypes[optType2] = RunFilter(optionTypes, optType2, 1, logger)
	optionTypes[optType3] = RunFilter(optionTypes, optType3, 1, logger)

	if reFilterFirstOpt == true || reFilterSecondOpt == true || reFilterThirdOpt == true {
		return Filter3OptionsShareQuota(optionTypes, optType1, optType2, optType3, logger)

	}

	return optionTypes
}

func RunFilter(optionTypes map[string][]*domain.PpdbOption, optType string, loop int, logger *logrus.Logger) []*domain.PpdbOption {

	runtime.GOMAXPROCS(2)
	var messages = make(chan []*domain.PpdbOption)
	logger.Info("Filter2OptionsShareQuota")

	var getFiltered = func(objs chan []*domain.PpdbOption, option []*domain.PpdbOption) {
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

	//fmt.Println(controller)

	return data
}

func Filter1Options(optionTypes map[string][]*domain.PpdbOption, optType string, loop int, logger *logrus.Logger) []*domain.PpdbOption {

	runtime.GOMAXPROCS(1)
	var messages = make(chan []*domain.PpdbOption)
	logger.Info("Filter1Options")

	var getFiltered = func(objs chan []*domain.PpdbOption, option []*domain.PpdbOption) {
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

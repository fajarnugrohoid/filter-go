package logic

import (
	"filterisasi/models/domain"
	"fmt"
	"github.com/sirupsen/logrus"
)

func DoFilterVocational(optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger) map[string][]*domain.PpdbOption {
	optionTypes["abk"] = Filter1Options(optionTypes, "abk", 0, logger) //no.1
	for i := 0; i < len(optionTypes["abk"]); i++ {
		fmt.Println("", optionTypes["abk"][i].Name, "-Q:", optionTypes["abk"][i].Quota, "R:", len(optionTypes["abk"][i].PpdbRegistration))
	}
	SendAllQuota(optionTypes, "abk", "ketm", logger) //no.1
	//os.Exit(0)
	optionTypes = Doing2OptionsShareQuota(optionTypes, "ketm", "kondisi-tertentu", logger) //no.1

	optionTypes = Filter3OptionsShareQuota(optionTypes, "prioritas-terdekat", "ketm", "kondisi-tertentu", logger)

	return optionTypes
}

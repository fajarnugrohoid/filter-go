package logic

import (
	"filterisasi/models/domain"
	"fmt"
	"github.com/sirupsen/logrus"
)

func DoFilterSenior(optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger) map[string][]*domain.PpdbOption {
	optionTypes["abk"] = Filter1Options(optionTypes, "abk", 0, logger) //no.1
	for i := 0; i < len(optionTypes["abk"]); i++ {
		fmt.Println("", optionTypes["abk"][i].Name, "-Q:", optionTypes["abk"][i].Quota, "R:", len(optionTypes["abk"][i].PpdbRegistration))
	}
	SendAllQuota(optionTypes, "abk", "ketm", logger) //no.1
	//os.Exit(0)
	optionTypes = Doing2OptionsShareQuota(optionTypes, "ketm", "kondisi-tertentu", logger) //no.1

	optionTypes = Doing2OptionsShareQuota(optionTypes, "anak-guru", "perpindahan", logger) //no.2

	optionTypes, _ = NeedSuplyQuota1Option(optionTypes, "anak-guru", "kondisi-tertentu", false, logger)   //no.4
	optionTypes, _ = NeedSuplyQuota1Option(optionTypes, "anak-guru", "ketm", false, logger)               //no.4
	optionTypes, _ = NeedSuplyQuota1Option(optionTypes, "perpindahan", "kondisi-tertentu", false, logger) //no.4
	optionTypes, _ = NeedSuplyQuota1Option(optionTypes, "perpindahan", "ketm", false, logger)
	optionTypes["anak-guru"] = Filter1Options(optionTypes, "anak-guru", 0, logger)     //no.4
	optionTypes["perpindahan"] = Filter1Options(optionTypes, "perpindahan", 0, logger) //no.4

	SendAllQuota(optionTypes, "ketm", "prestasi-rapor", logger)             //no.5
	SendAllQuota(optionTypes, "kondisi-tertentu", "prestasi-rapor", logger) //no.5

	optionTypes = Doing2OptionsShareQuota(optionTypes, "prestasi-rapor", "perlombaan", logger) //no.3

	//tinggal semua sisa dikirim ke tahap 2 (zonasi)

	return optionTypes
}

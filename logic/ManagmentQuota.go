package logic

import (
	"filterisasi/models/domain"
	"github.com/sirupsen/logrus"
)

func SendAllQuota(optionTypes map[string][]*domain.PpdbOption, currentType string, targetType string, logger *logrus.Logger) {
	for i := 0; i < len(optionTypes[currentType]); i++ {
		if optionTypes[currentType][i].Quota > len(optionTypes[currentType][i].PpdbRegistration) {
			sisaQuota := optionTypes[currentType][i].Quota - len(optionTypes[currentType][i].PpdbRegistration)
			optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota + sisaQuota
			optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota - sisaQuota
			optionTypes[currentType][i].AddQuota += sisaQuota
		}
	}
}

func NeedSuplyQuota1Option(optionTypes map[string][]*domain.PpdbOption, currentType string, targetType string, reFilter bool, logger *logrus.Logger) (map[string][]*domain.PpdbOption, bool) {
	logger.Info("===========================need quota==============================")

	for i := 0; i < len(optionTypes[currentType]); i++ {
		logger.Debug(i, "-", optionTypes[currentType][i].Id, " - ",
			optionTypes[currentType][i].Name,
			" - q:", optionTypes[currentType][i].Quota,
			" - Need:", optionTypes[currentType][i].NeedQuota,
			" - Added:", optionTypes[currentType][i].AddQuota)

		if i == len(optionTypes[currentType])-1 {
			continue
		}

		if optionTypes[currentType][i].NeedQuota > 0 {

			sisaQuota := optionTypes[targetType][i].Quota - len(optionTypes[targetType][i].PpdbRegistration)
			logger.Debug("sisaQuota:", sisaQuota)
			if sisaQuota > 0 {

				/*
					if optionTypes[currentType][i].NeedQuota >= sisaQuota {
						optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + sisaQuota
						optionTypes[currentType][i].NeedQuota = optionTypes[currentType][i].NeedQuota - sisaQuota
						optionTypes[currentType][i].AddQuota += sisaQuota
						optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - sisaQuota
						logger.Debug("targetType.quota:", optionTypes[targetType][i].Quota)
					} else {
						//jika kebutuhan kuota hanya sedikit artinya hanya membutuhkan sedikit dari sisa kuota lawannya
						logger.Debug("sisaQuota:", sisaQuota, " - NeedQuota", optionTypes[currentType][i].NeedQuota)
						optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + optionTypes[currentType][i].NeedQuota
						optionTypes[currentType][i].AddQuota += optionTypes[currentType][i].NeedQuota
						optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - optionTypes[currentType][i].NeedQuota
						optionTypes[currentType][i].NeedQuota = 0

					}*/
				optionTypes, _ = runCalculateQuota(optionTypes, currentType, targetType, i, sisaQuota, logger)

				optionTypes[currentType][i].Filtered = 0
				optionTypes[currentType][i].UpdateQuota = true
				optionTypes[currentType] = PullStudentToFirstChoice(optionTypes[currentType], i, logger)
				reFilter = true
			}
		}
	}

	logger.Debug("optionType1:", currentType)
	for i := 0; i < len(optionTypes[currentType]); i++ {
		logger.Debug(i, "-", optionTypes[currentType][i].Id, " - ", optionTypes[currentType][i].Name,
			" : q: ", optionTypes[currentType][i].Quota,
			" : p: ", len(optionTypes[currentType][i].PpdbRegistration),
			" - needQuota:", optionTypes[currentType][i].NeedQuota,
			" - AddQuota:", optionTypes[currentType][i].AddQuota,
		)
	}
	logger.Debug("optionType2:", targetType)
	for i := 0; i < len(optionTypes[targetType]); i++ {
		logger.Debug(i, "-", optionTypes[targetType][i].Id, " - ", optionTypes[targetType][i].Name,
			" : q: ", optionTypes[targetType][i].Quota,
			" : p: ", len(optionTypes[targetType][i].PpdbRegistration),
			" - needQuota:", optionTypes[targetType][i].NeedQuota,
			" - AddQuota:", optionTypes[targetType][i].AddQuota,
		)
	}

	return optionTypes, reFilter
}

func NeedSupplyQuota2Options(optionTypes map[string][]*domain.PpdbOption, currentType string, targetType1 string, targetType2 string, reFilter bool, logger *logrus.Logger) (map[string][]*domain.PpdbOption, bool) {
	logger.Info("===========================need quota==============================")

	for i := 0; i < len(optionTypes[currentType]); i++ {
		logger.Debug(i, "-", optionTypes[currentType][i].Id, " - ",
			optionTypes[currentType][i].Name,
			" - q:", optionTypes[currentType][i].Quota,
			" - Need:", optionTypes[currentType][i].NeedQuota,
			" - Added:", optionTypes[currentType][i].AddQuota)

		if i == len(optionTypes[currentType])-1 {
			continue
		}

		if optionTypes[currentType][i].NeedQuota > 0 {

			sisaQuota1 := optionTypes[targetType1][i].Quota - len(optionTypes[targetType1][i].PpdbRegistration)
			sisaQuota2 := optionTypes[targetType2][i].Quota - len(optionTypes[targetType2][i].PpdbRegistration)
			logger.Debug("sisaQuota:", sisaQuota1)
			var stillNeedQuota bool
			if sisaQuota1 > 0 || sisaQuota2 > 0 {

				optionTypes, stillNeedQuota = runCalculateQuota(optionTypes, currentType, targetType1, i, sisaQuota1, logger)
				if stillNeedQuota {
					optionTypes, stillNeedQuota = runCalculateQuota(optionTypes, currentType, targetType2, i, sisaQuota2, logger)
				}

				optionTypes[currentType][i].Filtered = 0
				optionTypes[currentType][i].UpdateQuota = true
				optionTypes[currentType] = PullStudentToFirstChoice(optionTypes[currentType], i, logger)
				reFilter = true
			}
		}
	}

	logger.Debug("optionType1:", currentType)
	for i := 0; i < len(optionTypes[currentType]); i++ {
		logger.Debug(i, "-", optionTypes[currentType][i].Id, " - ", optionTypes[currentType][i].Name,
			" : q: ", optionTypes[currentType][i].Quota,
			" : p: ", len(optionTypes[currentType][i].PpdbRegistration),
			" - needQuota:", optionTypes[currentType][i].NeedQuota,
			" - AddQuota:", optionTypes[currentType][i].AddQuota,
		)
	}
	logger.Debug("optionType2:", targetType1)
	for i := 0; i < len(optionTypes[targetType1]); i++ {
		logger.Debug(i, "-", optionTypes[targetType1][i].Id, " - ", optionTypes[targetType1][i].Name,
			" : q: ", optionTypes[targetType1][i].Quota,
			" : p: ", len(optionTypes[targetType1][i].PpdbRegistration),
			" - needQuota:", optionTypes[targetType1][i].NeedQuota,
			" - AddQuota:", optionTypes[targetType1][i].AddQuota,
		)
	}

	return optionTypes, reFilter
}

func runCalculateQuota(optionTypes map[string][]*domain.PpdbOption, currentType string, targetType string, i int, sisaQuota int, logger *logrus.Logger) (map[string][]*domain.PpdbOption, bool) {
	var stillNeedQuota bool
	if optionTypes[currentType][i].NeedQuota >= sisaQuota {
		optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + sisaQuota
		optionTypes[currentType][i].NeedQuota = optionTypes[currentType][i].NeedQuota - sisaQuota
		optionTypes[currentType][i].AddQuota += sisaQuota
		optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - sisaQuota
		logger.Debug("targetType.quota:", optionTypes[targetType][i].Quota)
		stillNeedQuota = true
	} else {
		//jika kebutuhan kuota hanya sedikit artinya hanya membutuhkan sedikit dari sisa kuota lawannya
		logger.Debug("sisaQuota:", sisaQuota, " - NeedQuota:", optionTypes[currentType][i].NeedQuota)
		optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + optionTypes[currentType][i].NeedQuota
		optionTypes[currentType][i].AddQuota += optionTypes[currentType][i].NeedQuota
		optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - optionTypes[currentType][i].NeedQuota
		optionTypes[currentType][i].NeedQuota = 0
		stillNeedQuota = false
	}
	return optionTypes, stillNeedQuota
}

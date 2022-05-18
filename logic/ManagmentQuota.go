package logic

import (
	"filterisasi/models"
	"github.com/sirupsen/logrus"
)

func SendQuota(optionTypes map[string][]*models.PpdbOption, currentType string, targetType string, logger *logrus.Logger) {
	for i := 0; i < len(optionTypes[currentType]); i++ {
		if optionTypes[currentType][i].Quota > len(optionTypes[currentType][i].PpdbRegistration) {
			sisaQuota := optionTypes[currentType][i].Quota - len(optionTypes[currentType][i].PpdbRegistration)
			optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota + sisaQuota
			optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota - sisaQuota
			optionTypes[currentType][i].AddQuota += sisaQuota
		}
	}
}

func CheckQuota(optionTypes map[string][]*models.PpdbOption, currentType string, targetType string, reFilter bool, logger *logrus.Logger) (map[string][]*models.PpdbOption, bool) {
	logger.Info("===========================need quota==============================")

	for i := 0; i < len(optionTypes[currentType]); i++ {
		logger.Debug(i, "-", optionTypes[currentType][i].Id, " - ",
			optionTypes[currentType][i].Name,
			" - q:", optionTypes[currentType][i].Quota,
			" - Need:", optionTypes[currentType][i].NeedQuotaFirstOpt,
			" - Added:", optionTypes[currentType][i].AddQuota)

		if i == len(optionTypes[currentType])-1 {
			continue
		}

		if optionTypes[currentType][i].NeedQuotaFirstOpt > 0 {

			sisaQuota := optionTypes[targetType][i].Quota - len(optionTypes[targetType][i].PpdbRegistration)
			logger.Debug("sisaQuota:", sisaQuota)
			if sisaQuota > 0 {

				if optionTypes[currentType][i].NeedQuotaFirstOpt >= sisaQuota {
					optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + sisaQuota
					optionTypes[currentType][i].NeedQuotaFirstOpt = optionTypes[currentType][i].NeedQuotaFirstOpt - sisaQuota
					optionTypes[currentType][i].AddQuota += sisaQuota
					optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - sisaQuota
					logger.Debug("targetType.quota:", optionTypes[targetType][i].Quota)
				} else {
					//jika kebutuhan kuota hanya sedikit artinya hanya membutuhkan sedikit dari sisa kuota lawannya
					logger.Debug("sisaQuota:", sisaQuota, " - NeedQuotaFirstOpt:", optionTypes[currentType][i].NeedQuotaFirstOpt)
					optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + optionTypes[currentType][i].NeedQuotaFirstOpt
					optionTypes[currentType][i].AddQuota += optionTypes[currentType][i].NeedQuotaFirstOpt
					optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - optionTypes[currentType][i].NeedQuotaFirstOpt
					optionTypes[currentType][i].NeedQuotaFirstOpt = 0

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
			" - needQuota:", optionTypes[currentType][i].NeedQuotaFirstOpt,
			" - AddQuota:", optionTypes[currentType][i].AddQuota,
		)
	}
	logger.Debug("optionType2:", targetType)
	for i := 0; i < len(optionTypes[targetType]); i++ {
		logger.Debug(i, "-", optionTypes[targetType][i].Id, " - ", optionTypes[targetType][i].Name,
			" : q: ", optionTypes[targetType][i].Quota,
			" : p: ", len(optionTypes[targetType][i].PpdbRegistration),
			" - needQuota:", optionTypes[targetType][i].NeedQuotaFirstOpt,
			" - AddQuota:", optionTypes[targetType][i].AddQuota,
		)
	}

	return optionTypes, reFilter
}

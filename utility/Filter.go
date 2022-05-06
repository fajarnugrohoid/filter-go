package utility

import (
	"filterisasi/models"
	"fmt"
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

func DoFilter(optionTypes map[string][]*models.PpdbOption) map[string][]*models.PpdbOption {
	optionTypes["ketm"] = Filter2OptionsShareQuota(optionTypes, "ketm")
	optionTypes["kondisi-tertentu"] = Filter2OptionsShareQuota(optionTypes, "kondisi-tertentu")

	fmt.Println("===========================res==============================")
	for _, opt := range optionTypes["ketm"] {
		fmt.Println(opt.Id, " - ", opt.Name, " : q: ", opt.Quota, " len.std:", len(opt.PpdbRegistration), "")
		for i, std := range opt.PpdbRegistration {
			fmt.Println(">ori:", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
				" distance1: ", std.Distance1)
		}
		for i, std := range opt.RegistrationHistory {
			fmt.Println(">hist2:", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
				" AcceptedIndex: ", std.AcceptedIndex)
		}
		fmt.Println("\n")
	}

	for _, opt := range optionTypes["kondisi-tertentu"] {
		fmt.Println(opt.Id, " - ", opt.Name, " : q: ", opt.Quota, " len.std:", len(opt.PpdbRegistration), "")
		for i, std := range opt.PpdbRegistration {
			fmt.Println(">ori:", i, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1)
		}
		fmt.Println("\n")
	}

	//share quota
	optionTypes = CheckQuota(optionTypes, "ketm", "kondisi-tertentu")
	optionTypes = CheckQuota(optionTypes, "kondisi-tertentu", "ketm")

	optionTypes["ketm"] = Filter2OptionsShareQuota(optionTypes, "ketm")
	optionTypes["kondisi-tertentu"] = Filter2OptionsShareQuota(optionTypes, "kondisi-tertentu")

	return optionTypes
}

func CheckQuota(optionTypes map[string][]*models.PpdbOption, currentType string, targetType string) map[string][]*models.PpdbOption {
	fmt.Println("===========================need quota==============================")

	for i := 0; i < len(optionTypes[currentType]); i++ {
		fmt.Println(i, "-", optionTypes[currentType][i].Id, " - ", optionTypes[currentType][i].Name, " : q: ", optionTypes[currentType][i].Quota, " - needQuota:", optionTypes[currentType][i].IsNeedQuota)

		if i == len(optionTypes[currentType])-1 {
			continue
		}

		if optionTypes[currentType][i].NeedQuota > 0 {

			sisaQuota := optionTypes[targetType][i].Quota - len(optionTypes[targetType][i].PpdbRegistration)
			if sisaQuota > 0 {

				if optionTypes[currentType][i].NeedQuota > sisaQuota {
					optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + sisaQuota
					optionTypes[currentType][i].NeedQuota = optionTypes[currentType][i].NeedQuota - sisaQuota
					optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - sisaQuota

				} else {
					//jika kebutuhan kuota hanya sedikit artinya hanya membutuhkan sedikit dari sisa kuota lawannya
					quotaNeeded := sisaQuota - optionTypes[currentType][i].NeedQuota
					optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + quotaNeeded
					optionTypes[currentType][i].NeedQuota = 0
					optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - quotaNeeded
				}

				optionTypes[currentType][i].Filtered = 0
				optionTypes[currentType] = PullStudentToFirstChoice(optionTypes[currentType], i)

			}
		}
	}
	return optionTypes
}

func PullStudentToFirstChoice(optionList []*models.PpdbOption, currTargetIdxOpt int) []*models.PpdbOption {
	for j := 0; j < len(optionList[currTargetIdxOpt].RegistrationHistory); j++ {
		if optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus != 0 {

			var targetIdxStd int
			var nextTargetIdxOpt int
			if optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex == -1 {
				nextTargetIdxOpt = len(optionList) - 1
				targetIdxStd = models.FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)

				fmt.Println("Yg tidak diterima == :",
					optionList[currTargetIdxOpt].RegistrationHistory[j].Id,
					" - ", optionList[currTargetIdxOpt].RegistrationHistory[j].Name,
					" - AccStatus:", optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus,
					" - targetIdxOpt:", nextTargetIdxOpt,
					" - TargetIdxStd:", targetIdxStd,
				)
			} else {
				nextTargetIdxOpt = optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex
				targetIdxStd = models.FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)
				fmt.Println("Yg tidak diterima !=:",
					optionList[currTargetIdxOpt].RegistrationHistory[j].Id,
					" - ", optionList[currTargetIdxOpt].RegistrationHistory[j].Name,
					" - AccStatus:", optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus,
					" - targetIdxOpt:", nextTargetIdxOpt,
					" - TargetIdxStd:", targetIdxStd,
				)
			}

			optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedStatus = 0
			optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus = 0

			optionList[currTargetIdxOpt].AddStd(optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd])

			optionList[nextTargetIdxOpt].RemoveStd(targetIdxStd)
			if nextTargetIdxOpt != len(optionList)-1 {
				optionList[nextTargetIdxOpt].Filtered = 0
				return PullStudentToFirstChoice(optionList, nextTargetIdxOpt)
			}

		}
	}

	return optionList
}

func Filter2OptionsShareQuota(optionTypes map[string][]*models.PpdbOption, optType string) []*models.PpdbOption {

	runtime.GOMAXPROCS(2)
	var messages = make(chan []*models.PpdbOption)
	fmt.Println("Filter2OptionsShareQuota")

	var getFiltered = func(objs chan []*models.PpdbOption, option []*models.PpdbOption) {
		fmt.Println("bef getFiltered:")
		ppdbOptions := models.ProcessFilter(option, true)
		fmt.Println("aft getFiltered")
		messages <- ppdbOptions
	}
	fmt.Println("end getFiltered")
	if optType == "ketm" {
		go getFiltered(messages, optionTypes["ketm"])
		fmt.Println("go done getFiltered")

	} else if optType == "kondisi-tertentu" {
		go getFiltered(messages, optionTypes["kondisi-tertentu"])
		fmt.Println("go done getFiltered")

	}
	fmt.Println("go done getFiltered")
	data := <-messages // read from channel a
	close(messages)

	fmt.Println("close messages")

	//fmt.Println(data)

	return data
}

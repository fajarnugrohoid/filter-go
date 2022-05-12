package models

import (
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

func DoFilter(optionTypes map[string][]*PpdbOption) map[string][]*PpdbOption {
	optionTypes["ketm"] = Filter2OptionsShareQuota(optionTypes, "ketm", 0)
	optionTypes["kondisi-tertentu"] = Filter2OptionsShareQuota(optionTypes, "kondisi-tertentu", 0)

	/*
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
		}*/

	//share quota
	var reFilterKondisiTertentu, reFilterKetm bool

	optionTypes, reFilterKetm = CheckQuota(optionTypes, "ketm", "kondisi-tertentu", false)
	optionTypes, reFilterKondisiTertentu = CheckQuota(optionTypes, "kondisi-tertentu", "ketm", false)

	optionTypes["ketm"] = Filter2OptionsShareQuota(optionTypes, "ketm", 1)
	optionTypes["kondisi-tertentu"] = Filter2OptionsShareQuota(optionTypes, "kondisi-tertentu", 1)

	if reFilterKondisiTertentu == true || reFilterKetm == true {
		return DoFilter(optionTypes)
	}

	return optionTypes
}

func CheckQuota(optionTypes map[string][]*PpdbOption, currentType string, targetType string, reFilter bool) (map[string][]*PpdbOption, bool) {
	fmt.Println("===========================need quota==============================")

	for i := 0; i < len(optionTypes[currentType]); i++ {
		fmt.Println(i, "-", optionTypes[currentType][i].Id, " - ",
			optionTypes[currentType][i].Name,
			" - q:", optionTypes[currentType][i].Quota,
			" - Need:", optionTypes[currentType][i].NeedQuotaFirstOpt,
			" - Added:", optionTypes[currentType][i].AddQuota)

		if i == len(optionTypes[currentType])-1 {
			continue
		}

		if optionTypes[currentType][i].NeedQuotaFirstOpt > 0 {

			sisaQuota := optionTypes[targetType][i].Quota - len(optionTypes[targetType][i].PpdbRegistration)
			fmt.Println("sisaQuota:", sisaQuota)
			if sisaQuota > 0 {

				if optionTypes[currentType][i].NeedQuotaFirstOpt >= sisaQuota {
					optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + sisaQuota
					optionTypes[currentType][i].NeedQuotaFirstOpt = optionTypes[currentType][i].NeedQuotaFirstOpt - sisaQuota
					optionTypes[currentType][i].AddQuota += sisaQuota
					optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - sisaQuota
					fmt.Println("targetType.quota:", optionTypes[targetType][i].Quota)
				} else {
					//jika kebutuhan kuota hanya sedikit artinya hanya membutuhkan sedikit dari sisa kuota lawannya
					fmt.Println("sisaQuota:", sisaQuota, " - NeedQuotaFirstOpt:", optionTypes[currentType][i].NeedQuotaFirstOpt)
					optionTypes[currentType][i].Quota = optionTypes[currentType][i].Quota + optionTypes[currentType][i].NeedQuotaFirstOpt
					optionTypes[currentType][i].AddQuota += optionTypes[currentType][i].NeedQuotaFirstOpt
					optionTypes[targetType][i].Quota = optionTypes[targetType][i].Quota - optionTypes[currentType][i].NeedQuotaFirstOpt
					optionTypes[currentType][i].NeedQuotaFirstOpt = 0

				}

				optionTypes[currentType][i].Filtered = 0
				optionTypes[currentType][i].UpdateQuota = true
				optionTypes[currentType] = PullStudentToFirstChoice(optionTypes[currentType], i)
				reFilter = true
			}
		}
	}

	fmt.Println("optionType1:", currentType)
	for i := 0; i < len(optionTypes[currentType]); i++ {
		fmt.Println(i, "-", optionTypes[currentType][i].Id, " - ", optionTypes[currentType][i].Name,
			" : q: ", optionTypes[currentType][i].Quota,
			" : p: ", len(optionTypes[currentType][i].PpdbRegistration),
			" - needQuota:", optionTypes[currentType][i].NeedQuotaFirstOpt,
			" - AddQuota:", optionTypes[currentType][i].AddQuota,
		)
	}
	fmt.Println("optionType2:", targetType)
	for i := 0; i < len(optionTypes[targetType]); i++ {
		fmt.Println(i, "-", optionTypes[targetType][i].Id, " - ", optionTypes[targetType][i].Name,
			" : q: ", optionTypes[targetType][i].Quota,
			" : p: ", len(optionTypes[targetType][i].PpdbRegistration),
			" - needQuota:", optionTypes[targetType][i].NeedQuotaFirstOpt,
			" - AddQuota:", optionTypes[targetType][i].AddQuota,
		)
	}

	return optionTypes, reFilter
}

func PullStudentToFirstChoice(optionList []*PpdbOption, currTargetIdxOpt int) []*PpdbOption {
	fmt.Println("currTargetIdxOpt:", currTargetIdxOpt)
	//pull student from backup history pilihan 1

	var listPullOpt = make([]int, 0)
	for j := 0; j < len(optionList[currTargetIdxOpt].RegistrationHistory); j++ {
		if optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus != 0 {

			var targetIdxStd int
			var nextTargetIdxOpt int
			if optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex == -1 {
				nextTargetIdxOpt = len(optionList) - 1
				targetIdxStd = FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)

				fmt.Println("Yg tidak diterima == :",
					optionList[currTargetIdxOpt].RegistrationHistory[j].Id,
					" - ", optionList[currTargetIdxOpt].RegistrationHistory[j].Name,
					" - AccStatus:", optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus,
					" - targetIdxOpt:", nextTargetIdxOpt,
					" - TargetIdxStd:", targetIdxStd,
				)
			} else {
				nextTargetIdxOpt = optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex
				targetIdxStd = FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)
				fmt.Println("Yg tidak diterima !=:",
					optionList[currTargetIdxOpt].RegistrationHistory[j].Id,
					" - ", optionList[currTargetIdxOpt].RegistrationHistory[j].Name,
					" - AccStatus:", optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus,
					" - targetIdxOpt:", nextTargetIdxOpt,
					" - TargetIdxStd:", targetIdxStd,
				)
			}

			/*
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedStatus = 0
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedIndex = currTargetIdxOpt
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedChoiceId = optionList[currTargetIdxOpt].Id
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance1
				optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus = 0
				optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex = currTargetIdxOpt
				optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedChoiceId = optionList[currTargetIdxOpt].Id
			*/
			dataChange := StudentUpdate{
				curOptIdx:     currTargetIdxOpt,
				nextOptIdx:    nextTargetIdxOpt,
				nextIdxStd:    targetIdxStd,
				histIdxStd:    j,
				accStatus:     0,
				NextOptChoice: optionList[currTargetIdxOpt].Id,
				Distance:      optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance1,
			}
			UpdatePullStudentFirstChoice(optionList, dataChange)

			optionList[currTargetIdxOpt].AddStd(optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd])
			optionList[nextTargetIdxOpt].RemoveStd(targetIdxStd)
			j--
			if nextTargetIdxOpt != len(optionList)-1 {
				optionList[nextTargetIdxOpt].Filtered = 0
				listPullOpt = append(listPullOpt, nextTargetIdxOpt)
			}

		}
	}

	//pull student from history shifting
	for j := 0; j < len(optionList[currTargetIdxOpt].HistoryShifting); j++ {
		var levelAcc int
		if optionList[currTargetIdxOpt].HistoryShifting[j].FirstChoiceOption.String() ==
			optionList[currTargetIdxOpt].Id.String() {
			levelAcc = 0
		} else if optionList[currTargetIdxOpt].HistoryShifting[j].SecondChoiceOption.String() ==
			optionList[currTargetIdxOpt].Id.String() {
			levelAcc = 1
		} else if optionList[currTargetIdxOpt].HistoryShifting[j].ThirdChoiceOption.String() ==
			optionList[currTargetIdxOpt].Id.String() {
			levelAcc = 2
		} else {
			levelAcc = 3
		}

		//idxStd := models.FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)

		optIdxFirstChoice := FindIndex(optionList[currTargetIdxOpt].HistoryShifting[j].FirstChoiceOption, optionList)
		stdIdx := FindIndexStudent(optionList[currTargetIdxOpt].HistoryShifting[j].Id, optionList[optIdxFirstChoice].RegistrationHistory)
		fmt.Println("shifting name :", optionList[currTargetIdxOpt].HistoryShifting[j].Name)
		fmt.Println("optIdxFirstChoice:", optionList[optIdxFirstChoice].Name, " stdIdx:", stdIdx)

		if optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus > levelAcc {
			fmt.Println("levelAcc:", levelAcc)
			nextTargetIdxOpt := optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex
			fmt.Println("nextTargetIdxOpt.len:", len(optionList[nextTargetIdxOpt].PpdbRegistration))
			targetIdxStd := FindIndexStudentTest(optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].Id, optionList[nextTargetIdxOpt].PpdbRegistration)
			fmt.Println("nextTargetIdxOpt:", nextTargetIdxOpt, " ", optionList[nextTargetIdxOpt].Name, " targetIdxStd:", targetIdxStd)

			/*
				optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = levelAcc
				optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = currTargetIdxOpt
				optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedChoiceId = optionList[currTargetIdxOpt].Id
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedStatus = levelAcc
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedIndex = currTargetIdxOpt
				optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedChoiceId = optionList[currTargetIdxOpt].Id
			*/
			var tmpDistance float64
			if levelAcc == 0 {
				//optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance1
				tmpDistance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance1
			} else if levelAcc == 1 {
				//optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance2
				tmpDistance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance2
			} else if levelAcc == 2 {
				//optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance3
				tmpDistance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance3
			} else {
				//optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance3
				tmpDistance = optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance3
			}

			dataChange := StudentUpdate{
				firstOptIdx:   optIdxFirstChoice,
				accStatus:     levelAcc,
				histIdxStd:    stdIdx,
				curOptIdx:     currTargetIdxOpt,
				nextOptIdx:    nextTargetIdxOpt,
				nextIdxStd:    targetIdxStd,
				NextOptChoice: optionList[currTargetIdxOpt].Id,
				Distance:      tmpDistance,
			}
			UpdatePullStudentByChoice(optionList, dataChange)

			optionList[currTargetIdxOpt].AddStd(optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd])
			optionList[nextTargetIdxOpt].RemoveStd(targetIdxStd)

		}
	}

	for i := 0; i < len(listPullOpt); i++ {
		optionList = PullStudentToFirstChoice(optionList, listPullOpt[i])
	}

	return optionList
}

func Filter2OptionsShareQuota(optionTypes map[string][]*PpdbOption, optType string, loop int) []*PpdbOption {

	runtime.GOMAXPROCS(2)
	var messages = make(chan []*PpdbOption)
	fmt.Println("Filter2OptionsShareQuota")

	var getFiltered = func(objs chan []*PpdbOption, option []*PpdbOption) {
		fmt.Println("bef getFiltered:")
		ppdbOptions := ProcessFilter(option, true, loop)
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

func Filter2OptionsShareQuotaSingleGoroutine(optionTypes map[string][]*PpdbOption, optType string, loop int) []*PpdbOption {

	fmt.Println("Filter2OptionsShareQuota")

	ppdbOptions := ProcessFilter(optionTypes[optType], true, loop)

	fmt.Println("close messages")

	return ppdbOptions
}

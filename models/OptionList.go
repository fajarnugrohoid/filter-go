package models

import (
	"fmt"
)

type PpdbOptionList struct {
	options []PpdbOption
}

func (optionList *PpdbOptionList) AddOpt(item PpdbOption) {
	optionList.options = append(optionList.options, item)
}

func (optionList *PpdbOptionList) UpdateStudent(currTargetIdxOpt int, nextTargetIdxOpt int, targetIdxStd int, j int) {
	optionList.options[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedStatus = 0
	optionList.options[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedIndex = currTargetIdxOpt
	optionList.options[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].AcceptedChoiceId = optionList.options[currTargetIdxOpt].Id
	optionList.options[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance = optionList.options[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance1
	optionList.options[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus = 0
	optionList.options[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex = currTargetIdxOpt
	optionList.options[currTargetIdxOpt].RegistrationHistory[j].AcceptedChoiceId = optionList.options[currTargetIdxOpt].Id
}

func ProcessFilter(optionList []*PpdbOption, status bool, loop int) []*PpdbOption {
	var optIdx int
	var stdIdx int
	for i := 0; i < len(optionList); i++ {
		//fmt.Println("ProcessFilter:", optionList[i].Id)
		if optionList[i].Filtered == 0 {
			SortByDistanceAndAge(optionList[i].PpdbRegistration)

			fmt.Println("afterSortByDistanceAndAge")
			fmt.Println(optionList[i].Id, " - ", optionList[i].Name,
				" len.std:", len(optionList[i].PpdbRegistration),
				" : q: ", optionList[i].Quota, " \n ")
			for y, std := range optionList[i].PpdbRegistration {
				fmt.Println("", y, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
					" AccIdx: ", std.AcceptedIndex,
					" AccId: ", std.AcceptedChoiceId,
				)
			}

			if len(optionList[i].PpdbRegistration) > optionList[i].Quota { //cek jml pendaftar lebih dari quota sekolah

				if optionList[i].UpdateQuota == true {
					optionList[i].NeedQuotaFirstOpt = (len(optionList[i].PpdbRegistration) - optionList[i].Quota)
					optionList[i].UpdateQuota = false
				}
				fmt.Println(optionList[i].Name, " NeedQuotaFirstOpt:", optionList[i].NeedQuotaFirstOpt)

				x := 0
				for j := optionList[i].Quota; j < len(optionList[i].PpdbRegistration); j++ { ////cut siswa yg lebih dari quota move to sec, third choice

					optIdxFirstChoice := FindIndex(optionList[i].PpdbRegistration[j].FirstChoiceOption, optionList)
					stdIdx = FindIndexStudent(optionList[i].PpdbRegistration[j].Id, optionList[optIdxFirstChoice].RegistrationHistory)
					fmt.Println(x, "-findIdxStd:",
						optionList[i].Name, "-",
						optionList[i].PpdbRegistration[j].Id, "-",
						optionList[i].PpdbRegistration[j].Name,
						" - ", optIdxFirstChoice,
						" - ", optionList[optIdxFirstChoice].Name,
						" - ", stdIdx,
						" - AcceptedStatus:", optionList[i].PpdbRegistration[j].AcceptedStatus,
					)

					/*
						for x, std := range optionList[optIdxFirstChoice].RegistrationHistory {
							fmt.Println(">hist3:", x, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
								" AcceptedIndex: ", std.AcceptedIndex)
						}
						fmt.Println("\n") */

					if optionList[i].PpdbRegistration[j].AcceptedStatus == 0 {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 1
						optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[i].PpdbRegistration[j].SecondChoiceOption
						optionList[i].PpdbRegistration[j].Distance = optionList[i].PpdbRegistration[j].Distance2

						optIdx = FindIndex(optionList[i].PpdbRegistration[j].SecondChoiceOption, optionList)

						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 1
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = optIdx
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedChoiceId = optionList[i].PpdbRegistration[j].SecondChoiceOption
						fmt.Println("          >sec ori:", j, ":",
							optionList[i].PpdbRegistration[j].Name, "-",
							optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].Name, "-",
							optionList[i].PpdbRegistration[j].SecondChoiceOption, " - ",
							optIdx)
					} else if optionList[i].PpdbRegistration[j].AcceptedStatus == 1 {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 2
						optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[i].PpdbRegistration[j].ThirdChoiceOption
						optionList[i].PpdbRegistration[j].Distance = optionList[i].PpdbRegistration[j].Distance3

						optIdx = FindIndex(optionList[i].PpdbRegistration[j].ThirdChoiceOption, optionList)

						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 2
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = optIdx
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedChoiceId = optionList[i].PpdbRegistration[j].ThirdChoiceOption
						fmt.Println("          >third ori:", j, ":", optionList[i].PpdbRegistration[j].Name, "-", optionList[i].PpdbRegistration[j].SecondChoiceOption, " - ", optIdx)
					} else {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 3

						optIdx = len(optionList) - 1
						optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[optIdx].Id
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 3
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = optIdx
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedChoiceId = optionList[optIdx].Id
					}

					if optIdx == -1 || optIdx == len(optionList)-1 { //jika tidak ada option dan telah dilempar ke pembuangan
						fmt.Println("in if -1 idx:", optIdx, "-", optionList[i].PpdbRegistration[j].Name, "-", len(optionList)-1)
						optionList[i].PpdbRegistration[j].AcceptedStatus = 3
						optionList[i].PpdbRegistration[j].AcceptedIndex = len(optionList) - 1
						optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[len(optionList)-1].Id
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 3
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = len(optionList) - 1
						optionList[len(optionList)-1].AddStd(optionList[i].PpdbRegistration[j])
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedChoiceId = optionList[len(optionList)-1].Id
						optionList[i].RemoveStd(j)
						j--
					} else {
						fmt.Println(optionList[i].PpdbRegistration[j].Name, "-idx:", optIdx)
						optionList[optIdx].AddStd(optionList[i].PpdbRegistration[j])

						optionList[i].AddHistory(optionList[i].PpdbRegistration[j], optIdx)

						optionList[i].RemoveStd(j)
						j--

						optionList[optIdx].Filtered = 0
						status = true

					}
					x++
				}

			}
			optionList[i].Filtered = 1
		}
	}

	if status == true {
		return ProcessFilter(optionList, false, 1)
	}
	return optionList
}

package models

import (
	"fmt"
	"sort"
)

type PpdbOptionList struct {
	options []PpdbOption
}

func (optionList *PpdbOptionList) AddOpt(item PpdbOption) {
	optionList.options = append(optionList.options, item)
}

func ProcessFilter(optionList []*PpdbOption, status bool) []*PpdbOption {
	var optIdx int
	var stdIdx int
	for i := 0; i < len(optionList); i++ {
		fmt.Println("ProcessFilter:", optionList[i].Id)
		if optionList[i].Filtered == 0 {
			sort.Sort(ByDistance(optionList[i].PpdbRegistration))
			fmt.Println(optionList[i].Id, " - ", optionList[i].Name,
				" len.std:", len(optionList[i].PpdbRegistration),
				" : q: ", optionList[i].Quota, " \n ")

			if len(optionList[i].PpdbRegistration) > optionList[i].Quota { //cek jml pendaftar lebih dari quota sekolah

				for j := optionList[i].Quota; j < len(optionList[i].PpdbRegistration); j++ { ////cut siswa yg lebih dari quota move to sec, third choice
					//idx = -1

					optIdxFirstChoice := FindIndex(optionList[i].PpdbRegistration[j].FirstChoiceOption, optionList)
					stdIdx = FindIndexStudent(optionList[i].PpdbRegistration[j].Id, optionList[optIdxFirstChoice].RegistrationHistory)
					fmt.Println("findIdxStd:",
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

						optIdx = FindIndex(optionList[i].PpdbRegistration[j].SecondChoiceOption, optionList)

						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 1
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = optIdx
						fmt.Println(">sec ori:", j, ":",
							optionList[i].PpdbRegistration[j].Name, "-",
							optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].Name, "-",
							optionList[i].PpdbRegistration[j].SecondChoiceOption, " - ",
							optIdx)
					} else if optionList[i].PpdbRegistration[j].AcceptedStatus == 1 {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 2
						optIdx = FindIndex(optionList[i].PpdbRegistration[j].ThirdChoiceOption, optionList)

						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 2
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = optIdx
						fmt.Println(">third ori:", j, ":", optionList[i].PpdbRegistration[j].Name, "-", optionList[i].PpdbRegistration[j].SecondChoiceOption, " - ", optIdx)
					}

					fmt.Println("idx:", optIdx, "-", optionList[i].PpdbRegistration[j].Name)
					if optIdx == -1 || optIdx == len(optionList)-1 { //jika tidak ada option dan telah dilempar ke pembuangan
						fmt.Println("in if -1 idx:", optIdx, "-", optionList[i].PpdbRegistration[j].Name, "-", len(optionList)-1)
						optionList[i].PpdbRegistration[j].AcceptedStatus = 3
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus = 3
						optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex = len(optionList) - 1
						optionList[len(optionList)-1].AddStd(optionList[i].PpdbRegistration[j])
						optionList[i].RemoveStd(j)
						j--
					} else {
						fmt.Println(optionList[i].PpdbRegistration[j].Name, "-idx:", optIdx)
						optionList[optIdx].AddStd(optionList[i].PpdbRegistration[j])
						optionList[i].RemoveStd(j)
						j--

						optionList[optIdx].Filtered = 0
						status = true

					}

				}
				optionList[i].IsNeedQuota = true
			}
			optionList[i].Filtered = 1
		}
	}

	if status == true {
		return ProcessFilter(optionList, false)
	}
	return optionList
}

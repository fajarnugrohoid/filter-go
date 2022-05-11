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

func ProcessFilter(optionList []*PpdbOption, status bool, loop int) []*PpdbOption {
	var nextOptIdx int
	var histIdxStd int
	for curOptIdx := 0; curOptIdx < len(optionList); curOptIdx++ {
		//fmt.Println("ProcessFilter:", optionList[i].Id)
		if optionList[curOptIdx].Filtered == 0 {
			SortByDistanceAndAge(optionList[curOptIdx].PpdbRegistration)

			fmt.Println("afterSortByDistanceAndAge")
			fmt.Println(optionList[curOptIdx].Id, " - ", optionList[curOptIdx].Name,
				" len.std:", len(optionList[curOptIdx].PpdbRegistration),
				" : q: ", optionList[curOptIdx].Quota, " \n ")
			for y, std := range optionList[curOptIdx].PpdbRegistration {
				fmt.Println("", y, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
					" AccIdx: ", std.AcceptedIndex,
					" AccId: ", std.AcceptedChoiceId,
				)
			}

			if len(optionList[curOptIdx].PpdbRegistration) > optionList[curOptIdx].Quota { //cek jml pendaftar lebih dari quota sekolah

				if optionList[curOptIdx].UpdateQuota == true {
					optionList[curOptIdx].NeedQuotaFirstOpt = (len(optionList[curOptIdx].PpdbRegistration) - optionList[curOptIdx].Quota)
					optionList[curOptIdx].UpdateQuota = false
				}
				fmt.Println(optionList[curOptIdx].Name, " NeedQuotaFirstOpt:", optionList[curOptIdx].NeedQuotaFirstOpt)

				x := 0
				for curIdxStd := optionList[curOptIdx].Quota; curIdxStd < len(optionList[curOptIdx].PpdbRegistration); curIdxStd++ { ////cut siswa yg lebih dari quota move to sec, third choice

					firstOptIdx := FindIndex(optionList[curOptIdx].PpdbRegistration[curIdxStd].FirstChoiceOption, optionList)
					histIdxStd = FindIndexStudent(optionList[curOptIdx].PpdbRegistration[curIdxStd].Id, optionList[firstOptIdx].RegistrationHistory)
					fmt.Println(x, "-findIdxStd:",
						optionList[curOptIdx].Name, "-",
						optionList[curOptIdx].PpdbRegistration[curIdxStd].Id, "-",
						optionList[curOptIdx].PpdbRegistration[curIdxStd].Name,
						" - ", firstOptIdx,
						" - ", optionList[firstOptIdx].Name,
						" - ", histIdxStd,
						" - AcceptedStatus:", optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus,
					)

					if optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus == 0 {

						nextOptIdx = FindIndex(optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption, optionList)

						/*
							optionList[i].PpdbRegistration[j].AcceptedStatus = 1
							optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[i].PpdbRegistration[j].SecondChoiceOption
							optionList[i].PpdbRegistration[j].Distance = optionList[i].PpdbRegistration[j].Distance2
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedStatus = 1
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedIndex = optIdx
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[i].PpdbRegistration[j].SecondChoiceOption
						*/
						//UpdateMoveStudent(optionList, curOptIdx, nextOptIdx, firstOptIdx, j, histIdxStd, 1)
						dataChange := &StudentUpdate{
							curOptIdx:     curOptIdx,
							nextOptIdx:    nextOptIdx,
							firstOptIdx:   firstOptIdx,
							curIdxStd:     curIdxStd,
							histIdxStd:    histIdxStd,
							accStatus:     1,
							NextOptChoice: optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance2,
						}
						dataChange.UpdateMoveStudent(optionList)
						fmt.Println("          >sec ori:", curIdxStd, ":",
							optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-",
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].Name, "-",
							optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption, " - ",
							nextOptIdx)

					} else if optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus == 1 {

						nextOptIdx = FindIndex(optionList[curOptIdx].PpdbRegistration[curIdxStd].ThirdChoiceOption, optionList)
						/*
							optionList[i].PpdbRegistration[j].AcceptedStatus = 2
							optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[i].PpdbRegistration[j].ThirdChoiceOption
							optionList[i].PpdbRegistration[j].Distance = optionList[i].PpdbRegistration[j].Distance3

							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedStatus = 2
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedIndex = nextOptIdx
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[i].PpdbRegistration[j].ThirdChoiceOption
						*/
						// curOptIdx, nextOptIdx, firstOptIdx, j, histIdxStd, 2
						dataChange := &StudentUpdate{
							curOptIdx:     curOptIdx,
							nextOptIdx:    nextOptIdx,
							firstOptIdx:   firstOptIdx,
							curIdxStd:     curIdxStd,
							histIdxStd:    histIdxStd,
							accStatus:     2,
							NextOptChoice: optionList[curOptIdx].PpdbRegistration[curIdxStd].ThirdChoiceOption,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance3,
						}
						dataChange.UpdateMoveStudent(optionList)
						fmt.Println("          >third ori:", curIdxStd, ":", optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-", optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption, " - ", nextOptIdx)
					} else {
						nextOptIdx = len(optionList) - 1
						/*
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus = 3
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedChoiceId = optionList[nextOptIdx].Id
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedStatus = 3
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedIndex = nextOptIdx
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[nextOptIdx].Id
						*/
						dataChange := &StudentUpdate{
							curOptIdx:     curOptIdx,
							nextOptIdx:    nextOptIdx,
							firstOptIdx:   firstOptIdx,
							curIdxStd:     curIdxStd,
							histIdxStd:    histIdxStd,
							accStatus:     3,
							NextOptChoice: optionList[nextOptIdx].Id,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance3,
						}
						dataChange.UpdateMoveStudent(optionList)
					}

					if nextOptIdx == -1 || nextOptIdx == len(optionList)-1 { //jika tidak ada option dan telah dilempar ke pembuangan
						fmt.Println("in if -1 idx:", nextOptIdx, "-", optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-", len(optionList)-1)
						nextOptIdx = len(optionList) - 1
						/*
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus = 3
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedIndex = nextOptIdx
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedChoiceId = optionList[nextOptIdx].Id
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedStatus = 3
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedIndex = nextOptIdx
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[len(optionList)-1].Id
						*/
						dataChange := &StudentUpdate{
							curOptIdx:     curOptIdx,
							nextOptIdx:    nextOptIdx,
							firstOptIdx:   firstOptIdx,
							curIdxStd:     curIdxStd,
							histIdxStd:    histIdxStd,
							accStatus:     3,
							NextOptChoice: optionList[nextOptIdx].Id,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance3,
						}
						dataChange.UpdateMoveStudent(optionList)

						optionList[len(optionList)-1].AddStd(optionList[curOptIdx].PpdbRegistration[curIdxStd])
						optionList[curOptIdx].RemoveStd(curIdxStd)
						curIdxStd--
					} else {
						fmt.Println(optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-idx:", nextOptIdx)
						optionList[nextOptIdx].AddStd(optionList[curOptIdx].PpdbRegistration[curIdxStd])
						optionList[curOptIdx].AddHistory(optionList[curOptIdx].PpdbRegistration[curIdxStd], nextOptIdx)
						optionList[curOptIdx].RemoveStd(curIdxStd)
						curIdxStd--

						optionList[nextOptIdx].Filtered = 0
						status = true

					}
					x++
				}

			}
			optionList[curOptIdx].Filtered = 1
		}
	}

	if status == true {
		return ProcessFilter(optionList, false, 1)
	}
	return optionList
}

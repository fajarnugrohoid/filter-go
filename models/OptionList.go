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

func ProcessFilter(optionList []*PpdbOption, status bool) []*PpdbOption {

	for i := 0; i < len(optionList); i++ {
		if optionList[i].Filtered == 0 {
			//sort.Sort(ByScore(optionList.options[i].PpdbRegistration))
			fmt.Println(optionList[i].Id, " - ", optionList[i].Name,
				" len.std:", len(optionList[i].PpdbRegistration),
				" : q: ", optionList[i].Quota, " \n ")

			if len(optionList[i].PpdbRegistration) > optionList[i].Quota {

				for j := optionList[i].Quota; j < len(optionList[i].PpdbRegistration); j++ {
					idx := -1
					if optionList[i].PpdbRegistration[j].AcceptedStatus == 0 {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 1
						idx := FindIndex(optionList[i].PpdbRegistration[j].SecondChoiceOption, optionList)
						fmt.Println(">ori:", j, ":", optionList[i].PpdbRegistration[j].Name, "-", optionList[i].PpdbRegistration[j].SecondChoiceOption, " - ", idx)
					} else if optionList[i].PpdbRegistration[j].AcceptedStatus == 1 {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 2
						idx := FindIndex(optionList[i].PpdbRegistration[j].ThirdChoiceOption, optionList)
						fmt.Println(">ori:", j, ":", optionList[i].PpdbRegistration[j].Name, "-", optionList[i].PpdbRegistration[j].SecondChoiceOption, " - ", idx)
					}
					if idx == -1 {
						optionList[i].PpdbRegistration[j].AcceptedStatus = 3
						optionList[len(optionList)-1].AddStd(optionList[i].PpdbRegistration[j])
						optionList[i].RemoveStd(j)
						j--
					} else {

						optionList[idx].AddStd(optionList[i].PpdbRegistration[j])
						optionList[i].RemoveStd(j)
						j--
						optionList[idx].Filtered = 0
						status = true
					}
				}

			}
			optionList[i].Filtered = 1
		}
	}

	if status == true {
		return ProcessFilter(optionList, false)
	}
	return optionList
}

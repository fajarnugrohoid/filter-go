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

func Filter2OptionsShareQuota(optionTypes map[string][]*models.PpdbOption) []*models.PpdbOption {

	runtime.GOMAXPROCS(2)
	var messages = make(chan []*models.PpdbOption)
	fmt.Println("Filter2OptionsShareQuota")
	/*
		var getFiltered = func(option []models.PpdbOption) {
			ppdbOptions := models.ProcessFilter(option, false)
			messages <- ppdbOptions
		}*/

	var getFiltered = func(objs chan []*models.PpdbOption, option []*models.PpdbOption) {
		fmt.Println("bef getFiltered:")
		ppdbOptions := models.ProcessFilter(option, true)
		fmt.Println("aft getFiltered")
		messages <- ppdbOptions
	}
	fmt.Println("end getFiltered")
	go getFiltered(messages, optionTypes["ketm"])
	fmt.Println("go done getFiltered")
	data := <-messages // read from channel a

	close(messages)

	fmt.Println("close messages")

	//fmt.Println(data)

	return data
}

package logic

import (
	"filterisasi/controller"
	"filterisasi/models/domain"
	"github.com/sirupsen/logrus"
)

func PullStudentToFirstChoice(optionList []*domain.PpdbOption, currTargetIdxOpt int, logger *logrus.Logger) []*domain.PpdbOption {
	logger.Debug("currTargetIdxOpt:", currTargetIdxOpt)
	//pull student from backup history pilihan 1

	var listPullOpt = make([]int, 0)
	for j := 0; j < len(optionList[currTargetIdxOpt].RegistrationHistory); j++ {
		if optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus != 0 {

			var targetIdxStd int
			var nextTargetIdxOpt int
			if optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex == -1 {
				nextTargetIdxOpt = len(optionList) - 1
				targetIdxStd = domain.FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)

				logger.Debug("Yg tidak diterima == :",
					optionList[currTargetIdxOpt].RegistrationHistory[j].Id,
					" - ", optionList[currTargetIdxOpt].RegistrationHistory[j].Name,
					" - AccStatus:", optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedStatus,
					" - targetIdxOpt:", nextTargetIdxOpt,
					" - TargetIdxStd:", targetIdxStd,
				)
			} else {
				nextTargetIdxOpt = optionList[currTargetIdxOpt].RegistrationHistory[j].AcceptedIndex
				targetIdxStd = domain.FindIndexStudent(optionList[currTargetIdxOpt].RegistrationHistory[j].Id, optionList[nextTargetIdxOpt].PpdbRegistration)
				logger.Debug("Yg tidak diterima !=:",
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
			dataChange := controller.StudentUpdate{
				CurOptIdx:     currTargetIdxOpt,
				NextOptIdx:    nextTargetIdxOpt,
				NextIdxStd:    targetIdxStd,
				HistIdxStd:    j,
				AccStatus:     0,
				NextOptChoice: optionList[currTargetIdxOpt].Id,
				Distance:      optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd].Distance1,
			}
			controller.UpdatePullStudentFirstChoice(optionList, dataChange)

			optionList[currTargetIdxOpt].AddStd(optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd], logger)
			optionList[nextTargetIdxOpt].RemoveStd(targetIdxStd, logger)
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

		optIdxFirstChoice := domain.FindIndex(optionList[currTargetIdxOpt].HistoryShifting[j].FirstChoiceOption, optionList)
		stdIdx := domain.FindIndexStudent(optionList[currTargetIdxOpt].HistoryShifting[j].Id, optionList[optIdxFirstChoice].RegistrationHistory)
		logger.Debug("shifting name :", optionList[currTargetIdxOpt].HistoryShifting[j].Name)
		logger.Debug("optIdxFirstChoice:", optionList[optIdxFirstChoice].Name, " stdIdx:", stdIdx)

		if optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedStatus > levelAcc {
			logger.Debug("levelAcc:", levelAcc)
			nextTargetIdxOpt := optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].AcceptedIndex
			logger.Debug("nextTargetIdxOpt.len:", len(optionList[nextTargetIdxOpt].PpdbRegistration))
			targetIdxStd := domain.FindIndexStudentTest(optionList[optIdxFirstChoice].RegistrationHistory[stdIdx].Id, optionList[nextTargetIdxOpt].PpdbRegistration)
			logger.Debug("nextTargetIdxOpt:", nextTargetIdxOpt, " ", optionList[nextTargetIdxOpt].Name, " targetIdxStd:", targetIdxStd)

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

			dataChange := controller.StudentUpdate{
				FirstOptIdx:   optIdxFirstChoice,
				AccStatus:     levelAcc,
				HistIdxStd:    stdIdx,
				CurOptIdx:     currTargetIdxOpt,
				NextOptIdx:    nextTargetIdxOpt,
				NextIdxStd:    targetIdxStd,
				NextOptChoice: optionList[currTargetIdxOpt].Id,
				Distance:      tmpDistance,
			}
			controller.UpdatePullStudentByChoice(optionList, dataChange)

			optionList[currTargetIdxOpt].AddStd(optionList[nextTargetIdxOpt].PpdbRegistration[targetIdxStd], logger)
			optionList[nextTargetIdxOpt].RemoveStd(targetIdxStd, logger)

		}
	}

	for i := 0; i < len(listPullOpt); i++ {
		optionList = PullStudentToFirstChoice(optionList, listPullOpt[i], logger)
	}

	return optionList
}

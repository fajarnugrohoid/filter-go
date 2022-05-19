package logic

import (
	"filterisasi/models"
	"filterisasi/repositories"
	"github.com/sirupsen/logrus"
)

func ProcessFilter(optionList []*models.PpdbOption, status bool, loop int, logger *logrus.Logger) []*models.PpdbOption {
	var nextOptIdx int
	var histIdxStd int

	for curOptIdx := 0; curOptIdx < len(optionList); curOptIdx++ {
		//fmt.Println("ProcessFilter:", optionList[i].Id)
		if optionList[curOptIdx].Filtered == 0 {
			models.SortByDistanceAndAge(optionList[curOptIdx].PpdbRegistration)

			logger.Debug("afterSortByDistanceAndAge")
			logger.Debug(optionList[curOptIdx].Id, " - ", optionList[curOptIdx].Name,
				" len.std:", len(optionList[curOptIdx].PpdbRegistration),
				" : q: ", optionList[curOptIdx].Quota, " \n ")
			for y, std := range optionList[curOptIdx].PpdbRegistration {
				logger.Debug("", y, ":", std.Name, " - acc:", std.AcceptedStatus, " distance1: ", std.Distance1,
					" AccIdx: ", std.AcceptedIndex,
					" AccId: ", std.AcceptedChoiceId,
				)
			}

			if len(optionList[curOptIdx].PpdbRegistration) > optionList[curOptIdx].Quota { //cek jml pendaftar lebih dari quota sekolah

				if optionList[curOptIdx].UpdateQuota == true {
					optionList[curOptIdx].NeedQuota = len(optionList[curOptIdx].PpdbRegistration) - optionList[curOptIdx].Quota
					optionList[curOptIdx].UpdateQuota = false
				}
				logger.Debug(optionList[curOptIdx].Name, " NeedQuota:", optionList[curOptIdx].NeedQuota)

				x := 0
				for curIdxStd := optionList[curOptIdx].Quota; curIdxStd < len(optionList[curOptIdx].PpdbRegistration); curIdxStd++ { ////cut siswa yg lebih dari quota move to sec, third choice

					firstOptIdx := models.FindIndex(optionList[curOptIdx].PpdbRegistration[curIdxStd].FirstChoiceOption, optionList)
					histIdxStd = models.FindIndexStudent(optionList[curOptIdx].PpdbRegistration[curIdxStd].Id, optionList[firstOptIdx].RegistrationHistory)
					logger.Debug(x, "-findIdxStd:",
						optionList[curOptIdx].Name, "-",
						optionList[curOptIdx].PpdbRegistration[curIdxStd].Id, "-",
						optionList[curOptIdx].PpdbRegistration[curIdxStd].Name,
						" - ", firstOptIdx,
						" - ", optionList[firstOptIdx].Name,
						" - ", histIdxStd,
						" - AcceptedStatus:", optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus,
					)

					if optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus == 0 {

						nextOptIdx = models.FindIndex(optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption, optionList)

						/*
							optionList[i].PpdbRegistration[j].AcceptedStatus = 1
							optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[i].PpdbRegistration[j].SecondChoiceOption
							optionList[i].PpdbRegistration[j].Distance = optionList[i].PpdbRegistration[j].Distance2
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedStatus = 1
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedIndex = optIdx
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[i].PpdbRegistration[j].SecondChoiceOption
						*/
						//UpdateMoveStudent(optionList, curOptIdx, nextOptIdx, firstOptIdx, j, histIdxStd, 1)
						dataChange := repositories.StudentUpdate{
							CurOptIdx:     curOptIdx,
							NextOptIdx:    nextOptIdx,
							FirstOptIdx:   firstOptIdx,
							CurIdxStd:     curIdxStd,
							HistIdxStd:    histIdxStd,
							AccStatus:     1,
							NextOptChoice: optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance2,
						}
						repositories.UpdateMoveStudent(optionList, dataChange)
						logger.Debug("          >sec ori:", curIdxStd, ":",
							optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-",
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].Name, "-",
							optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption, " - ",
							nextOptIdx)

					} else if optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus == 1 {

						nextOptIdx = models.FindIndex(optionList[curOptIdx].PpdbRegistration[curIdxStd].ThirdChoiceOption, optionList)
						/*
							optionList[i].PpdbRegistration[j].AcceptedStatus = 2
							optionList[i].PpdbRegistration[j].AcceptedChoiceId = optionList[i].PpdbRegistration[j].ThirdChoiceOption
							optionList[i].PpdbRegistration[j].Distance = optionList[i].PpdbRegistration[j].Distance3

							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedStatus = 2
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedIndex = nextOptIdx
							optionList[optIdxFirstChoice].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[i].PpdbRegistration[j].ThirdChoiceOption
						*/
						// curOptIdx, nextOptIdx, firstOptIdx, j, histIdxStd, 2
						dataChange := repositories.StudentUpdate{
							CurOptIdx:     curOptIdx,
							NextOptIdx:    nextOptIdx,
							FirstOptIdx:   firstOptIdx,
							CurIdxStd:     curIdxStd,
							HistIdxStd:    histIdxStd,
							AccStatus:     2,
							NextOptChoice: optionList[curOptIdx].PpdbRegistration[curIdxStd].ThirdChoiceOption,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance3,
						}
						repositories.UpdateMoveStudent(optionList, dataChange)
						logger.Debug("          >third ori:", curIdxStd, ":", optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-", optionList[curOptIdx].PpdbRegistration[curIdxStd].SecondChoiceOption, " - ", nextOptIdx)
					} else {
						nextOptIdx = len(optionList) - 1
						/*
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus = 3
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedChoiceId = optionList[nextOptIdx].Id
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedStatus = 3
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedIndex = nextOptIdx
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[nextOptIdx].Id
						*/
						dataChange := repositories.StudentUpdate{
							CurOptIdx:     curOptIdx,
							NextOptIdx:    nextOptIdx,
							FirstOptIdx:   firstOptIdx,
							CurIdxStd:     curIdxStd,
							HistIdxStd:    histIdxStd,
							AccStatus:     3,
							NextOptChoice: optionList[nextOptIdx].Id,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance3,
						}
						repositories.UpdateMoveStudent(optionList, dataChange)
					}

					if nextOptIdx == -1 || nextOptIdx == len(optionList)-1 { //jika tidak ada option dan telah dilempar ke pembuangan

						nextOptIdx = len(optionList) - 1
						logger.Debug("in if -1 idx:", nextOptIdx, "-",
							optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-", len(optionList)-1, " accId:", optionList[nextOptIdx].Id)
						/*
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedStatus = 3
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedIndex = nextOptIdx
							optionList[curOptIdx].PpdbRegistration[curIdxStd].AcceptedChoiceId = optionList[nextOptIdx].Id
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedStatus = 3
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedIndex = nextOptIdx
							optionList[firstOptIdx].RegistrationHistory[histIdxStd].AcceptedChoiceId = optionList[len(optionList)-1].Id
						*/
						dataChange := repositories.StudentUpdate{
							CurOptIdx:     curOptIdx,
							NextOptIdx:    nextOptIdx,
							FirstOptIdx:   firstOptIdx,
							CurIdxStd:     curIdxStd,
							HistIdxStd:    histIdxStd,
							AccStatus:     3,
							NextOptChoice: optionList[nextOptIdx].Id,
							Distance:      optionList[curOptIdx].PpdbRegistration[curIdxStd].Distance3,
						}
						repositories.UpdateMoveStudent(optionList, dataChange)

						optionList[len(optionList)-1].AddStd(optionList[curOptIdx].PpdbRegistration[curIdxStd], logger)
						optionList[curOptIdx].RemoveStd(curIdxStd, logger)
						for n := 0; n < len(optionList[len(optionList)-1].PpdbRegistration); n++ {
							logger.Debug(optionList[len(optionList)-1].PpdbRegistration[n].Name, " accId:", optionList[len(optionList)-1].PpdbRegistration[n].AcceptedChoiceId)
						}
						curIdxStd--
					} else {
						logger.Debug(optionList[curOptIdx].PpdbRegistration[curIdxStd].Name, "-idx:", nextOptIdx)
						optionList[nextOptIdx].AddStd(optionList[curOptIdx].PpdbRegistration[curIdxStd], logger)
						optionList[curOptIdx].AddHistory(optionList[curOptIdx].PpdbRegistration[curIdxStd], nextOptIdx)
						optionList[curOptIdx].RemoveStd(curIdxStd, logger)
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
		return ProcessFilter(optionList, false, 1, logger)
	}
	return optionList
}

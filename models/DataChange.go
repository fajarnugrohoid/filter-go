package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentUpdate struct {
	curOptIdx     int
	nextOptIdx    int
	firstOptIdx   int
	curIdxStd     int
	nextIdxStd    int
	histIdxStd    int
	accStatus     int
	NextOptChoice primitive.ObjectID
	Distance      float64
}

func UpdateMoveStudent(optionList []*PpdbOption, d StudentUpdate) {
	optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].AcceptedStatus = d.accStatus
	optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].SecondChoiceOption
	optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].Distance = d.Distance
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedStatus = d.accStatus
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedIndex = d.nextOptIdx
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedChoiceId = d.NextOptChoice
}

//tarik siswa semua siswa yg terlempar ke pilihan pertama
func UpdatePullStudentFirstChoice(optionList []*PpdbOption, d StudentUpdate) {
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].AcceptedStatus = d.accStatus
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].AcceptedIndex = d.curOptIdx
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].Id
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].Distance = d.Distance
	optionList[d.curOptIdx].RegistrationHistory[d.histIdxStd].AcceptedStatus = d.accStatus
	optionList[d.curOptIdx].RegistrationHistory[d.histIdxStd].AcceptedIndex = d.curOptIdx
	optionList[d.curOptIdx].RegistrationHistory[d.histIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].Id
	optionList[d.curOptIdx].RegistrationHistory[d.histIdxStd].Distance = d.Distance
}

//tarik siswa ke pilihan sebelumnya yg terlempar ke pilihan yg lebih tinggi (selanjutnya)
func UpdatePullStudentByChoice(optionList []*PpdbOption, d StudentUpdate) {
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].AcceptedStatus = d.accStatus
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].AcceptedIndex = d.curOptIdx
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].Id
	optionList[d.nextOptIdx].PpdbRegistration[d.nextIdxStd].Distance = d.Distance

	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedStatus = d.accStatus
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedIndex = d.curOptIdx
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].Id
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].Distance = d.Distance

}

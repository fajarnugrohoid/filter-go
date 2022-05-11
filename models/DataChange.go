package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentUpdate struct {
	curOptIdx     int
	nextOptIdx    int
	firstOptIdx   int
	curIdxStd     int
	histIdxStd    int
	accStatus     int
	NextOptChoice primitive.ObjectID
	Distance      float64
}

func (d *StudentUpdate) UpdateMoveStudent(optionList []*PpdbOption) {
	optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].AcceptedStatus = d.accStatus
	optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].SecondChoiceOption
	optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].Distance = d.Distance
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedStatus = d.accStatus
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedIndex = d.nextOptIdx
	optionList[d.firstOptIdx].RegistrationHistory[d.histIdxStd].AcceptedChoiceId = d.NextOptChoice
}

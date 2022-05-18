package repositories

import (
	"filterisasi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentUpdate struct {
	CurOptIdx     int
	NextOptIdx    int
	FirstOptIdx   int
	CurIdxStd     int
	NextIdxStd    int
	HistIdxStd    int
	AccStatus     int
	NextOptChoice primitive.ObjectID
	Distance      float64
}

func UpdateMoveStudent(optionList []*models.PpdbOption, d StudentUpdate) {
	optionList[d.CurOptIdx].PpdbRegistration[d.CurIdxStd].AcceptedStatus = d.AccStatus
	optionList[d.CurOptIdx].PpdbRegistration[d.CurIdxStd].AcceptedIndex = d.NextOptIdx
	//optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].AcceptedChoiceId = optionList[d.curOptIdx].PpdbRegistration[d.curIdxStd].SecondChoiceOption
	optionList[d.CurOptIdx].PpdbRegistration[d.CurIdxStd].AcceptedChoiceId = d.NextOptChoice
	optionList[d.CurOptIdx].PpdbRegistration[d.CurIdxStd].Distance = d.Distance

	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedStatus = d.AccStatus
	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedIndex = d.NextOptIdx
	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedChoiceId = d.NextOptChoice
}

//tarik siswa semua siswa yg terlempar ke pilihan pertama
func UpdatePullStudentFirstChoice(optionList []*models.PpdbOption, d StudentUpdate) {
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].AcceptedStatus = d.AccStatus
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].AcceptedIndex = d.CurOptIdx
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].AcceptedChoiceId = optionList[d.CurOptIdx].Id
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].Distance = d.Distance
	optionList[d.CurOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedStatus = d.AccStatus
	optionList[d.CurOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedIndex = d.CurOptIdx
	optionList[d.CurOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedChoiceId = optionList[d.CurOptIdx].Id
	optionList[d.CurOptIdx].RegistrationHistory[d.HistIdxStd].Distance = d.Distance
}

//tarik siswa ke pilihan sebelumnya yg terlempar ke pilihan yg lebih tinggi (selanjutnya)
func UpdatePullStudentByChoice(optionList []*models.PpdbOption, d StudentUpdate) {
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].AcceptedStatus = d.AccStatus
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].AcceptedIndex = d.CurOptIdx
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].AcceptedChoiceId = optionList[d.CurOptIdx].Id
	optionList[d.NextOptIdx].PpdbRegistration[d.NextIdxStd].Distance = d.Distance

	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedStatus = d.AccStatus
	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedIndex = d.CurOptIdx
	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].AcceptedChoiceId = optionList[d.CurOptIdx].Id
	optionList[d.FirstOptIdx].RegistrationHistory[d.HistIdxStd].Distance = d.Distance

}

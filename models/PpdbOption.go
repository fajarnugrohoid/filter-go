package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PpdbOption struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty"`
	Name                string             `bson:"name,omitempty"`
	Type                string             `bson:"type,omitempty"`
	Quota               int                `bson:"quota,omitempty"`
	QuotaOld            int                `bson:"quota_old,omitempty"`
	TotalQuota          int                `bson:"total_quota,omitempty"`
	SchoolId            primitive.ObjectID `bson:"school,omitempty"`
	Filtered            int
	UpdateQuota         bool
	NeedQuota           int
	AddQuota            int
	PpdbSchool          PpdbSchool `bson:"ppdb_schools,omitempty"`
	PpdbRegistration    []PpdbRegistration
	RegistrationHistory []PpdbRegistration
	HistoryShifting     []PpdbRegistration
}

func (ppdbOption PpdbOption) addItem(options []PpdbOption) []PpdbOption {
	return append(options, ppdbOption)
}

func (option *PpdbOption) AddStd(item PpdbRegistration, logger *logrus.Logger) {
	logger.Debug("addStd:", item.Name)
	option.PpdbRegistration = append(option.PpdbRegistration, item)
}
func (option *PpdbOption) RemoveStd(i int, logger *logrus.Logger) {
	logger.Debug("RemoveStd:", option.PpdbRegistration[i].Name)
	option.PpdbRegistration = append(option.PpdbRegistration[:i], option.PpdbRegistration[i+1:]...)
}

func (option *PpdbOption) AddHistory(item PpdbRegistration, accIndex int) {
	item.AcceptedIndex = accIndex

	var isExist = false
	var stdIdx int
	for i := 0; i < len(option.HistoryShifting); i++ {
		if option.HistoryShifting[i].Id.String() == item.Id.String() {
			isExist = true
			stdIdx = i
			break
		}
	}
	if isExist == true {
		option.HistoryShifting[stdIdx].AcceptedIndex = accIndex
	} else {
		option.HistoryShifting = append(option.HistoryShifting, item)
	}

}
func (option *PpdbOption) RemoveHistory(i int) {
	option.HistoryShifting = append(option.HistoryShifting[:i], option.HistoryShifting[i+1:]...)
}

func FindIndex(element primitive.ObjectID, data []*PpdbOption) int {
	for k, v := range data {
		//fmt.Println("element:", element.String(), "==", v.Id.String())
		if element.String() == v.Id.String() {
			return k
		}
	}
	return -1 //not found.
}

func FindIndexStudent(element primitive.ObjectID, data []PpdbRegistration) int {

	for k, v := range data {
		//fmt.Println("element:", element.String(), "==", v.Id.String(), " - ", v.Name)
		if element.String() == v.Id.String() {
			return k
		}
	}
	return -1 //not found.
}
func FindIndexStudentTest(element primitive.ObjectID, data []PpdbRegistration) int {

	for k, v := range data {
		fmt.Println("element:", element.String(), "==", v.Id.String(), " - ", v.Name)
		if element.String() == v.Id.String() {
			return k
		}
	}
	return -1 //not found.
}

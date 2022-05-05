package models

import (
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
	IsNeedQuota         bool
	PpdbSchool          PpdbSchool `bson:"ppdb_schools,omitempty"`
	PpdbRegistration    []PpdbRegistration
	RegistrationHistory []PpdbRegistration
}

func (ppdbOption PpdbOption) addItem(options []PpdbOption) []PpdbOption {
	return append(options, ppdbOption)
}

func (option *PpdbOption) AddStd(item PpdbRegistration) {
	option.PpdbRegistration = append(option.PpdbRegistration, item)
}
func (option *PpdbOption) RemoveStd(i int) {
	option.PpdbRegistration = append(option.PpdbRegistration[:i], option.PpdbRegistration[i+1:]...)
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

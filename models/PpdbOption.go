package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PpdbOption struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name,omitempty"`
	Type             string             `bson:"type,omitempty"`
	Quota            int                `bson:"quota,omitempty"`
	QuotaOld         int                `bson:"quota_old,omitempty"`
	TotalQuota       int                `bson:"total_quota,omitempty"`
	SchoolId         primitive.ObjectID `bson:"school,omitempty"`
	PpdbSchool       PpdbSchool         `bson:"ppdb_schools,omitempty"`
	PpdbRegistration []PpdbRegistration
}

func (ppdbOption PpdbOption) addItem(options []PpdbOption) []PpdbOption {
	return append(options, ppdbOption)
}

type PpdbOptionList struct {
	options []PpdbOption
}

func (optionList *PpdbOptionList) AddOpt(item PpdbOption) {
	optionList.options = append(optionList.options, item)
}

func (option *PpdbOption) AddStd(item PpdbRegistration) {
	option.PpdbRegistration = append(option.PpdbRegistration, item)
}
func (option *PpdbOption) RemoveStd(i int) {
	option.PpdbRegistration = append(option.PpdbRegistration[:i], option.PpdbRegistration[i+1:]...)
}

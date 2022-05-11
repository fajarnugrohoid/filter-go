package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

type PpdbRegistration struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name,omitempty"`
	OptionType         string             `bson:"option_type,omitempty"`
	FirstChoiceOption  primitive.ObjectID `bson:"first_choice_option,omitempty"`
	SecondChoiceOption primitive.ObjectID `bson:"second_choice_option,omitempty"`
	ThirdChoiceOption  primitive.ObjectID `bson:"third_choice_option,omitempty"`
	Score              float64            `bson:"score,omitempty"`
	Distance           float64
	Distance1          float64            `bson:"distance1,omitempty"`
	Distance2          float64            `bson:"distance2,omitempty"`
	Distance3          float64            `bson:"distance3,omitempty"`
	BirthDate          primitive.DateTime `bson:"birth_date,omitempty"`
	AcceptedStatus     int                `bson:"accepted_status"`
	AcceptedIndex      int                `bson:"accepted_index"`
	AcceptedChoiceId   primitive.ObjectID `bson:"accepted_choice_id"`
}

type ByScore []PpdbRegistration

func (m ByScore) Len() int           { return len(m) }
func (m ByScore) Less(i, j int) bool { return m[i].Score > m[j].Score }
func (m ByScore) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type ByDistance []PpdbRegistration

func (m ByDistance) Len() int           { return len(m) }
func (m ByDistance) Less(i, j int) bool { return m[i].Distance1 < m[j].Distance1 }
func (m ByDistance) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func SortByDistanceAndAge(members []PpdbRegistration) {
	sort.SliceStable(members, func(i, j int) bool {
		mi, mj := members[i], members[j]
		switch {
		case mi.Distance != mj.Distance:
			return mi.Distance < mj.Distance
		default:
			return mi.BirthDate < mj.BirthDate
		}
	})
}

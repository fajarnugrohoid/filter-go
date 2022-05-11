package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PpdbStatistic struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name,omitempty"`
	OptionType string             `bson:"option_type,omitempty"`
	Quota      int                `bson:"quota,omitempty"`
	SchoolId   primitive.ObjectID `bson:"school,omitempty"`
	Pg         float64
}

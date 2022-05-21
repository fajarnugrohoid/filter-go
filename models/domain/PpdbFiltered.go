package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PpdbFiltered struct {
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
	AcceptedIndex      int
}

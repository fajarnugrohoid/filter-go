package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PpdbRegistration struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name,omitempty"`
	FirstChoiceOption  primitive.ObjectID `bson:"first_choice_option,omitempty"`
	SecondChoiceOption primitive.ObjectID `bson:"second_choice_option,omitempty"`
	ThirdChoiceOption  primitive.ObjectID `bson:"third_choice_option,omitempty"`
	Score              float32            `bson:"score,omitempty"`
	AcceptedStatus     int
}

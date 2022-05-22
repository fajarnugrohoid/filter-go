package repository

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PpdbFilteredRepository interface {
	Save(ctx context.Context, database *mongo.Database, ppdbOptions []*domain.PpdbOption, optionType string) (*mongo.InsertManyResult, error)
	DeleteByOptionType(ctx context.Context, database *mongo.Database, option_type string)
	GetByOpt(ctx context.Context, database *mongo.Database, optionType string, optId primitive.ObjectID) []domain.PpdbFiltered
}

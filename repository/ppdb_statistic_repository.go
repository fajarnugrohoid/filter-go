package repository

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PpdbStatisticRepository interface {
	Insert(ctx context.Context, database *mongo.Database, ppdbStatistics []domain.PpdbStatistic, option_type string)
	DeleteByOptionType(ctx context.Context, database *mongo.Database, option_type string)
	GetAll(ctx context.Context, database *mongo.Database, optionType string) []domain.PpdbStatistic
	GetById(ctx context.Context, database *mongo.Database, optionType string, id primitive.ObjectID) []domain.PpdbStatistic
}

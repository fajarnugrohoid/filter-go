package repository

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type PpdbSchoolOptionRepository interface {
	GetSchoolOptionByLevel(ctx context.Context, database *mongo.Database, level string) []domain.PpdbOption
}

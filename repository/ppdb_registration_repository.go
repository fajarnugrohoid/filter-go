package repository

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PpdbRegistrationRepositoy interface {
	Save(ctx context.Context, database *mongo.Database, category domain.PpdbRegistration) domain.PpdbRegistration
	Update(ctx context.Context, database *mongo.Database, category domain.PpdbRegistration) domain.PpdbRegistration
	Delete(ctx context.Context, database *mongo.Database, category domain.PpdbRegistration)
	GetByFirstChoiceLevel(ctx context.Context, database *mongo.Database, level string, firstChoice primitive.ObjectID) []domain.PpdbRegistration
	GetAll(ctx context.Context, database *mongo.Database) []domain.PpdbRegistration
}

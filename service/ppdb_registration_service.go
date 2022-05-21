package service

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PpdbRegistrationService interface {
	FindByFirstChoiceLevel(ctx context.Context, level string, firstChoice primitive.ObjectID) []domain.PpdbRegistration
}

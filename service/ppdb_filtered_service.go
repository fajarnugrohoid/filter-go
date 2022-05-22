package service

import (
	"context"
	"filterisasi/models/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PpdbFilteredService interface {
	Save(ctx context.Context, ppdbOptions []*domain.PpdbOption, optionType string)
	DeleteByOptionType(ctx context.Context, option_type string)
	GetByOpt(ctx context.Context, optionType string, optId primitive.ObjectID) []domain.PpdbFiltered
}

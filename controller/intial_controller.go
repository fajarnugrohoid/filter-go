package controller

import (
	"context"
	"filterisasi/models/domain"
)

type InitialController interface {
	InitData(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, schoolOption []domain.PpdbOption) map[string][]*domain.PpdbOption
}

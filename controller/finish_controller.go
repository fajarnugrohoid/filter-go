package controller

import (
	"context"
	"filterisasi/models/domain"
	"github.com/sirupsen/logrus"
)

type FinishController interface {
	UpdateFiltered(ctx context.Context, optionTypes map[string][]*domain.PpdbOption)
	UpdateAllStatistic(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger)
	UpdateStatisticByOpt(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, optType string, logger *logrus.Logger)
	UpdateFilteredStatistic(ctx context.Context, optionTypes map[string][]*domain.PpdbOption, logger *logrus.Logger)
}

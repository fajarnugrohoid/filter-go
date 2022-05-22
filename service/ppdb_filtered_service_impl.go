package service

import (
	"context"
	"filterisasi/models/domain"
	"filterisasi/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PpdbFilteredServiceImpl struct {
	PpdbFilteredRepository repository.PpdbFilteredRepository
	DbClient               *mongo.Client
	DB                     *mongo.Database
}

func NewPpdbFilteredService(ppdbFilteredRepository repository.PpdbFilteredRepository, DB *mongo.Database) PpdbFilteredService {
	return &PpdbFilteredServiceImpl{PpdbFilteredRepository: ppdbFilteredRepository, DB: DB}
}

func (service PpdbFilteredServiceImpl) Save(ctx context.Context, ppdbOptions []*domain.PpdbOption, optionType string) {
	//TODO implement me

	db := service.DB
	var session mongo.Session
	var _ error

	//ppdbFiltereds := service.PpdbFilteredRepository.Save(ctx, db, ppdbOptions, optionType)

	_ = db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}

		_, err = service.PpdbFilteredRepository.Save(ctx, db, ppdbOptions, optionType)
		if err != nil {
			return err
		}

		if err = session.CommitTransaction(sessionContext); err != nil {
			return err
		}
		return nil
	})

}

func (service PpdbFilteredServiceImpl) DeleteByOptionType(ctx context.Context, option_type string) {
	//TODO implement me
	panic("implement me")
}

func (service PpdbFilteredServiceImpl) GetByOpt(ctx context.Context, optionType string, optId primitive.ObjectID) []domain.PpdbFiltered {
	//TODO implement me
	panic("implement me")
}

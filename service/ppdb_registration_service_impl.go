package service

import (
	"context"
	"filterisasi/models/domain"
	"filterisasi/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PpdbRegistrationServiceImpl struct {
	PpdbRegistrationRepository repository.PpdbRegistrationRepositoy
	DB                         *mongo.Database
}

func NewPpdbRegistrationService(ppdbRegistrationRepository repository.PpdbRegistrationRepositoy, DB *mongo.Database) PpdbRegistrationService {
	return &PpdbRegistrationServiceImpl{PpdbRegistrationRepository: ppdbRegistrationRepository, DB: DB}
}

func (service *PpdbRegistrationServiceImpl) FindByFirstChoiceLevel(ctx context.Context, level string, firstChoice primitive.ObjectID) []domain.PpdbRegistration {
	//TODO implement me
	// transaction
	db := service.DB
	ppdbRegistrations := service.PpdbRegistrationRepository.GetByFirstChoiceLevel(ctx, db, level, firstChoice)
	/*
		err = database.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
			err := sessionContext.StartTransaction()
			if err != nil {
				return err
			}

			_, err = col.InsertOne(sessionContext, bson.M{"_id": "1", "name": "berry"})
			if err != nil {
				return err
			}

			_, err = col.InsertOne(sessionContext, bson.M{"_id": "2", "name": "gucci"})
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)
				return err
			}
			if err = session.CommitTransaction(sessionContext); err != nil {
				return err
			}
			return nil
		})*/
	return ppdbRegistrations
}

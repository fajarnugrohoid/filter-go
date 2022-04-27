package utility

import (
	"filterisasi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindIndex(element primitive.ObjectID, data []models.PpdbOption) int {
	for k, v := range data {
		if element == v.Id {
			return k
		}
	}
	return -1 //not found.
}

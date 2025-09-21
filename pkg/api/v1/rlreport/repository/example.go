package repository

import (
	"fhir-sirs/pkg/api/v1/rlreport/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// the abstraction of resource functions. this part manages the resource operations e.g: view, create, update or delete
type Example interface {
	GetExampleValues(mongoConn *mongo.Client) ([]models.Example, error)
}

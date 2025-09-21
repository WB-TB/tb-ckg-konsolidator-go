package impl

import (
	"context"
	"fhir-sirs/app/config"
	"fhir-sirs/pkg/api/v1/rlreport/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Example struct{}

func NewExample() *Example {
	return &Example{}
}

// the implementation of resource functions
func (e *Example) GetExampleValues(mongoConn *mongo.Client) ([]models.Example, error) {
	var (
		results []models.Example
	)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.D{{}}
	collection := mongoConn.Database(config.GetConfig().MongoDatabaseName).Collection("examples")
	documentsReturned, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	if err = documentsReturned.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

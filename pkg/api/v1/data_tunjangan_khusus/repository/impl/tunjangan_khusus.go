package impl

import (
	"context"
	"fhir-sirs/app/config"
	"fhir-sirs/pkg/api/v1/data_tunjangan_khusus/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TunjanganKhususRepo struct{}

func NewTunjanganKhususRepo() *TunjanganKhususRepo {
	return &TunjanganKhususRepo{}
}

func (r *TunjanganKhususRepo) GetRealDataFiltered(ctx context.Context, conn *mongo.Client, orgID, nikDrSp, tanggal string) ([]models.DataTunjanganKhusus, error) {
	filter := bson.D{}

	if orgID != "" {
		filter = append(filter, bson.E{Key: "organization_id", Value: orgID})
	}
	if nikDrSp != "" {
		filter = append(filter, bson.E{Key: "nik_drSp", Value: nikDrSp})
	}
	if tanggal != "" {
		filter = append(filter, bson.E{Key: "tanggal", Value: tanggal})
	}

	collection := conn.Database(config.GetConfig().MongoDatabaseName).Collection(config.GetConfig().MongoCollectionName)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []models.DataTunjanganKhusus
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

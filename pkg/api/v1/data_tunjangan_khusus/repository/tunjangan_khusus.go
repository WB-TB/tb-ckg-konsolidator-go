package repository

import (
	"context"
	"fhir-sirs/pkg/api/v1/data_tunjangan_khusus/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TunjanganKhusus interface {
	GetRealDataFiltered(ctx context.Context, conn *mongo.Client, orgID, nikDrSp, tanggal string) ([]models.DataTunjanganKhusus, error)
}

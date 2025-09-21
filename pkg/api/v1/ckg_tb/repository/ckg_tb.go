package repository

import (
	"context"
	"fhir-sirs/pkg/api/v1/ckg_tb/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CKGTB interface {
	GeTbtDataFiltered(ctx context.Context, conn *mongo.Client, tanggal string, halaman int) (*models.DataSkriningTBOutput, error)
	PostTbPatientStatus(ctx context.Context, conn *mongo.Client, input []models.StatusPasienTBInput) ([]models.StatusPasienTBResult, error)
}

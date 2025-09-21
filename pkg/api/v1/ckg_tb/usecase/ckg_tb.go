package usecase

import (
	"context"
	"fhir-sirs/pkg/api/v1/ckg_tb/models"
	"fhir-sirs/pkg/api/v1/ckg_tb/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type DataCKGTB struct {
	Repo repository.CKGTB
	Conn *mongo.Client
}

func NewDataCKGTB(repo repository.CKGTB, conn *mongo.Client) *DataCKGTB {
	return &DataCKGTB{
		Repo: repo,
		Conn: conn,
	}
}

func (u *DataCKGTB) GeTbtDataFiltered(tanggal string, halaman int) (*models.DataSkriningTBOutput, error) {
	ctx := context.Background()
	return u.Repo.GeTbtDataFiltered(ctx, u.Conn, tanggal, halaman)
}

func (u *DataCKGTB) PostTbPatientStatus(input []models.StatusPasienTBInput) ([]models.StatusPasienTBResult, error) {
	ctx := context.Background()
	return u.Repo.PostTbPatientStatus(ctx, u.Conn, input)
}

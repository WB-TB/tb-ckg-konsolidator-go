package usecase

import (
	"context"
	"fhir-sirs/pkg/api/v1/data_tunjangan_khusus/models"
	"fhir-sirs/pkg/api/v1/data_tunjangan_khusus/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type DataTunjanganKhusus struct {
	Repo repository.TunjanganKhusus
	Conn *mongo.Client
}

func NewDataTunjanganKhusus(repo repository.TunjanganKhusus, conn *mongo.Client) *DataTunjanganKhusus {
	return &DataTunjanganKhusus{
		Repo: repo,
		Conn: conn,
	}
}

func (u *DataTunjanganKhusus) GetDummyData() []models.DataTunjanganKhusus {
	return []models.DataTunjanganKhusus{
		{
			NamaDrSp:                   "Dr. Andi SpPD",
			NikDrSp:                    "12345678901234",
			PractitionerID:             "13193826263",
			NamaFasyankes:              "RSUD Kota A",
			OrganizationID:             "100121469",
			Tanggal:                    "2025-06-15",
			JumlahPasien:               23,
			JumlahBacaanHasilPenunjang: 15,
			JenisPelayanan:             "Rawat Inap",
		},
	}
}

func (u *DataTunjanganKhusus) GetDummyDataFiltered(orgID, nikDrSp, tanggal string) []models.DataTunjanganKhusus {
	dummy := u.GetDummyData()
	var filtered []models.DataTunjanganKhusus

	for _, d := range dummy {
		if (orgID == "" || d.OrganizationID == orgID) &&
			(nikDrSp == "" || d.NikDrSp == nikDrSp) &&
			(tanggal == "" || d.Tanggal == tanggal) {
			filtered = append(filtered, d)
		}
	}
	return filtered
}

func (u *DataTunjanganKhusus) GetRealDataFiltered(orgID, nikDrSp, tanggal string) ([]models.DataTunjanganKhusus, error) {
	ctx := context.Background()
	return u.Repo.GetRealDataFiltered(ctx, u.Conn, orgID, nikDrSp, tanggal)
}

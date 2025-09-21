package data_tunjangan_khusus

import "fhir-sirs/pkg/api/v1/data_tunjangan_khusus/models"

type Service interface {
	GetDummyData() []models.DataTunjanganKhusus
}

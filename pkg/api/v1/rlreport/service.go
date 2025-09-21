package rlreport

import (
	"fhir-sirs/pkg/api/v1/rlreport/models"
)

type Service interface {
	// GetExampleValues() []models.Example
	GetRL34() models.RL34Visitors
	GetRL35() models.RL35Visits
	GetRL51() models.RL51Morbidities
}

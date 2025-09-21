package usecase

import "fhir-sirs/pkg/api/v1/rlreport/models"

type RLReport struct{}

func NewRLReport() *RLReport {
	return &RLReport{}
}

func (r *RLReport) GetRL34() models.RL34Visitors {
	return models.RL34Visitors{
		Month:             "2025-06",
		OrganizationID:    "ORG001",
		NewVisitors:       1234,
		ReturningVisitors: 4321,
		TotalVisitors:     5555,
	}
}

func (r *RLReport) GetRL35() models.RL35Visits {
	return models.RL35Visits{
		Month:          "2025-06",
		OrganizationID: "ORG001",
		Domicile:       "local",
		Visits: []models.RL35Visit{
			{
				ActivityID:   "ACT001",
				ActivityName: "Penyakit Dalam",
				Male:         321,
				Female:       418,
				Total:        739,
			},
		},
	}
}

func (r *RLReport) GetRL51() models.RL51Morbidities {
	return models.RL51Morbidities{
		Month:          "2025-06",
		OrganizationID: "ORG001",
		Records: []models.MorbidityRecord{
			{
				ICD10:     "I10",
				Diagnosis: "Essential (primary) hypertension",
				NewCases: []models.NewCase{
					{
						AgeID:          "001",
						AgeName:        "< 1 jam",
						MaleNewCases:   150,
						FemaleNewCases: 200,
					},
				},
				MaleNewCases:   150,
				FemaleNewCases: 200,
				TotalNewCases:  350,
				MaleVisits:     150,
				FemaleVisits:   200,
				TotalVisits:    350,
			},
		},
	}
}

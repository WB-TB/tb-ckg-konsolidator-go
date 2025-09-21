package models

type RL34Visitors struct {
	Month             string `json:"month" bson:"month"`
	OrganizationID    string `json:"organization_id" bson:"organization_id"`
	NewVisitors       int    `json:"new_visitors" bson:"new_visitors"`
	ReturningVisitors int    `json:"returning_visitors" bson:"returning_visitors"`
	TotalVisitors     int    `json:"total_visitors" bson:"total_visitors"`
}

type RL35Visit struct {
	ActivityID   string `json:"activity_id" bson:"activity_id"`
	ActivityName string `json:"activity_name" bson:"activity_name"`
	Male         int    `json:"male" bson:"male"`
	Female       int    `json:"female" bson:"female"`
	Total        int    `json:"total" bson:"total"`
}

type RL35Visits struct {
	Month          string      `json:"month" bson:"month"`
	OrganizationID string      `json:"organization_id" bson:"organization_id"`
	Domicile       string      `json:"domicile" bson:"domicile"`
	Visits         []RL35Visit `json:"visits" bson:"visits"`
}

type NewCase struct {
	AgeID          string `json:"age_id" bson:"age_id"`
	AgeName        string `json:"age_name" bson:"age_name"`
	MaleNewCases   int    `json:"male_new_cases" bson:"male_new_cases"`
	FemaleNewCases int    `json:"female_new_cases" bson:"female_new_cases"`
}

type MorbidityRecord struct {
	ICD10          string    `json:"icd10" bson:"icd10"`
	Diagnosis      string    `json:"diagnosis" bson:"diagnosis"`
	NewCases       []NewCase `json:"new_cases" bson:"new_cases"`
	MaleNewCases   int       `json:"male_new_cases" bson:"male_new_cases"`
	FemaleNewCases int       `json:"female_new_cases" bson:"female_new_cases"`
	TotalNewCases  int       `json:"total_new_cases" bson:"total_new_cases"`
	MaleVisits     int       `json:"male_visits" bson:"male_visits"`
	FemaleVisits   int       `json:"female_visits" bson:"female_visits"`
	TotalVisits    int       `json:"total_visits" bson:"total_visits"`
}

type RL51Morbidities struct {
	Month          string            `json:"month" bson:"month"`
	OrganizationID string            `json:"organization_id" bson:"organization_id"`
	Records        []MorbidityRecord `json:"records" bson:"records"`
}

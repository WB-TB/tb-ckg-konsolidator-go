package usecase

import "fhir-sirs/pkg/api/v1/rlreport/models"

// the implementation of service, this part manages all the business flows
func (e *Example) GetExampleValues() []models.Example {
	var (
		results []models.Example
	)

	values, err := e.repository.example.GetExampleValues(e.conn.readWrite)
	if err != nil {
		return nil
	}

	if len(values) > 0 {
		for i, v := range values {
			v.Field3 = e.GetFieldValueExample(i) // call the helper function
			results = append(results, v)
		}
	}

	return results
}

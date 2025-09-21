package models

// data model of resource
type Example struct {
	Field1 string `json:"field_1" bson:"field_1"`
	Field2 string `json:"field_2" bson:"field_2"`
	Field3 string `json:"field_3"`
}

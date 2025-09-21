package usecase

import (
	example "fhir-sirs/pkg/api/v1/rlreport/repository"
	exampleImpl "fhir-sirs/pkg/api/v1/rlreport/repository/impl"

	"go.mongodb.org/mongo-driver/mongo"
)

type Connection struct {
	readWrite *mongo.Client
}

type Repository struct {
	example example.Example
}

// this struct will be used in the implementation of services
type Example struct {
	conn       Connection
	repository Repository
}

func NewExampleRepository() Repository {
	return Repository{
		example: exampleImpl.NewExample(),
	}
}

func New(mongoConn *mongo.Client, repository Repository) *Example {
	return &Example{
		conn: Connection{
			readWrite: mongoConn,
		},
		repository: repository,
	}
}

// wrap up the db connector to be used on the entrypoint
func Initialize(mongoConn *mongo.Client) *Example {
	return New(mongoConn, NewExampleRepository())
}

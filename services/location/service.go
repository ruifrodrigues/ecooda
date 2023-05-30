package location

import (
	"github.com/ruifrodrigues/ecooda/api"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

var allowedSorts = []string{"id", "uuid"}

type Service struct {
	pb.UnimplementedLocationServiceServer
	dbCtx  Database
	fields Fields
}

func NewLocationService(conf Config) *Service {
	conn, err := conf.Database.OpenConnection()
	if err != nil {
		panic("Database Connection is not open!")
	}

	fields := api.NewFields()

	fields.Location = &api.FieldTypes{
		Guarded: []string{
			"id",
		},
		Default: []string{
			"uuid",
			"created_at",
			"updated_at",
		},
		Available: []string{
			"uuid",
			"name",
			"type",
			"parent",
			"created_at",
			"updated_at",
		},
	}

	return &Service{
		dbCtx:  conn,
		fields: fields,
	}
}

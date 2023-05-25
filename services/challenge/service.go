package challenge

import (
	"github.com/ruifrodrigues/ecooda/api"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

var allowedSorts = []string{"id", "uuid"}

type Service struct {
	pb.UnimplementedChallengeServiceServer
	dbCtx  Database
	fields Fields
}

func NewChallengeService(conf Config) *Service {
	conn, err := conf.Database.OpenConnection()
	if err != nil {
		panic("Database Connection is not open!")
	}

	fields := api.NewFields()

	fields.Challenge = &api.FieldTypes{
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
			"description",
			"street",
			"postcode",
			"latitude",
			"longitude",
			"thumbnail",
			"gallery",
			"categories",
			"created_at",
			"updated_at",
		},
	}

	fields.Category = &api.FieldTypes{
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
			"created_at",
			"updated_at",
		},
	}

	return &Service{
		dbCtx:  conn,
		fields: fields,
	}
}

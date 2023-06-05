package location

import (
	"github.com/ruifrodrigues/ecooda/api"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

var allowedSorts = []string{"id", "uuid"}

type Service struct {
	pb.UnimplementedLocationServiceServer
	conf   config.Config
	fields api.Fields
}

func NewLocationService(conf config.Config) *Service {
	if err := conf.Database.OpenConnection(); err != nil {
		panic("Database Connection not open!")
	}

	fields := api.NewFields()
	fields.Location = new(api.FieldTypes)

	fields.Location.Guarded = []string{
		"id",
	}
	fields.Location.Default = []string{
		"uuid",
		"created_at",
		"updated_at",
	}
	fields.Location.Available = []string{
		"uuid",
		"name",
		"type",
		"parents",
		"created_at",
		"updated_at",
	}

	return &Service{
		conf:   conf,
		fields: fields,
	}
}

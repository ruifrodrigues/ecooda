package challenge

import (
	"github.com/ruifrodrigues/ecooda/api"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

var allowedSorts = []string{"id", "uuid"}

type Service struct {
	pb.UnimplementedChallengeServiceServer
	conf   config.Config
	fields api.Fields
}

func NewChallengeService(conf config.Config) *Service {
	if err := conf.Database.OpenConnection(); err != nil {
		panic("Database Connection is not open!")
	}

	fields := api.NewFields()
	fields.Challenge = new(api.FieldTypes)
	fields.Challenge.Guarded = []string{
		"id",
	}

	fields.Challenge.Default = []string{
		"uuid",
		"created_at",
		"updated_at",
	}

	fields.Challenge.Available = []string{
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
		"locations",
		"created_at",
		"updated_at",
	}

	fields.Category = new(api.FieldTypes)
	fields.Category.Guarded = []string{
		"id",
	}
	fields.Category.Default = []string{
		"uuid",
		"created_at",
		"updated_at",
	}
	fields.Category.Available = []string{
		"uuid",
		"name",
		"created_at",
		"updated_at",
	}

	return &Service{
		conf:   conf,
		fields: fields,
	}
}

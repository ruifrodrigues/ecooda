package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository struct {
	config   config.Config
	location *Location
}

type Behaviour interface {
	Get(uuid string) (Aggregate, error)
	Save() error
}

func NewRepository(config config.Config) Behaviour {
	return &Repository{
		config:   config,
		location: &Location{},
	}
}

func (r *Repository) Get(uuid string) (Aggregate, error) {
	dbCtx := r.config.Database.Ctx()

	transaction := dbCtx.Select("*").
		Where("uuid=?", uuid).
		Preload("Challenges").
		First(r.location)

	if transaction.Error != nil {
		return nil, status.Error(codes.NotFound, "Aggregate Not Found >> "+transaction.Error.Error())
	}

	return NewAggregateRoot(r.config, r.location), nil
}

func (r *Repository) Save() error {
	dbCtx := r.config.Database.Ctx()

	transaction := dbCtx.Save(r.location)
	if transaction.Error != nil {
		return status.Error(codes.Internal, "Aggregate Not Saved >> "+transaction.Error.Error())
	}

	return nil
}

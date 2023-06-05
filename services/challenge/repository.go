package challenge

import (
	"github.com/ruifrodrigues/ecooda/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository struct {
	dbCtx     config.Database
	challenge *Challenge
}

type Behaviour interface {
	Get(uuid string) (Aggregate, error)
	Save() error
}

func NewRepository(dbCtx config.Database) Behaviour {
	return &Repository{
		dbCtx,
		&Challenge{},
	}
}

func (r *Repository) Get(uuid string) (Aggregate, error) {
	transaction := r.dbCtx.Select("*").
		Preload("Categories").
		Where("uuid=?", uuid).
		First(r.challenge)

	if transaction.Error != nil {
		return nil, status.Error(codes.NotFound, "Aggregate Not Found >> "+transaction.Error.Error())
	}

	return NewAggregateRoot(r.challenge), nil
}

func (r *Repository) Save() error {
	err := r.dbCtx.Model(r.challenge).Association("Categories").Replace(r.challenge.Categories)
	if err != nil {
		return status.Error(codes.Internal, "Aggregate Not Saved >> "+err.Error())
	}

	return nil
}

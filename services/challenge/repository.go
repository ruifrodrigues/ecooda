package challenge

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository struct {
	dbCtx     Database
	challenge *Challenge
}

type IRepository interface {
	Get(uuid string) *AggregateRoot
	Save() error
}

func NewRepository(dbCtx Database) *Repository {
	return &Repository{
		dbCtx,
		&Challenge{},
	}
}

func (r *Repository) Get(uuid string) *AggregateRoot {
	r.dbCtx.Select("*").
		Preload("Categories").
		Where("uuid=?", uuid).
		First(r.challenge)

	return NewAggregateRoot(r.challenge)
}

func (r *Repository) Save() error {
	err := r.dbCtx.Model(r.challenge).Association("Categories").Replace(r.challenge.Categories)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

package challenge

import (
	"github.com/ruifrodrigues/ecooda/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository struct {
	conf      config.Config
	challenge *Challenge
}

type Behaviour interface {
	Get(uuid string) (Aggregate, error)
	Save() error
}

func NewRepository(conf config.Config) Behaviour {
	repository := new(Repository)
	repository.conf = conf

	return repository
}

func (r *Repository) Get(uuid string) (Aggregate, error) {
	var err error

	dbCtx := r.conf.Database.Ctx()

	r.challenge, err = NewQuery(dbCtx).LoadAggregateRoot(uuid)
	if err != nil {
		return nil, err
	}

	return NewAggregateRoot(r.conf, r.challenge), nil
}

func (r *Repository) Save() error {
	dbCtx := r.conf.Database.Ctx()

	err := NewQuery(dbCtx).SaveChallenge(r.challenge)
	if err != nil {
		return status.Error(codes.Internal, "Aggregate Not Saved >> "+err.Error())
	}

	return nil
}

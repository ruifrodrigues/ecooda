package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	"gorm.io/gorm"
)

type Repository struct {
	conf     config.Config
	location *Location
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

	r.location, err = NewQuery(dbCtx).LoadAggregateRoot(uuid)
	if err != nil {
		return nil, err
	}

	return NewAggregateRoot(r.conf, r.location), nil
}

func (r *Repository) Save() error {
	dbCtx := r.conf.Database.Ctx()

	// Start Transaction
	err := dbCtx.Transaction(func(tx *gorm.DB) error {
		query := NewQuery(dbCtx)

		challengesUuids, err := query.GetChallengesUuids(r.location, 0, 0, "")
		if err != nil {
			return err
		}

		if err = query.DeleteChallenge(r.location, challengesUuids); err != nil {
			return err
		}

		if err = query.SaveLocation(r.location); err != nil {
			return err
		}

		return nil
	})

	return err
}

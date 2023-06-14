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
	return &Repository{
		conf: conf,
	}
}

func (r *Repository) Get(uuid string) (Aggregate, error) {
	var err error

	r.location, err = NewQuery(r.conf).LoadAggregateRoot(uuid)
	if err != nil {
		return nil, err
	}

	return NewAggregateRoot(r.conf, r.location), nil
}

func (r *Repository) Save() error {
	// Start Transaction
	err := r.conf.Database.Ctx().Transaction(func(tx *gorm.DB) error {
		query := NewQuery(r.conf)

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

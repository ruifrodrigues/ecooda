package challenge

import (
	"github.com/ruifrodrigues/ecooda/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RecordNotFound     = "record not found"
	NoRecordsAvailable = "No records available"
)

type Reader interface {
	LoadAggregateRoot(uuid string) (*Challenge, error)
	CountChallenges() int32
	CountCategories() int32
	GetChallengeByUuid(uuid string) (*Challenge, error)
	GetChallenges(offset, limit int, order string, uuids []string) ([]*Challenge, error)
	GetCategoryByUuid(uuid string) (*Category, error)
	GetCategories(offset, limit int, order string) ([]*Category, error)
	CreateChallenge(fields map[string]interface{}) error
	CreateCategory(fields map[string]interface{}) error
	UpdateCategory(uuid string, fields map[string]interface{}) error
	DeleteChallenge(uuid string) error
	DeleteCategory(uuid string) error
	SaveChallenge(location *Challenge) error
}

type Query struct {
	conf config.Config
}

func NewQuery(conf config.Config) Reader {
	return &Query{
		conf,
	}
}

func (q *Query) LoadAggregateRoot(uuid string) (*Challenge, error) {
	var challenge *Challenge

	query := q.conf.Database.Ctx().
		Select("*").
		Preload("Categories").
		Where("uuid=?", uuid).
		First(&challenge)

	if query.Error != nil {
		return nil, status.Error(codes.NotFound, "Aggregate Not Found >> "+query.Error.Error())
	}

	return challenge, nil
}

func (q *Query) CountChallenges() int32 {
	var records []*Challenge
	var count int64 = 0

	q.conf.Database.Ctx().
		Find(&records).
		Count(&count)

	return int32(count)
}

func (q *Query) CountCategories() int32 {
	var records []*Category
	var count int64 = 0

	q.conf.Database.Ctx().
		Find(&records).
		Count(&count)

	return int32(count)
}

func (q *Query) GetChallengeByUuid(uuid string) (*Challenge, error) {
	var record *Challenge

	query := q.conf.Database.Ctx().
		Preload("Categories").
		Where("uuid=?", uuid).
		Find(&record)

	if query.Error != nil {
		return nil, status.Error(codes.Internal, query.Error.Error())
	}

	if record.UUID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, status.Error(codes.NotFound, RecordNotFound)
	}

	return record, nil
}

func (q *Query) GetChallenges(offset, limit int, order string, uuids []string) ([]*Challenge, error) {
	var records []*Challenge

	query := q.conf.Database.Ctx().
		Preload("Categories")

	if len(uuids) > 0 {
		query = query.Where("uuid in ?", uuids)
	}

	query.
		Offset(offset).
		Limit(limit).
		Order(order).
		Find(&records)

	if len(records) == 0 {
		return nil, status.Error(codes.NotFound, NoRecordsAvailable)
	}

	return records, nil
}

func (q *Query) GetCategoryByUuid(uuid string) (*Category, error) {
	var record *Category

	query := q.conf.Database.Ctx().
		Where("uuid=?", uuid).
		Find(&record)

	if query.Error != nil {
		return nil, status.Error(codes.Internal, query.Error.Error())
	}

	if record.UUID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, status.Error(codes.NotFound, RecordNotFound)
	}

	return record, nil
}

func (q *Query) GetCategories(offset, limit int, order string) ([]*Category, error) {
	var records []*Category

	q.conf.Database.Ctx().
		Offset(offset).
		Limit(limit).
		Order(order).
		Find(&records)

	if len(records) == 0 {
		return nil, status.Error(codes.NotFound, NoRecordsAvailable)
	}

	return records, nil
}

func (q *Query) CreateChallenge(fields map[string]interface{}) error {
	query := q.conf.Database.Ctx().
		Model(&Challenge{}).
		Create(fields)

	if query.Error != nil {
		return status.Error(codes.AlreadyExists, query.Error.Error())
	}

	return nil
}

func (q *Query) CreateCategory(fields map[string]interface{}) error {
	query := q.conf.Database.Ctx().
		Model(&Category{}).
		Create(fields)

	if query.Error != nil {
		return status.Error(codes.AlreadyExists, query.Error.Error())
	}

	return nil
}

func (q *Query) UpdateCategory(uuid string, fields map[string]interface{}) error {
	query := q.conf.Database.Ctx().
		Model(&Category{}).
		Where("uuid=?", uuid).
		Updates(fields)

	if query.Error != nil && query.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, query.Error.Error())
	}

	return nil
}

func (q *Query) DeleteChallenge(uuid string) error {
	query := q.conf.Database.Ctx().
		Where("uuid=?", uuid).
		Delete(&Challenge{})

	if query.Error != nil && query.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, query.Error.Error())
	}

	return nil
}

func (q *Query) DeleteCategory(uuid string) error {
	query := q.conf.Database.Ctx().
		Where("uuid=?", uuid).
		Delete(&Category{})

	if query.Error != nil && query.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, query.Error.Error())
	}

	return nil
}

func (q *Query) SaveChallenge(challenge *Challenge) error {
	query := q.conf.Database.Ctx().
		Model(challenge).
		Association("Categories").
		Replace(challenge.Categories)

	if query != nil {
		return status.Error(codes.Internal, "Aggregate Not Saved >> "+query.Error())
	}

	return nil
}

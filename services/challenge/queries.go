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
	CountChallenges() int32
	FindChallengeByUuid(uuid string) (*Challenge, error)
	FindChallenges(offset, limit int, order string) ([]*Challenge, error)
	CreateChallenge(fields map[string]interface{}) error
	DeleteChallenge(uuid string) error
	CountCategories() int32
	FindCategoryByUuid(uuid string) (*Category, error)
	FindCategories(offset, limit int, order string) ([]*Category, error)
	CreateCategory(fields map[string]interface{}) error
	UpdateCategory(uuid string, fields map[string]interface{}) error
	DeleteCategory(uuid string) error
}

type Query struct {
	dbCtx config.Database
}

func NewQuery(dbCtx config.Database) Reader {
	return &Query{
		dbCtx,
	}
}

func (q *Query) CountChallenges() int32 {
	var records []*Challenge
	var count int64 = 0

	q.dbCtx.Find(&records).Count(&count)

	return int32(count)
}

func (q *Query) FindChallengeByUuid(uuid string) (*Challenge, error) {
	var record *Challenge

	transaction := q.dbCtx.
		Preload("Categories").
		Where("uuid=?", uuid).
		Find(&record)

	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return nil, status.Error(codes.NotFound, transaction.Error.Error())
	}

	return record, nil
}

func (q *Query) FindChallenges(offset, limit int, order string) ([]*Challenge, error) {
	var records []*Challenge

	q.dbCtx.
		Preload("Categories").
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
	result := q.dbCtx.Model(&Challenge{}).Create(fields)
	if result.Error != nil {
		return status.Error(codes.AlreadyExists, result.Error.Error())
	}

	return nil
}

func (q *Query) DeleteChallenge(uuid string) error {
	transaction := q.dbCtx.Where("uuid=?", uuid).Delete(&Challenge{})
	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, transaction.Error.Error())
	}

	return nil
}

func (q *Query) CountCategories() int32 {
	var records []*Category
	var count int64 = 0

	q.dbCtx.Find(&records).Count(&count)

	return int32(count)
}

func (q *Query) FindCategoryByUuid(uuid string) (*Category, error) {
	var record *Category

	transaction := q.dbCtx.
		Preload("Categories").
		Where("uuid=?", uuid).
		Find(&record)

	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return nil, status.Error(codes.NotFound, transaction.Error.Error())
	}

	return record, nil
}

func (q *Query) FindCategories(offset, limit int, order string) ([]*Category, error) {
	var records []*Category

	q.dbCtx.
		Offset(offset).
		Limit(limit).
		Order(order).
		Find(&records)

	if len(records) == 0 {
		return nil, status.Error(codes.NotFound, NoRecordsAvailable)
	}

	return records, nil
}

func (q *Query) CreateCategory(fields map[string]interface{}) error {
	result := q.dbCtx.Model(&Category{}).Create(fields)
	if result.Error != nil {
		return status.Error(codes.AlreadyExists, result.Error.Error())
	}

	return nil
}

func (q *Query) UpdateCategory(uuid string, fields map[string]interface{}) error {
	transaction := q.dbCtx.
		Model(&Category{}).
		Where("uuid=?", uuid).
		Updates(fields)

	if transaction.Error != nil && transaction.Error.Error() == "record not found" {
		return status.Error(codes.NotFound, transaction.Error.Error())
	}

	return nil
}

func (q *Query) DeleteCategory(uuid string) error {
	transaction := q.dbCtx.Where("uuid=?", uuid).Delete(&Category{})
	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, transaction.Error.Error())
	}

	return nil
}

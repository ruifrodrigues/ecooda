package location

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
	Count() int32
	FindItemByUuid(uuid string, fields string) (*Location, error)
	FindItems(offset, limit int, order string) ([]*Location, error)
	Create(fields map[string]interface{}) error
	Delete(uuid string) error
}

type Query struct {
	dbCtx config.Database
}

func NewQuery(dbCtx config.Database) Reader {
	return &Query{
		dbCtx,
	}
}

func (q *Query) Count() int32 {
	var records []*Location
	var count int64 = 0

	q.dbCtx.Find(&records).Count(&count)

	return int32(count)
}

func (q *Query) FindItemByUuid(uuid string, fields string) (*Location, error) {
	var record *Location

	transaction := q.dbCtx.
		Select(fields).
		Where("uuid=?", uuid).
		Find(&record)

	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return nil, status.Error(codes.NotFound, transaction.Error.Error())
	}

	return record, nil
}

func (q *Query) FindItems(offset, limit int, order string) ([]*Location, error) {
	var records []*Location

	q.dbCtx.Offset(offset).Limit(limit).Order(order).Find(&records)
	if len(records) == 0 {
		return nil, status.Error(codes.NotFound, NoRecordsAvailable)
	}

	return records, nil
}

func (q *Query) Create(fields map[string]interface{}) error {
	result := q.dbCtx.Model(&Location{}).Create(fields)
	if result.Error != nil {
		return status.Error(codes.AlreadyExists, result.Error.Error())
	}

	return nil
}

func (q *Query) Delete(uuid string) error {
	transaction := q.dbCtx.Where("uuid=?", uuid).Delete(&Location{})
	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, transaction.Error.Error())
	}

	return nil
}

package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RecordNotFound     = "record not found"
	NoRecordsAvailable = "No records available"
)

type Reader interface {
	LoadAggregateRoot(uuid string) (*Location, error)
	CountLocations() int32
	CountChallenges(location *Location) int32
	GetLocationByUuid(uuid string, fields string) (*Location, error)
	GetLocations(offset, limit int, order string) ([]*Location, error)
	GetLocationsByChallengeUuid(uuid string, fields string) (*Location, error)
	GetChallengesUuids(location *Location, offset int, limit int, order string) ([]string, error)
	GetLocation(uuid string, locationType pb.LocationType) (*Location, error)
	DeleteLocation(uuid string) error
	DeleteChallenge(location *Location, challengeUuids []string) error
	CreateLocation(fields map[string]interface{}) error
	SaveLocation(location *Location) error
}

type Query struct {
	dbCtx config.Database
}

func NewQuery(dbCtx config.Database) Reader {
	return &Query{
		dbCtx,
	}
}

func (q *Query) LoadAggregateRoot(uuid string) (*Location, error) {
	var location *Location

	query := q.dbCtx.Select("*").
		Where("uuid=?", uuid).
		Preload("Challenges").
		First(&location)

	if query.Error != nil {
		return nil, status.Error(codes.NotFound, "Aggregate Not Found >> "+query.Error.Error())
	}

	return location, nil
}

func (q *Query) CountLocations() int32 {
	var records []*Location
	var count int64 = 0

	q.dbCtx.Find(&records).Count(&count)

	return int32(count)
}

func (q *Query) CountChallenges(location *Location) int32 {
	var records []*LocationChallenges
	var count int64 = 0

	q.dbCtx.Where("location_id=?", location.ID).Find(&records).Count(&count)

	return int32(count)
}

func (q *Query) GetLocationByUuid(uuid string, fields string) (*Location, error) {
	var record *Location

	query := q.dbCtx.
		Select(fields).
		Where("uuid=?", uuid).
		Find(&record)

	if query.Error != nil && query.Error.Error() == RecordNotFound {
		return nil, status.Error(codes.NotFound, query.Error.Error())
	}

	return record, nil
}

func (q *Query) GetLocations(offset, limit int, order string) ([]*Location, error) {
	var records []*Location

	q.dbCtx.Offset(offset).Limit(limit).Order(order).Find(&records)
	if len(records) == 0 {
		return nil, status.Error(codes.NotFound, NoRecordsAvailable)
	}

	return records, nil
}

func (q *Query) GetLocationsByChallengeUuid(uuid string, fields string) (*Location, error) {
	var record *LocationChallenges

	query := q.dbCtx.
		Select(fields).
		Preload("Location").
		Where("challenge_uuid=?", uuid).
		First(&record)

	if query.Error != nil && query.Error.Error() == RecordNotFound {
		return nil, status.Error(codes.NotFound, query.Error.Error())
	}

	return record.Location, nil
}

func (q *Query) GetChallengesUuids(location *Location, offset int, limit int, order string) ([]string, error) {
	var locationChallenges []*LocationChallenges

	query := q.dbCtx.
		Select("challenge_uuid").
		Where("location_id=?", location.ID)

	if offset != 0 && limit != 0 && order != "" {
		query = query.Offset(offset).Limit(limit).Order(order)
	}

	query = query.Find(&locationChallenges)

	if query.Error != nil {
		return nil, status.Error(codes.Internal, "Could not find Challenge on LocationChallenges table >> "+query.Error.Error())
	}

	return extractChallengeUuids(locationChallenges), nil
}

func (q *Query) GetLocation(uuid string, locationType pb.LocationType) (*Location, error) {
	country := new(Location)

	query := q.dbCtx.Ctx().
		Select("id").
		Where("uuid=?", uuid).
		Where("type=?", int32(locationType)).
		First(country)

	if query.Error != nil {
		return nil, status.Error(codes.NotFound, "Location Not Found >> "+query.Error.Error())
	}

	return country, nil
}

func (q *Query) DeleteLocation(uuid string) error {

	query := q.dbCtx.Where("uuid=?", uuid).Delete(&Location{})
	if query.Error != nil && query.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, query.Error.Error())
	}

	return nil
}

func (q *Query) DeleteChallenge(location *Location, challengeUuids []string) error {
	for _, challengeUuid := range challengeUuids {
		if !utils.InArray(challengeUuid, extractChallengeUuids(location.Challenges)) {
			query := q.dbCtx.
				Where("location_id=?", location.ID).
				Where("challenge_uuid=?", challengeUuid).
				Delete(&LocationChallenges{})

			if query.Error != nil {
				return status.Error(codes.Internal, "Could Not Delete Challenge from Location >> "+query.Error.Error())
			}
		}
	}

	return nil
}

func (q *Query) CreateLocation(fields map[string]interface{}) error {
	result := q.dbCtx.Model(&Location{}).Create(fields)
	if result.Error != nil {
		return status.Error(codes.AlreadyExists, result.Error.Error())
	}

	return nil
}

func (q *Query) SaveLocation(location *Location) error {
	if query := q.dbCtx.Save(location); query.Error != nil {
		return status.Error(codes.Internal, "Aggregate Not Saved >> "+query.Error.Error())
	}

	return nil
}

func extractChallengeUuids(locationChallenges []*LocationChallenges) []string {
	var challengeUuids []string
	for _, locationChallenge := range locationChallenges {
		challengeUuids = append(challengeUuids, locationChallenge.ChallengeUUID.String())
	}

	return challengeUuids
}

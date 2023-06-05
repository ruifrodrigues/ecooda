package location

import (
	_uuid "github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"github.com/ruifrodrigues/ecooda/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Aggregate interface {
	ChangeName(name string)
	ChangeType(locationType pb.LocationType)
	AddCountry(uuid string) error
	AddRegion(uuid string) error
	RemoveCountry(uuid string) error
	RemoveRegion(uuid string) error
	AddChallenge(uuid string) error
	RemoveChallenge(uuid string) error
}

type AggregateRoot struct {
	conf     config.Config
	location *Location
}

func NewAggregateRoot(conf config.Config, location *Location) Aggregate {
	return &AggregateRoot{
		conf:     conf,
		location: location,
	}
}

func (a *AggregateRoot) ChangeName(name string) {
	if name != "" {
		a.location.Name = name
	}
}

func (a *AggregateRoot) ChangeType(locationType pb.LocationType) {
	if locationType != pb.LocationType_LOCATION_TYPE_UNSPECIFIED {
		a.location.Type = int32(locationType)
	}
}

func (a *AggregateRoot) AddCountry(uuid string) error {
	if pb.LocationType(a.location.Type) != pb.LocationType_LOCATION_TYPE_REGION {
		message := "Can only add LOCATION_TYPE_COUNTRY to LOCATION_TYPE_REGION, current type is " +
			pb.LocationType_name[a.location.Type]

		return status.Error(codes.FailedPrecondition, message)
	}

	if err := utils.ValidateUuid(uuid); err != nil {
		return err
	}

	country := new(Location)

	transaction := a.conf.Database.Ctx().
		Select("id").
		Where("uuid=?", uuid).
		Where("type=?", int32(pb.LocationType_LOCATION_TYPE_COUNTRY)).
		First(country)

	if transaction.Error != nil {
		return status.Error(codes.NotFound, "Location Not Found >> "+transaction.Error.Error())
	}

	region := a.location
	region.ParentID = &country.ID

	return nil
}

func (a *AggregateRoot) AddRegion(uuid string) error {
	if pb.LocationType(a.location.Type) != pb.LocationType_LOCATION_TYPE_CITY {
		message := "Can only add LOCATION_TYPE_REGION to LOCATION_TYPE_CITY, current type is " +
			pb.LocationType_name[a.location.Type]

		return status.Error(codes.FailedPrecondition, message)
	}

	if err := utils.ValidateUuid(uuid); err != nil {
		return err
	}

	region := new(Location)

	transaction := a.conf.Database.Ctx().
		Select("id").
		Where("uuid=?", uuid).
		Where("type=?", int32(pb.LocationType_LOCATION_TYPE_REGION)).
		First(region)

	if transaction.Error != nil {
		return status.Error(codes.NotFound, "Location Not Found >> "+transaction.Error.Error())
	}

	city := a.location
	city.ParentID = &region.ID

	return nil
}

func (a *AggregateRoot) RemoveCountry(uuid string) error {
	if err := utils.ValidateUuid(uuid); err != nil {
		return err
	}

	country := new(Location)

	transaction := a.conf.Database.Ctx().
		Select("id").
		Where("uuid=?", uuid).
		Where("type=?", int32(pb.LocationType_LOCATION_TYPE_COUNTRY)).
		First(country)

	if transaction.Error != nil && transaction.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, transaction.Error.Error())
	}

	if *a.location.ParentID == country.ID {
		a.location.ParentID = nil
	}

	return nil
}

func (a *AggregateRoot) RemoveRegion(uuid string) error {
	if err := utils.ValidateUuid(uuid); err != nil {
		return err
	}

	region := new(Location)

	result := a.conf.Database.Ctx().
		Select("id").
		Where("uuid=?", uuid).
		Where("type=?", int32(pb.LocationType_LOCATION_TYPE_REGION)).
		First(region)

	if result.Error != nil && result.Error.Error() == RecordNotFound {
		return status.Error(codes.NotFound, result.Error.Error())
	}

	if *a.location.ParentID == region.ID {
		a.location.ParentID = nil
	}

	return nil
}

func (a *AggregateRoot) AddChallenge(uuid string) error {
	if err := utils.ValidateUuid(uuid); err != nil {
		return err
	}

	request := new(pb.GetChallengeItemRequest)
	request.Uuid = uuid

	grpcPort := a.conf.Values.GrpcPort

	challenge, err := NewGrpcClient(grpcPort).GetChallengeItem(request)
	if err != nil {
		return status.Error(codes.NotFound, "Challenge not found >> "+err.Error())
	}

	challenges := new(LocationChallenges)
	challenges.ChallengeUUID, err = _uuid.Parse(challenge.Data.Uuid)
	if err != nil {
		return status.Error(codes.Internal, "Could not parse uuid to string >> "+err.Error())
	}

	a.location.Challenges = append(a.location.Challenges, challenges)

	return nil
}

func (a *AggregateRoot) RemoveChallenge(uuid string) error {
	return nil
}

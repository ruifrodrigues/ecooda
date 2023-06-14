package location

import (
	_uuid "github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"github.com/ruifrodrigues/ecooda/utils"
)

var Seeders = config.Seeds{
	&Location{}: seedLocationTable,
}

func seedLocationTable(db config.Database) {
	var c = []Location{
		{
			ID:   1,
			UUID: _uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Africa",
		},
		{
			ID:   2,
			UUID: _uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Asia",
		},
		{
			ID:   3,
			UUID: _uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Europe",
		},
		{
			ID:   4,
			UUID: _uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "North America",
		},
		{
			ID:   5,
			UUID: _uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "South America",
		},
		{
			ID:   6,
			UUID: _uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Oceania",
		},
		{
			ID:       7,
			ParentID: utils.Pointer[uint](3),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_COUNTRY),
			Name:     "Portugal",
		},
		{
			ID:       8,
			ParentID: utils.Pointer[uint](7),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_REGION),
			Name:     "Estremadura",
		},
		{
			ID:       9,
			ParentID: utils.Pointer[uint](8),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_CITY),
			Name:     "Lisbon",
		},
		{
			ID:       10,
			ParentID: utils.Pointer[uint](3),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_COUNTRY),
			Name:     "Netherlands",
		},
		{
			ID:       11,
			ParentID: utils.Pointer[uint](10),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_REGION),
			Name:     "South Holland",
		},
		{
			ID:       12,
			ParentID: utils.Pointer[uint](11),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_CITY),
			Name:     "Amsterdam",
		},
		{
			ID:       13,
			ParentID: utils.Pointer[uint](5),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_COUNTRY),
			Name:     "Brazil",
		},
		{
			ID:       14,
			ParentID: utils.Pointer[uint](13),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_REGION),
			Name:     "Southeast Region",
		},
		{
			ID:       15,
			ParentID: utils.Pointer[uint](14),
			UUID:     _uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_CITY),
			Name:     "Rio de Janeiro",
		},
	}

	db.Create(&c)
}

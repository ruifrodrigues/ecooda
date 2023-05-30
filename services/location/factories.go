package location

import (
	"github.com/google/uuid"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

var Seeders = config.Seeds{
	&Location{}: seedLocationTable,
}

func seedLocationTable(db config.Database) {
	var c = []Location{
		{
			ID:   1,
			UUID: uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Africa",
		},
		{
			ID:   2,
			UUID: uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Asia",
		},
		{
			ID:   3,
			UUID: uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Europe",
		},
		{
			ID:   4,
			UUID: uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "North America",
		},
		{
			ID:   5,
			UUID: uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "South America",
		},
		{
			ID:   6,
			UUID: uuid.New(),
			Type: int32(pb.LocationType_LOCATION_TYPE_CONTINENT),
			Name: "Oceania",
		},
		{
			ID:       7,
			ParentID: 3,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_COUNTRY),
			Name:     "Portugal",
		},
		{
			ID:       8,
			ParentID: 7,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_REGION),
			Name:     "Estremadura",
		},
		{
			ID:       9,
			ParentID: 8,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_CITY),
			Name:     "Lisbon",
		},
		{
			ID:       10,
			ParentID: 3,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_COUNTRY),
			Name:     "Netherlands",
		},
		{
			ID:       11,
			ParentID: 10,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_REGION),
			Name:     "South Holland",
		},
		{
			ID:       12,
			ParentID: 11,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_CITY),
			Name:     "Amsterdam",
		},
		{
			ID:       13,
			ParentID: 5,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_COUNTRY),
			Name:     "Brazil",
		},
		{
			ID:       14,
			ParentID: 13,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_REGION),
			Name:     "Southeast Region",
		},
		{
			ID:       15,
			ParentID: 14,
			UUID:     uuid.New(),
			Type:     int32(pb.LocationType_LOCATION_TYPE_CITY),
			Name:     "Rio de Janeiro",
		},
	}

	db.Create(&c)
}

package location

import (
	"context"
	"github.com/ruifrodrigues/ecooda/api"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

type (
	Fields   = api.Fields
	Config   = config.Config
	Database = config.Database
	Context  = context.Context

	PbLocation             = pb.Location
	PbGetLocationsRequest  = pb.GetLocationsRequest
	PbGetLocationsResponse = pb.GetLocationsResponse
	PbGetLocationRequest   = pb.GetLocationRequest
	PbGetLocationResponse  = pb.GetLocationResponse
	PbPostLocationRequest  = pb.PostLocationRequest
	PbPostLocationResponse = pb.PostLocationResponse
	PbLocationName         = pb.Location_Name
	PbLocationParents      = pb.Location_Parents
	PbLocationType         = pb.Location_Type
	PbParents              = pb.Parents
)

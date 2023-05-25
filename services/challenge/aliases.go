package challenge

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

	PbChallenge                = pb.Challenge
	PbGetChallengesRequest     = pb.GetChallengesRequest
	PbGetChallengesResponse    = pb.GetChallengesResponse
	PbGetChallengeRequest      = pb.GetChallengeRequest
	PbGetChallengeResponse     = pb.GetChallengeResponse
	PbPostChallengeRequest     = pb.PostChallengeRequest
	PbPostChallengeResponse    = pb.PostChallengeResponse
	PbPutChallengeRequest      = pb.PutChallengeRequest
	PbPutChallengeResponse     = pb.PutChallengeResponse
	PbDeleteChallengeRequest   = pb.DeleteChallengeRequest
	PbDeleteChallengeResponse  = pb.DeleteChallengeResponse
	PbCommandChallengeRequest  = pb.CommandChallengeRequest
	PbCommandChallengeResponse = pb.CommandChallengeResponse

	PbChallengeName        = pb.Challenge_Name
	PbChallengeDescription = pb.Challenge_Description
	PbChallengeStreet      = pb.Challenge_Street
	PbChallengePostcode    = pb.Challenge_Postcode
	PbChallengeLatitude    = pb.Challenge_Latitude
	PbChallengeLongitude   = pb.Challenge_Longitude
	PbChallengeThumbnail   = pb.Challenge_Thumbnail
	PbChallengeGallery     = pb.Challenge_Gallery
	PbChallengeCategories  = pb.Challenge_Categories

	PbCategory               = pb.Category
	PbCategoryName           = pb.Category_Name
	PbGetCategoriesRequest   = pb.GetCategoriesRequest
	PbGetCategoriesResponse  = pb.GetCategoriesResponse
	PbGetCategoryRequest     = pb.GetCategoryRequest
	PbGetCategoryResponse    = pb.GetCategoryResponse
	PbPostCategoryRequest    = pb.PostCategoryRequest
	PbPostCategoryResponse   = pb.PostCategoryResponse
	PbPutCategoryRequest     = pb.PutCategoryRequest
	PbPutCategoryResponse    = pb.PutCategoryResponse
	PbDeleteCategoryRequest  = pb.DeleteCategoryRequest
	PbDeleteCategoryResponse = pb.DeleteCategoryResponse
)

package challenge

import (
	"github.com/ruifrodrigues/ecooda/hateoas"
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	locationv1 "github.com/ruifrodrigues/ecooda/stubs/go/location/v1"
)

const MaxLimit = 25

var (
	DefaultFields = []string{
		"uuid",
		"created_at",
		"updated_at",
	}

	GuardedFields = []string{
		"id",
	}
)

func Fields(resource *hateoas.Resource, field string) {
	if field == "uuid" {
		resource.Challenge.Uuid = resource.Tmp["uuid"].(string)
	}

	if field == "name" {
		resource.Challenge.Name = resource.Tmp["name"].(string)
	}

	if field == "description" {
		description := ""
		if resource.Tmp["description"] != nil {
			description = resource.Tmp["description"].(string)
		}
		resource.Challenge.OptionalDescription = &challengev1.Challenge_Description{
			Description: description,
		}
	}

	if field == "street" {
		street := ""
		if resource.Tmp["street"] != nil {
			street = resource.Tmp["street"].(string)
		}

		resource.Challenge.OptionalStreet = &challengev1.Challenge_Street{
			Street: street,
		}
	}

	if field == "postcode" {
		postcode := ""
		if resource.Tmp["postcode"] != nil {
			postcode = resource.Tmp["postcode"].(string)
		}

		resource.Challenge.OptionalPostcode = &challengev1.Challenge_Postcode{
			Postcode: postcode,
		}
	}

	if field == "latitude" {
		var latitude float32 = 0.0
		if resource.Tmp["latitude"] != nil {
			latitude = resource.Tmp["latitude"].(float32)
		}

		resource.Challenge.OptionalLatitude = &challengev1.Challenge_Latitude{
			Latitude: latitude,
		}
	}

	if field == "longitude" {
		var longitude float32 = 0.0
		if resource.Tmp["longitude"] != nil {
			longitude = resource.Tmp["longitude"].(float32)
		}

		resource.Challenge.OptionalLongitude = &challengev1.Challenge_Longitude{
			Longitude: longitude,
		}
	}

	if field == "thumbnail" {
		thumbnail := ""
		if resource.Tmp["thumbnail"] != nil {
			thumbnail = resource.Tmp["thumbnail"].(string)
		}

		resource.Challenge.OptionalThumbnail = &challengev1.Challenge_Thumbnail{
			Thumbnail: thumbnail,
		}
	}

	if field == "gallery" {
		gallery := ""
		if resource.Tmp["gallery"] != nil {
			gallery = resource.Tmp["gallery"].(string)
		}

		resource.Challenge.OptionalGallery = &challengev1.Challenge_Gallery{
			Gallery: gallery,
		}
	}

	if field == "created_at" {
		resource.Challenge.CreatedAt = resource.Tmp["created_at"].(string)
	}

	if field == "updated_at" {
		resource.Challenge.UpdatedAt = resource.Tmp["updated_at"].(string)
	}
}

func Includes(resource *hateoas.Resource, includes []string) {
	if len(includes) == 0 {
		return
	}

	var includeList []*challengev1.Include

	for _, include := range includes {
		if include == "locations" {
			includeList = append(includeList, &challengev1.Include{
				OptionalLocation: &challengev1.Include_Location{
					Location: &locationv1.Location{
						Uuid:      "",
						Continent: "Europe",
						Country:   "Portugal",
						Region:    "Lisbon Metro Area",
						City:      "Lisbon",
					},
				},
			})
		}
	}

	resource.Challenge.Includes = includeList
}

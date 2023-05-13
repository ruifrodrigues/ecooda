package challenge

import (
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	locationv1 "github.com/ruifrodrigues/ecooda/stubs/go/location/v1"
	"gorm.io/gorm"
	"strings"
)

var defaultFields = "uuid,{requested_fields},created_at,updated_at"

func transformer(model *gorm.DB, req *challengev1.GetChallengesRequest) []*challengev1.Challenge {
	var tmpChallenges []map[string]interface{}

	fields := buildFields(req)

	model.Select(fields).Find(&tmpChallenges)

	items := collection(tmpChallenges)

	includes(items)

	return items
}

func buildFields(req *challengev1.GetChallengesRequest) string {
	var fields string
	if len(req.Fields) == 0 {
		fields = strings.Replace(defaultFields, ",{requested_fields}", "", 1)
	} else {
		fields = strings.Replace(defaultFields, "{requested_fields}", req.Fields, 1)
	}

	return fields
}

func collection(tmpChallenges []map[string]interface{}) []*challengev1.Challenge {
	var items []*challengev1.Challenge

	for _, tmpChallenge := range tmpChallenges {
		challenge := &challengev1.Challenge{}

		for fieldName, fieldValue := range tmpChallenge {
			item(challenge, fieldName, fieldValue)
		}

		items = append(items, challenge)
	}

	return items
}

func item(challenge *challengev1.Challenge, fieldName string, fieldValue interface{}) {
	if fieldName == "uuid" {
		challenge.Uuid = fieldValue.(string)
	}

	if fieldName == "name" {
		challenge.OptionalName = &challengev1.Challenge_Name{
			Name: fieldValue.(string),
		}
	}

	if fieldName == "description" {
		challenge.OptionalDescription = &challengev1.Challenge_Description{
			Description: fieldValue.(string),
		}
	}

	if fieldName == "street" {
		val := ""
		if fieldValue != nil {
			val = fieldValue.(string)
		}

		challenge.OptionalStreet = &challengev1.Challenge_Street{
			Street: val,
		}
	}

	if fieldName == "postcode" {
		val := ""
		if fieldValue != nil {
			val = fieldValue.(string)
		}

		challenge.OptionalPostcode = &challengev1.Challenge_Postcode{
			Postcode: val,
		}
	}

	if fieldName == "latitude" {
		var val float32 = 0.0
		if fieldValue != nil {
			val = fieldValue.(float32)
		}

		challenge.OptionalLatitude = &challengev1.Challenge_Latitude{
			Latitude: val,
		}
	}

	if fieldName == "longitude" {
		var val float32 = 0.0
		if fieldValue != nil {
			val = fieldValue.(float32)
		}

		challenge.OptionalLongitude = &challengev1.Challenge_Longitude{
			Longitude: val,
		}
	}

	if fieldName == "thumbnail" {
		val := ""
		if fieldValue != nil {
			val = fieldValue.(string)
		}

		challenge.OptionalThumbnail = &challengev1.Challenge_Thumbnail{
			Thumbnail: val,
		}
	}

	if fieldName == "gallery" {
		val := ""
		if fieldValue != nil {
			val = fieldValue.(string)
		}

		challenge.OptionalGallery = &challengev1.Challenge_Gallery{
			Gallery: val,
		}
	}

	if fieldName == "created_at" {
		challenge.CreatedAt = fieldValue.(string)
	}

	if fieldName == "updated_at" {
		challenge.UpdatedAt = fieldValue.(string)
	}
}

func includes(items []*challengev1.Challenge) {
	includes := []*challengev1.Include{
		{
			OptionalLocation: &challengev1.Include_Location{
				Location: &locationv1.Location{
					Uuid:      "",
					Continent: "Europe",
					Country:   "Portugal",
					Region:    "Lisbon Metro Area",
					City:      "Lisbon",
				},
			},
		},
	}

	// Add Includes
	for _, item := range items {
		item.Includes = includes
	}
}

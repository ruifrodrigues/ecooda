package challenge

import (
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"sync"
)

func challengeFields(records []*Challenge, fields []string, limit int32) []*PbChallenge {
	var data []*PbChallenge

	dataChan := make(chan *PbChallenge, limit)

	wg := sync.WaitGroup{}
	wg.Add(len(records))

	for _, record := range records {
		item := &PbChallenge{}

		go func() {
			for _, field := range fields {
				if field == "uuid" {
					item.Uuid = record.UUID.String()
				}

				if field == "name" {
					name := ""
					if record.Name != "" {
						name = record.Name
					}
					item.OptionalName = &PbChallengeName{
						Name: name,
					}
				}

				if field == "description" {
					description := ""
					if record.Description != "" {
						description = record.Description
					}
					item.OptionalDescription = &PbChallengeDescription{
						Description: description,
					}
				}

				if field == "street" {
					street := ""
					if record.Street != "" {
						street = record.Street
					}

					item.OptionalStreet = &PbChallengeStreet{
						Street: street,
					}
				}

				if field == "postcode" {
					postcode := ""
					if record.Postcode != "" {
						postcode = record.Postcode
					}

					item.OptionalPostcode = &PbChallengePostcode{
						Postcode: postcode,
					}
				}

				if field == "latitude" {
					var latitude float32 = 0
					if record.Latitude != 0 {
						latitude = record.Latitude
					}

					item.OptionalLatitude = &PbChallengeLatitude{
						Latitude: latitude,
					}
				}

				if field == "longitude" {
					var longitude float32 = 0
					if record.Longitude != 0 {
						longitude = record.Longitude
					}

					item.OptionalLongitude = &PbChallengeLongitude{
						Longitude: longitude,
					}
				}

				if field == "thumbnail" {
					thumbnail := ""
					if record.Thumbnail != "" {
						thumbnail = record.Thumbnail
					}

					item.OptionalThumbnail = &PbChallengeThumbnail{
						Thumbnail: thumbnail,
					}
				}

				if field == "gallery" {
					gallery := ""
					if record.Gallery != nil {
						gallery = record.Gallery.String()
					}

					item.OptionalGallery = &PbChallengeGallery{
						Gallery: gallery,
					}
				}

				if field == "created_at" {
					item.CreatedAt = record.CreatedAt.String()
				}

				if field == "updated_at" {
					item.UpdatedAt = record.UpdatedAt.String()
				}

				if field == "categories" {
					item.OptionalCategories = &PbChallengeCategories{
						Categories: &pb.Categories{
							Data: categoryFields(record.Categories, []string{"uuid", "name", "created_at", "updated_at"}, 25),
						},
					}
				}
			}

			dataChan <- item

			wg.Done()
		}()

		data = append(data, <-dataChan)
	}

	wg.Wait()

	return data
}

func categoryFields(records []*Category, fields []string, limit int32) []*PbCategory {
	var data []*PbCategory

	dataChan := make(chan *PbCategory, limit)

	wg := sync.WaitGroup{}
	wg.Add(len(records))

	for _, record := range records {
		item := &PbCategory{}

		go func() {
			for _, field := range fields {
				if field == "uuid" {
					item.Uuid = record.UUID.String()
				}

				if field == "name" {
					name := ""
					if record.Name != "" {
						name = record.Name
					}
					item.OptionalName = &PbCategoryName{
						Name: name,
					}
				}

				if field == "created_at" {
					item.CreatedAt = record.CreatedAt.String()
				}

				if field == "updated_at" {
					item.UpdatedAt = record.UpdatedAt.String()
				}
			}

			dataChan <- item

			wg.Done()
		}()

		data = append(data, <-dataChan)
	}

	wg.Wait()

	return data
}

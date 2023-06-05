package challenge

import (
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"sync"
)

func challengeFields(records []*Challenge, fields []string, limit int32) []*pb.Challenge {
	var data []*pb.Challenge

	dataChan := make(chan *pb.Challenge, limit)

	wg := sync.WaitGroup{}
	wg.Add(len(records))

	for _, record := range records {
		item := &pb.Challenge{}

		go func() {
			for _, field := range fields {
				if field == "uuid" {
					item.Uuid = record.UUID.String()
				}

				if field == "name" {
					item.OptionalName = locationName(item, record)
				}

				if field == "description" {
					item.OptionalDescription = locationDescription(item, record)
				}

				if field == "street" {
					item.OptionalStreet = locationStreet(item, record)
				}

				if field == "postcode" {
					item.OptionalPostcode = locationPostcode(item, record)
				}

				if field == "latitude" {
					item.OptionalLatitude = locationLatitude(item, record)
				}

				if field == "longitude" {
					item.OptionalLongitude = locationLongitude(item, record)
				}

				if field == "thumbnail" {
					item.OptionalThumbnail = locationThumbnail(item, record)
				}

				if field == "gallery" {
					item.OptionalGallery = locationGallery(item, record)
				}

				if field == "created_at" {
					item.CreatedAt = record.CreatedAt.String()
				}

				if field == "updated_at" {
					item.UpdatedAt = record.UpdatedAt.String()
				}

				if field == "categories" {
					item.OptionalCategories = locationCategories(item, record)
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

func locationName(item *pb.Challenge, record *Challenge) *pb.Challenge_Name {
	name := ""
	if record.Name != "" {
		name = record.Name
	}
	item.OptionalName = &pb.Challenge_Name{
		Name: name,
	}

	return item.OptionalName.(*pb.Challenge_Name)
}

func locationDescription(item *pb.Challenge, record *Challenge) *pb.Challenge_Description {
	description := ""
	if record.Description != "" {
		description = record.Description
	}

	item.OptionalDescription = &pb.Challenge_Description{
		Description: description,
	}

	return item.OptionalDescription.(*pb.Challenge_Description)
}

func locationStreet(item *pb.Challenge, record *Challenge) *pb.Challenge_Street {
	street := ""
	if record.Street != "" {
		street = record.Street
	}

	item.OptionalStreet = &pb.Challenge_Street{
		Street: street,
	}

	return item.OptionalStreet.(*pb.Challenge_Street)
}

func locationPostcode(item *pb.Challenge, record *Challenge) *pb.Challenge_Postcode {
	postcode := ""
	if record.Postcode != "" {
		postcode = record.Postcode
	}

	item.OptionalPostcode = &pb.Challenge_Postcode{
		Postcode: postcode,
	}

	return item.OptionalPostcode.(*pb.Challenge_Postcode)
}

func locationLatitude(item *pb.Challenge, record *Challenge) *pb.Challenge_Latitude {
	var latitude float32 = 0
	if record.Latitude != 0 {
		latitude = record.Latitude
	}

	item.OptionalLatitude = &pb.Challenge_Latitude{
		Latitude: latitude,
	}

	return item.OptionalLatitude.(*pb.Challenge_Latitude)
}

func locationLongitude(item *pb.Challenge, record *Challenge) *pb.Challenge_Longitude {
	var longitude float32 = 0
	if record.Longitude != 0 {
		longitude = record.Longitude
	}

	item.OptionalLongitude = &pb.Challenge_Longitude{
		Longitude: longitude,
	}

	return item.OptionalLongitude.(*pb.Challenge_Longitude)
}

func locationThumbnail(item *pb.Challenge, record *Challenge) *pb.Challenge_Thumbnail {
	thumbnail := ""
	if record.Thumbnail != "" {
		thumbnail = record.Thumbnail
	}

	item.OptionalThumbnail = &pb.Challenge_Thumbnail{
		Thumbnail: thumbnail,
	}

	return item.OptionalThumbnail.(*pb.Challenge_Thumbnail)
}

func locationGallery(item *pb.Challenge, record *Challenge) *pb.Challenge_Gallery {
	gallery := ""
	if record.Gallery != nil {
		gallery = record.Gallery.String()
	}

	item.OptionalGallery = &pb.Challenge_Gallery{
		Gallery: gallery,
	}

	return item.OptionalGallery.(*pb.Challenge_Gallery)
}

func locationCategories(item *pb.Challenge, record *Challenge) *pb.Challenge_Categories {
	item.OptionalCategories = &pb.Challenge_Categories{
		Categories: &pb.Categories{
			Data: categoryFields(record.Categories, []string{"uuid", "name", "created_at", "updated_at"}, 25),
		},
	}

	return item.OptionalCategories.(*pb.Challenge_Categories)
}

func categoryFields(records []*Category, fields []string, limit int32) []*pb.Category {
	var data []*pb.Category

	dataChan := make(chan *pb.Category, limit)

	wg := sync.WaitGroup{}
	wg.Add(len(records))

	for _, record := range records {
		item := &pb.Category{}

		go func() {
			for _, field := range fields {
				if field == "uuid" {
					item.Uuid = record.UUID.String()
				}

				if field == "name" {
					item.OptionalName = categoryName(item, record)
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

func categoryName(item *pb.Category, record *Category) *pb.Category_Name {
	name := ""
	if record.Name != "" {
		name = record.Name
	}

	item.OptionalName = &pb.Category_Name{
		Name: name,
	}

	return item.OptionalName.(*pb.Category_Name)
}

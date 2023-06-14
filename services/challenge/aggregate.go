package challenge

import (
	"encoding/json"
	"github.com/ruifrodrigues/ecooda/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/datatypes"
)

type Aggregate interface {
	ChangeName(name string)
	ChangeDescription(description string)
	ChangeStreet(street string)
	ChangePostcode(postcode string)
	ChangeLatitude(latitude float32)
	ChangeLongitude(longitude float32)
	ChangeThumbnail(thumbnail string)
	ChangeGallery(gallery string) error
	AddCategory(uuid string) error
	RemoveCategory(uuid string) error
}

type AggregateRoot struct {
	conf      config.Config
	challenge *Challenge
}

func NewAggregateRoot(conf config.Config, challenge *Challenge) Aggregate {
	aggregateRoot := &AggregateRoot{
		conf:      conf,
		challenge: challenge,
	}

	return aggregateRoot
}

func (a *AggregateRoot) ChangeName(name string) {
	if name != "" {
		a.challenge.Name = name
	}
}

func (a *AggregateRoot) ChangeDescription(description string) {
	if description != "" {
		a.challenge.Description = description
	}
}

func (a *AggregateRoot) ChangeStreet(street string) {
	if street != "" {
		a.challenge.Street = street
	}
}

func (a *AggregateRoot) ChangePostcode(postcode string) {
	if postcode != "" {
		a.challenge.Postcode = postcode
	}
}

func (a *AggregateRoot) ChangeLatitude(latitude float32) {
	if latitude != 0 {
		a.challenge.Latitude = latitude
	}
}

func (a *AggregateRoot) ChangeLongitude(longitude float32) {
	if longitude != 0 {
		a.challenge.Longitude = longitude
	}
}

func (a *AggregateRoot) ChangeThumbnail(thumbnail string) {
	if thumbnail != "" {
		a.challenge.Thumbnail = thumbnail
	}
}

func (a *AggregateRoot) ChangeGallery(gallery string) error {
	if gallery != "" {
		var galleryJSON datatypes.JSON

		err := json.Unmarshal([]byte(gallery), &galleryJSON)
		if err != nil {
			return status.Error(codes.InvalidArgument, "Gallery field not a valid JSON >> "+err.Error())
		}

		a.challenge.Gallery = galleryJSON
	}

	return nil
}

func (a *AggregateRoot) AddCategory(uuid string) error {
	category, err := NewQuery(a.conf).GetCategoryByUuid(uuid)
	if err != nil {
		return err
	}

	a.challenge.Categories = append(a.challenge.Categories, category)

	return nil
}

func (a *AggregateRoot) RemoveCategory(uuid string) error {
	categories := a.challenge.Categories

	for i, category := range categories {
		if category.UUID.String() == uuid {
			categories[i] = categories[len(categories)-1]
			a.challenge.Categories = categories[:len(categories)-1]
		}
	}

	return nil
}

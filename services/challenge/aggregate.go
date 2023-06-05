package challenge

import (
	"encoding/json"
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
	AddCategory(category *Category)
	RemoveCategory(uuid string)
}

type AggregateRoot struct {
	challenge *Challenge
}

func NewAggregateRoot(challenge *Challenge) Aggregate {
	return &AggregateRoot{
		challenge: challenge,
	}
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

func (a *AggregateRoot) AddCategory(category *Category) {
	a.challenge.Categories = append(a.challenge.Categories, category)
}

func (a *AggregateRoot) RemoveCategory(uuid string) {
	categories := a.challenge.Categories

	for i, category := range categories {
		if category.UUID.String() == uuid {
			categories[i] = categories[len(categories)-1]
			a.challenge.Categories = categories[:len(categories)-1]
		}
	}
}

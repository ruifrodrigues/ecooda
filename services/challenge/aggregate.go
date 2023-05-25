package challenge

import (
	"gorm.io/datatypes"
)

type AggregateRoot struct {
	challenge *Challenge
}

type Aggregate interface {
	ChangeName(name string)
	ChangeDescription(description string)
	ChangeStreet(street string)
	ChangePostcode(postcode string)
	ChangeLatitude(latitude float32)
	ChangeLongitude(longitude float32)
	ChangeThumbnail(thumbnail string)
	ChangeGallery(gallery datatypes.JSON)
	AddCategory(category *Category)
	RemoveCategory(uuid string)
}

func NewAggregateRoot(challenge *Challenge) *AggregateRoot {
	return &AggregateRoot{
		challenge: challenge,
	}
}

func (a *AggregateRoot) ChangeName(name string) {
	a.challenge.Name = name
}

func (a *AggregateRoot) ChangeDescription(description string) {
	a.challenge.Description = description
}

func (a *AggregateRoot) ChangeStreet(street string) {
	a.challenge.Street = street
}

func (a *AggregateRoot) ChangePostcode(postcode string) {
	a.challenge.Postcode = postcode
}

func (a *AggregateRoot) ChangeLatitude(latitude float32) {
	a.challenge.Latitude = latitude
}

func (a *AggregateRoot) ChangeLongitude(longitude float32) {
	a.challenge.Longitude = longitude
}

func (a *AggregateRoot) ChangeThumbnail(thumbnail string) {
	a.challenge.Thumbnail = thumbnail
}

func (a *AggregateRoot) ChangeGallery(gallery datatypes.JSON) {
	a.challenge.Gallery = gallery
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

package challenge

import (
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"sync"
)

type InvokeChallengeFn func(item *pb.Challenge, record *Challenge)

type ChallengeData struct {
	conf    config.Config
	records []*Challenge
	fields  []string
	limit   int32
	invoke  map[string]InvokeChallengeFn
}

func NewChallengeData(conf config.Config, records []*Challenge, fields []string, limit int32) *ChallengeData {
	data := &ChallengeData{
		conf:    conf,
		records: records,
		fields:  fields,
		limit:   limit,
	}

	data.invoke = map[string]InvokeChallengeFn{
		"uuid":        data.uuid,
		"name":        data.name,
		"description": data.description,
		"street":      data.street,
		"postcode":    data.postcode,
		"latitude":    data.latitude,
		"longitude":   data.longitude,
		"thumbnail":   data.thumbnail,
		"gallery":     data.gallery,
		"created_at":  data.createdAt,
		"updated_at":  data.updatedAt,
		"categories":  data.categories,
		"locations":   data.locations,
	}

	return data
}

func (d *ChallengeData) Generate() []*pb.Challenge {
	var data []*pb.Challenge

	if len(d.records) == 0 {
		return data
	}

	dataChan := make(chan *pb.Challenge, d.limit)

	wg := sync.WaitGroup{}
	wg.Add(len(d.records))

	for _, record := range d.records {
		item := &pb.Challenge{}

		go func() {
			for _, field := range d.fields {
				d.invoke[field](item, record)
			}

			dataChan <- item

			wg.Done()
		}()

		data = append(data, <-dataChan)
	}

	wg.Wait()

	return data
}

func (d *ChallengeData) uuid(item *pb.Challenge, record *Challenge) {
	item.Uuid = record.UUID.String()
}

func (d *ChallengeData) name(item *pb.Challenge, record *Challenge) {
	name := ""
	if record.Name != "" {
		name = record.Name
	}

	item.OptionalName = &pb.Challenge_Name{
		Name: name,
	}
}

func (d *ChallengeData) description(item *pb.Challenge, record *Challenge) {
	description := ""
	if record.Description != "" {
		description = record.Description
	}

	item.OptionalDescription = &pb.Challenge_Description{
		Description: description,
	}
}

func (d *ChallengeData) street(item *pb.Challenge, record *Challenge) {
	street := ""
	if record.Street != "" {
		street = record.Street
	}

	item.OptionalStreet = &pb.Challenge_Street{
		Street: street,
	}
}

func (d *ChallengeData) postcode(item *pb.Challenge, record *Challenge) {
	postcode := ""
	if record.Postcode != "" {
		postcode = record.Postcode
	}

	item.OptionalPostcode = &pb.Challenge_Postcode{
		Postcode: postcode,
	}
}

func (d *ChallengeData) latitude(item *pb.Challenge, record *Challenge) {
	var latitude float32 = 0
	if record.Latitude != 0 {
		latitude = record.Latitude
	}

	item.OptionalLatitude = &pb.Challenge_Latitude{
		Latitude: latitude,
	}
}

func (d *ChallengeData) longitude(item *pb.Challenge, record *Challenge) {
	var longitude float32 = 0
	if record.Longitude != 0 {
		longitude = record.Longitude
	}

	item.OptionalLongitude = &pb.Challenge_Longitude{
		Longitude: longitude,
	}
}

func (d *ChallengeData) thumbnail(item *pb.Challenge, record *Challenge) {
	thumbnail := ""
	if record.Thumbnail != "" {
		thumbnail = record.Thumbnail
	}

	item.OptionalThumbnail = &pb.Challenge_Thumbnail{
		Thumbnail: thumbnail,
	}
}

func (d *ChallengeData) gallery(item *pb.Challenge, record *Challenge) {
	gallery := ""
	if record.Gallery != nil {
		gallery = record.Gallery.String()
	}

	item.OptionalGallery = &pb.Challenge_Gallery{
		Gallery: gallery,
	}
}

func (d *ChallengeData) createdAt(item *pb.Challenge, record *Challenge) {
	item.CreatedAt = record.CreatedAt.String()
}

func (d *ChallengeData) updatedAt(item *pb.Challenge, record *Challenge) {
	item.UpdatedAt = record.UpdatedAt.String()
}

func (d *ChallengeData) categories(item *pb.Challenge, record *Challenge) {
	item.OptionalCategories = &pb.Challenge_Categories{
		Categories: &pb.Categories{
			Data: NewCategoryData(
				d.conf,
				record.Categories,
				[]string{"uuid", "name", "created_at", "updated_at"},
				25,
			).Generate(),
		},
	}
}

func (d *ChallengeData) locations(item *pb.Challenge, record *Challenge) {
	request := &pb.GetLocationFromChallengeRequest{
		Uuid: record.UUID.String(),
		OptionalFields: &pb.GetLocationFromChallengeRequest_Fields{
			Fields: "uuid,name,type,parents,created_at,updated_at",
		},
	}

	location := NewGrpcClient(":50051").GetLocationFromChallenge(request)

	item.OptionalLocations = &pb.Challenge_Locations{
		Locations: &pb.Locations{
			Data: location.Data,
		},
	}
}

type InvokeCategoryFn func(item *pb.Category, record *Category)

type CategoryData struct {
	conf    config.Config
	records []*Category
	fields  []string
	limit   int32
	invoke  map[string]InvokeCategoryFn
}

func NewCategoryData(conf config.Config, records []*Category, fields []string, limit int32) *CategoryData {
	data := &CategoryData{
		conf:    conf,
		records: records,
		fields:  fields,
		limit:   limit,
	}

	data.invoke = map[string]InvokeCategoryFn{
		"uuid":       data.uuid,
		"name":       data.name,
		"created_at": data.createdAt,
		"updated_at": data.updatedAt,
	}

	return data
}

func (d *CategoryData) Generate() []*pb.Category {
	var data []*pb.Category

	if len(d.records) == 0 {
		return data
	}

	dataChan := make(chan *pb.Category, d.limit)

	wg := sync.WaitGroup{}
	wg.Add(len(d.records))

	for _, record := range d.records {
		item := &pb.Category{}

		go func() {
			for _, field := range d.fields {
				d.invoke[field](item, record)
			}

			dataChan <- item

			wg.Done()
		}()

		data = append(data, <-dataChan)
	}

	wg.Wait()

	return data
}

func (d *CategoryData) uuid(item *pb.Category, record *Category) {
	item.Uuid = record.UUID.String()
}

func (d *CategoryData) name(item *pb.Category, record *Category) {
	name := ""
	if record.Name != "" {
		name = record.Name
	}

	item.OptionalName = &pb.Category_Name{
		Name: name,
	}
}

func (d *CategoryData) createdAt(item *pb.Category, record *Category) {
	item.CreatedAt = record.CreatedAt.String()
}

func (d *CategoryData) updatedAt(item *pb.Category, record *Category) {
	item.UpdatedAt = record.UpdatedAt.String()
}

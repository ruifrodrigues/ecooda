package location

import (
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"sync"
)

func locationFields(dbCtx Database, records []*Location, fields []string, limit int32) []*PbLocation {
	var data []*PbLocation

	dataChan := make(chan *PbLocation, limit)

	wg := sync.WaitGroup{}
	wg.Add(len(records))

	for _, record := range records {
		item := &PbLocation{}

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
					item.OptionalName = &PbLocationName{
						Name: name,
					}
				}

				if field == "type" {
					var t int32 = 0
					if record.Type != 0 {
						t = record.Type
					}
					item.OptionalType = &PbLocationType{
						Type: pb.LocationType(t),
					}
				}

				if field == "created_at" {
					item.CreatedAt = record.CreatedAt.String()
				}

				if field == "updated_at" {
					item.UpdatedAt = record.UpdatedAt.String()
				}

				if field == "parent" {
					if record.ParentID != 0 {
						location := &Location{}
						dbCtx.Where(record.ParentID).First(&location)
						item.OptionalParent = parentLocation(dbCtx, location)
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

func parentLocation(dbCtx Database, l *Location) *PbLocationParent {
	location := &PbLocationParent{
		Parent: &pb.Parent{
			Data: &PbLocation{
				Uuid: l.UUID.String(),
				OptionalName: &PbLocationName{
					Name: l.Name,
				},
				OptionalType: &PbLocationType{
					Type: pb.LocationType(l.Type),
				},
				CreatedAt: l.CreatedAt.String(),
				UpdatedAt: l.UpdatedAt.String(),
			},
		},
	}

	if l.ParentID != 0 {
		p := &Location{}
		result := dbCtx.Where(l.ParentID).First(&p)
		if result.Error == nil {
			location.Parent.Data.OptionalParent = parentLocation(dbCtx, p)
		}
	}

	return location
}

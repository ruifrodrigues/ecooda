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

				if field == "parents" {
					var location *Location
					pbParents := &PbParents{Data: nil}
					item.OptionalParents = &PbLocationParents{}

					if record.ParentID != 0 {
						dbCtx.Where(record.ParentID).First(&location)
						parents(dbCtx, location, pbParents)

						item.OptionalParents = &PbLocationParents{
							Parents: pbParents,
						}
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

func parents(dbCtx Database, record *Location, pbParents *PbParents) {
	pbLocation := &PbLocation{
		Uuid: record.UUID.String(),
		OptionalName: &PbLocationName{
			Name: record.Name,
		},
		OptionalType: &PbLocationType{
			Type: pb.LocationType(record.Type),
		},
		CreatedAt: record.CreatedAt.String(),
		UpdatedAt: record.UpdatedAt.String(),
	}

	pbParents.Data = append(pbParents.Data, pbLocation)

	if record.ParentID == 0 {
		return
	}

	var location *Location
	result := dbCtx.Where(record.ParentID).First(&location)
	if result.Error == nil {
		parents(dbCtx, location, pbParents)
	}
}

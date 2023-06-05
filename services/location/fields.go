package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"sync"
)

func locationFields(dbCtx config.Database, records []*Location, fields []string, limit int32) []*pb.Location {
	var data []*pb.Location

	dataChan := make(chan *pb.Location, limit)

	wg := sync.WaitGroup{}
	wg.Add(len(records))

	for _, record := range records {
		item := &pb.Location{}

		go func() {
			for _, field := range fields {
				if field == "uuid" {
					item.Uuid = record.UUID.String()
				}

				if field == "name" {
					item.OptionalName = locationName(item, record)
				}

				if field == "type" {
					item.OptionalType = locationType(item, record)
				}

				if field == "created_at" {
					item.CreatedAt = record.CreatedAt.String()
				}

				if field == "updated_at" {
					item.UpdatedAt = record.UpdatedAt.String()
				}

				if field == "parents" {
					item.OptionalParents = locationsParents(dbCtx, item, record)
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

func locationName(item *pb.Location, record *Location) *pb.Location_Name {
	name := ""
	if record.Name != "" {
		name = record.Name
	}

	item.OptionalName = &pb.Location_Name{
		Name: name,
	}

	return item.OptionalName.(*pb.Location_Name)
}

func locationType(item *pb.Location, record *Location) *pb.Location_Type {
	var t int32 = 0
	if record.Type != 0 {
		t = record.Type
	}

	item.OptionalType = &pb.Location_Type{
		Type: pb.LocationType(t),
	}

	return item.OptionalType.(*pb.Location_Type)
}

func locationsParents(dbCtx config.Database, item *pb.Location, record *Location) *pb.Location_Parents {
	var dataFn func(dbCtx config.Database, record *Location, pbParents *pb.Parents)

	dataFn = func(dbCtx config.Database, record *Location, pbParents *pb.Parents) {
		pbLocation := &pb.Location{
			Uuid: record.UUID.String(),
			OptionalName: &pb.Location_Name{
				Name: record.Name,
			},
			OptionalType: &pb.Location_Type{
				Type: pb.LocationType(record.Type),
			},
			CreatedAt: record.CreatedAt.String(),
			UpdatedAt: record.UpdatedAt.String(),
		}

		pbParents.Data = append(pbParents.Data, pbLocation)

		if record.ParentID == nil {
			return
		}

		var l *Location

		result := dbCtx.Where(record.ParentID).First(&l)
		if result.Error == nil {
			dataFn(dbCtx, l, pbParents)
		}
	}

	var l *Location
	pbParents := &pb.Parents{Data: nil}
	item.OptionalParents = &pb.Location_Parents{}

	if record.ParentID != nil {
		dbCtx.Where(record.ParentID).First(&l)
		dataFn(dbCtx, l, pbParents)

		item.OptionalParents = &pb.Location_Parents{
			Parents: pbParents,
		}
	}

	return item.OptionalParents.(*pb.Location_Parents)
}

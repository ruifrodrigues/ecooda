package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"sync"
)

type InvokeFn func(conf config.Config, item *pb.Location, record *Location)

var Invoke = map[string]InvokeFn{
	"uuid":       uuid,
	"name":       name,
	"type":       lType,
	"created_at": createdAt,
	"updated_at": updatedAt,
	"parents":    parents,
}

type Data struct {
	conf    config.Config
	records []*Location
	fields  []string
	limit   int32
}

func NewData(conf config.Config, records []*Location, fields []string, limit int32) *Data {
	return &Data{
		conf:    conf,
		records: records,
		fields:  fields,
		limit:   limit,
	}
}

func (d *Data) Generate() []*pb.Location {
	var data []*pb.Location

	if len(d.records) == 0 {
		return data
	}

	dataChan := make(chan *pb.Location, d.limit)

	wg := sync.WaitGroup{}
	wg.Add(len(d.records))

	for _, record := range d.records {
		item := &pb.Location{}

		go func() {
			for _, field := range d.fields {
				Invoke[field](d.conf, item, record)
			}

			dataChan <- item

			wg.Done()
		}()

		data = append(data, <-dataChan)
	}

	wg.Wait()

	return data
}

func uuid(_ config.Config, item *pb.Location, record *Location) {
	item.Uuid = record.UUID.String()
}

func name(_ config.Config, item *pb.Location, record *Location) {
	name := ""
	if record.Name != "" {
		name = record.Name
	}

	item.OptionalName = &pb.Location_Name{
		Name: name,
	}
}

func lType(_ config.Config, item *pb.Location, record *Location) {
	var t int32 = 0
	if record.Type != 0 {
		t = record.Type
	}

	item.OptionalType = &pb.Location_Type{
		Type: pb.LocationType(t),
	}
}

func createdAt(_ config.Config, item *pb.Location, record *Location) {
	item.CreatedAt = record.CreatedAt.String()
}

func updatedAt(_ config.Config, item *pb.Location, record *Location) {
	item.UpdatedAt = record.UpdatedAt.String()
}

func parents(conf config.Config, item *pb.Location, record *Location) {
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

	dbCtx := conf.Database.Ctx()

	if record.ParentID != nil {
		dbCtx.Where(record.ParentID).First(&l)

		dataFn(dbCtx, l, pbParents)

		item.OptionalParents = &pb.Location_Parents{
			Parents: pbParents,
		}
	}
}

package api

import (
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"strings"
)

type Meta struct {
	*pb.Meta
}

func NewMeta() *Meta {
	return &Meta{
		&pb.Meta{},
	}
}

func (m *Meta) AddCursor(cursor *Cursor) *Meta {
	m.OptionalCursor = &pb.Meta_Cursor{
		Cursor: cursor.Cursor,
	}

	return m
}

func (m *Meta) AddFields(fields []string) *Meta {
	f := strings.Join(fields, ",")

	m.OptionalFields = &pb.Meta_Fields{
		Fields: f,
	}

	return m
}

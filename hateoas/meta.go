package hateoas

import (
	metav1 "github.com/ruifrodrigues/ecooda/stubs/go/meta/v1"
)

type Meta struct {
	*metav1.Meta
}

func NewMeta() *Meta {
	return &Meta{
		&metav1.Meta{},
	}
}

func (m *Meta) AddCursor(cursor *Cursor) *Meta {
	m.OptionalCursor = &metav1.Meta_Cursor{
		Cursor: cursor.Cursor,
	}

	return m
}

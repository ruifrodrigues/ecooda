package config

import (
	linkv1 "github.com/ruifrodrigues/ecooda/stubs/go/link/v1"
	"strings"
)

type Link struct {
	uuid string
	list []*linkv1.Link
}

func NewLink(uuid string, LinkList []*linkv1.Link) *Link {
	return &Link{
		uuid,
		LinkList,
	}
}

func (l *Link) List() []*linkv1.Link {
	for _, list := range l.list {
		list.Uri = strings.Replace(list.Uri, "{uuid}", l.uuid, 1)
	}

	return l.list
}

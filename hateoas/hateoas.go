package hateoas

import pb "github.com/ruifrodrigues/ecooda/stubs/go/meta/v1"

type Hateoas struct {
	metadata *pb.Meta
}

func New() *Hateoas {
	return &Hateoas{}
}

func (h *Hateoas) AddLinks(links []*pb.Link) {
	h.metadata.Links = links
}

func (h *Hateoas) AddCursor(cursor *pb.Cursor) {
	h.metadata.Cursor = cursor
}

func (h *Hateoas) Metadata() *pb.Meta {
	return h.metadata
}

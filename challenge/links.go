package challenge

import pb "github.com/ruifrodrigues/ecooda/stubs/go/meta/v1"

var links = []*pb.Link{
	{
		Rel:    "self",
		Uri:    "/challenges",
		Method: pb.Method_METHOD_POST,
	},
	{
		Rel:    "self",
		Uri:    "/challenges/{uuid}",
		Method: pb.Method_METHOD_GET,
	},
	{
		Rel:    "self",
		Uri:    "/challenges/{uuid}",
		Method: pb.Method_METHOD_PUT,
	},
	{
		Rel:    "self",
		Uri:    "/challenges/{uuid}",
		Method: pb.Method_METHOD_DELETE,
	},
}

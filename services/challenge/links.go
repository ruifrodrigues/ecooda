package challenge

import (
	linkv1 "github.com/ruifrodrigues/ecooda/stubs/go/link/v1"
)

var LinkList = []*linkv1.Link{
	{
		Rel:    "self",
		Uri:    "/challenges",
		Method: "POST",
		OptionalDescription: &linkv1.Link_Description{
			Description: "Store a new challenge",
		},
	},
	{
		Rel:    "self",
		Uri:    "/challenges/{uuid}",
		Method: "GET",
		OptionalDescription: &linkv1.Link_Description{
			Description: "Get a challenge by uuid",
		},
	},
	{
		Rel:    "self",
		Uri:    "/challenges/{uuid}",
		Method: "PUT",
		OptionalDescription: &linkv1.Link_Description{
			Description: "Update a challenge by uuid",
		},
	},
	{
		Rel:    "self",
		Uri:    "/challenges/{uuid}",
		Method: "DELETE",
		OptionalDescription: &linkv1.Link_Description{
			Description: "Delete a challenge by uuid",
		},
	},
}

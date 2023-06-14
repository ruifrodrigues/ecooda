package api

import (
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"math"
)

type Cursor struct {
	*pb.Cursor
}

func NewCursor(count, page, pageSize, maxLimit int32) *Cursor {
	cursor := &Cursor{
		&pb.Cursor{},
	}

	cursor.Count = count
	cursor.Previous = generatePrevious(page)
	cursor.Current = generateCurrent(page)
	cursor.Next = generateNext(cursor.Count, cursor.Current, pageSize, maxLimit)

	return cursor
}

func generatePrevious(page int32) int32 {
	currentPage := int(page)

	if currentPage <= 0 {
		return 0
	}

	return int32(currentPage - 1)
}

func generateCurrent(page int32) int32 {
	if page == 0 {
		return 1
	}

	return page
}

func generateNext(count, page, pageSize, maxLimit int32) int32 {
	size := size(pageSize, maxLimit)
	currentPage := int(page)
	nextPage := currentPage + 1
	maxPages := int(math.Ceil(float64(count) / float64(size)))

	if currentPage >= maxPages {
		return 0
	}

	return int32(nextPage)
}

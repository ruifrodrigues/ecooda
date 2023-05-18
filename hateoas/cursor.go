package hateoas

import (
	cursorv1 "github.com/ruifrodrigues/ecooda/stubs/go/cursor/v1"
	"gorm.io/gorm"
	"math"
)

type Cursor struct {
	*cursorv1.Cursor
}

func NewCursor[T Bag](model *gorm.DB, req *RequestBag[T]) *Cursor {
	cursor := &Cursor{
		&cursorv1.Cursor{},
	}
	cursor.Count = generateCount(model)
	cursor.Previous = generatePrevious(req)
	cursor.Current = generateCurrent(req)
	cursor.Next = generateNext(cursor.Count, req)

	return cursor
}

func generateCount(model *gorm.DB) int32 {
	var count int64 = 0
	model.Count(&count)

	return int32(count)
}

func generatePrevious[T Bag](req *RequestBag[T]) int32 {

	currentPage := int(req.GetPage())

	if currentPage <= 0 {
		return 0
	}

	return int32(currentPage - 1)
}

func generateCurrent[T Bag](req *RequestBag[T]) int32 {
	return req.GetPage()
}

func generateNext[T Bag](count int32, req *RequestBag[T]) int32 {
	currentPage := int(req.GetPage())
	nextPage := int(math.Ceil(float64(count) / float64(req.GetPageSize())))

	if currentPage == nextPage {
		return 0
	}

	if currentPage > nextPage {
		return int32(currentPage)
	}

	return int32(currentPage + 1)
}

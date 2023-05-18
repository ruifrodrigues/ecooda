package hateoas

import (
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type Resource struct {
	Tmp       map[string]interface{}
	Challenge *challengev1.Challenge
}

func NewResource() *Resource {
	return &Resource{
		Challenge: &challengev1.Challenge{},
	}
}

type Fields struct {
	Default []string
	Guarded []string
}

type IncludesFunc func(resource *Resource, includes []string)

type FieldsFunc func(resource *Resource, field string)

func NewFields(defaultFields, guardedFields []string) Fields {
	return Fields{
		Default: defaultFields,
		Guarded: guardedFields,
	}
}

type Transformer struct {
	Model           *gorm.DB
	RequestedFields []string
	Includes        []string
	Offset          int
	Limit           int
}

func NewTransformer[T Bag](model *gorm.DB, req *RequestBag[T], fields Fields, maxLimit int) *Transformer {
	return &Transformer{
		Model:           model,
		RequestedFields: requestedFields(req, fields),
		Includes:        includes(req),
		Offset:          offset(req),
		Limit:           limit(req, maxLimit),
	}
}

func requestedFields[T Bag](req *RequestBag[T], fb Fields) []string {
	if len(req.GetFields()) == 0 {
		return fb.Default
	}

	fieldsSlice := strings.Split(req.GetFields(), ",")
	for _, f := range fieldsSlice {
		for _, g := range fb.Guarded {
			if f == g {
				panic("field " + f + " not allowed")
			}
		}
	}

	fields := fb.Default
	fields = append(fields, strings.Split(req.GetFields(), ",")...)

	return fields
}

func includes[T Bag](req *RequestBag[T]) []string {
	return strings.Split(req.GetIncludes(), "|")
}

func limit[T Bag](req *RequestBag[T], maxLimit int) int {
	pageSize := int(req.GetPageSize())

	if pageSize > maxLimit {
		return maxLimit
	}

	if req.GetPageSize() == 0 {
		return maxLimit
	}

	return pageSize
}

func offset[T Bag](req *RequestBag[T]) int {
	currentPage := 0

	if req.GetPage() > 1 {
		currentPage = int(req.GetPage()) - 1
	}

	return currentPage * int(req.GetPageSize())
}

func Item(t *Transformer, fieldsFn FieldsFunc, includesFn IncludesFunc) *Resource {
	return &Resource{}
}

func Collection(t *Transformer, fieldsFn FieldsFunc, includesFn IncludesFunc) []*Resource {
	var (
		tmpCollection      []map[string]interface{}
		resourceCollection []*Resource
		resourceChan       = make(chan *Resource, t.Limit)
	)

	t.Model.Select("*").Offset(t.Offset).Limit(t.Limit).Find(&tmpCollection)

	wg := sync.WaitGroup{}
	wg.Add(len(tmpCollection))

	for _, tmpItem := range tmpCollection {
		resource := NewResource()
		resource.Tmp = tmpItem

		go func() {
			for _, field := range t.RequestedFields {
				fieldsFn(resource, field)
				includesFn(resource, t.Includes)
			}

			resourceChan <- resource

			wg.Done()
		}()

		resourceCollection = append(resourceCollection, <-resourceChan)
	}

	wg.Wait()

	return resourceCollection
}

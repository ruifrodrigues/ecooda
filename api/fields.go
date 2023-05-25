package api

import (
	"strings"
)

type FieldTypes struct {
	Guarded   []string
	Default   []string
	Available []string
}

type Fields struct {
	Challenge *FieldTypes
	Category  *FieldTypes
}

func NewFields() Fields {
	return Fields{}
}

func RequestedFields[T Bag](req *RequestBag[T], fb *FieldTypes) []string {
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

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
	Location  *FieldTypes
}

func NewFields() Fields {
	return Fields{}
}

func RequestedFields(fields string, fb *FieldTypes) []string {
	if len(fields) == 0 {
		return fb.Default
	}

	for _, f := range strings.Split(fields, ",") {
		for _, g := range fb.Guarded {
			if f == g {
				panic("field " + f + " not allowed")
			}
		}
	}

	f := fb.Default
	f = append(f, strings.Split(fields, ",")...)

	return f
}

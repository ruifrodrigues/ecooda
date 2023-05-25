package api

import (
	"fmt"
	"github.com/ruifrodrigues/ecooda/utils"
	"strings"
)

func Sort[T Bag](req *RequestBag[T], allowedSorts []string) string {
	if req.GetSort() == "" {
		return "id ASC"
	}

	field := strings.Replace(req.GetSort(), "-", "", 1)

	if !utils.InArray(field, allowedSorts) {
		return "id ASC"
	}

	if strings.HasPrefix(req.GetSort(), "-") {
		return fmt.Sprintf("%s DESC", field)
	}

	return fmt.Sprintf(" %s ASC", field)
}

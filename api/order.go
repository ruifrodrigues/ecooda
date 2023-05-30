package api

import (
	"fmt"
	"github.com/ruifrodrigues/ecooda/utils"
	"strings"
)

func Sort(sort string, allowedSorts []string) string {
	if sort == "" {
		return "id ASC"
	}

	field := strings.Replace(sort, "-", "", 1)

	if !utils.InArray(field, allowedSorts) {
		return "id ASC"
	}

	if strings.HasPrefix(sort, "-") {
		return fmt.Sprintf("%s DESC", field)
	}

	return fmt.Sprintf(" %s ASC", field)
}

package api

func size(pageSize, maxLimit int32) int {
	if pageSize == 0 {
		pageSize = maxLimit
	}

	return int(pageSize)
}

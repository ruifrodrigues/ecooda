package api

func Offset(page, pageSize, maxLimit int32) int {
	size := size(pageSize, maxLimit)

	if page == 1 {
		return 0
	}

	return int(page-1) * size
}

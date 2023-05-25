package api

func Limit[T Bag](req *RequestBag[T], maxLimit int) int {
	pageSize := int(req.GetPageSize())

	if pageSize > maxLimit {
		return maxLimit
	}

	if req.GetPageSize() == 0 {
		return maxLimit
	}

	return pageSize
}

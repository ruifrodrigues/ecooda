package api

func Limit(pageSize, maxLimit int32) int {
	pSize := int(pageSize)
	mLimit := int(maxLimit)

	if pSize > mLimit {
		return mLimit
	}

	if pSize == 0 {
		return mLimit
	}

	return pSize
}

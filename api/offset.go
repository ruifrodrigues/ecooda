package api

func Offset[T Bag](req *RequestBag[T]) int {
	currentPage := 0

	if req.GetPage() > 1 {
		currentPage = int(req.GetPage()) - 1
	}

	return currentPage * int(req.GetPageSize())
}

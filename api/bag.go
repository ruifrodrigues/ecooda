package api

import (
	ecoodav1 "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
)

type ChallengeItemRequest struct {
	Req *ecoodav1.GetChallengeRequest
}

func NewChallengeRequest(req *ecoodav1.GetChallengeRequest) *ChallengeItemRequest {
	return &ChallengeItemRequest{
		Req: req,
	}
}

func (cib *ChallengeItemRequest) GetFields() string {
	return cib.Req.GetFields()
}

func (cib *ChallengeItemRequest) GetSort() string {
	return ""
}

func (cib *ChallengeItemRequest) GetPage() int32 {
	return 0
}

func (cib *ChallengeItemRequest) GetPageSize() int32 {
	return 0
}

type ChallengeCollectionRequest struct {
	Req *ecoodav1.GetChallengesRequest
}

func NewChallengesRequest(req *ecoodav1.GetChallengesRequest) *ChallengeCollectionRequest {
	return &ChallengeCollectionRequest{
		Req: req,
	}
}

func (ccb *ChallengeCollectionRequest) GetFields() string {
	return ccb.Req.GetFields()
}

func (ccb *ChallengeCollectionRequest) GetSort() string {
	return ccb.Req.GetSort()
}

func (ccb *ChallengeCollectionRequest) GetPage() int32 {

	return ccb.Req.GetPage()
}

func (ccb *ChallengeCollectionRequest) GetPageSize() int32 {
	return ccb.Req.GetPageSize()
}

type CategoryItemRequest struct {
	Req *ecoodav1.GetCategoryRequest
}

func NewCategoryRequest(req *ecoodav1.GetCategoryRequest) *CategoryItemRequest {
	return &CategoryItemRequest{
		Req: req,
	}
}

func (cib *CategoryItemRequest) GetFields() string {
	return cib.Req.GetFields()
}

func (cib *CategoryItemRequest) GetSort() string {
	return ""
}

func (cib *CategoryItemRequest) GetPage() int32 {
	return 0
}

func (cib *CategoryItemRequest) GetPageSize() int32 {
	return 0
}

type CategoryCollectionRequest struct {
	Req *ecoodav1.GetCategoriesRequest
}

func NewCategoriesRequest(req *ecoodav1.GetCategoriesRequest) *CategoryCollectionRequest {
	return &CategoryCollectionRequest{
		Req: req,
	}
}

func (ccb *CategoryCollectionRequest) GetFields() string {
	return ccb.Req.GetFields()
}

func (ccb *CategoryCollectionRequest) GetSort() string {
	return ccb.Req.GetSort()
}

func (ccb *CategoryCollectionRequest) GetPage() int32 {

	return ccb.Req.GetPage()
}

func (ccb *CategoryCollectionRequest) GetPageSize() int32 {
	return ccb.Req.GetPageSize()
}

type Bag interface {
	*ChallengeItemRequest | *ChallengeCollectionRequest | *CategoryItemRequest | *CategoryCollectionRequest
	GetFields() string
	GetSort() string
	GetPage() int32
	GetPageSize() int32
}

type RequestBag[T Bag] struct {
	req T
}

func NewRequestBag[T Bag](req T) *RequestBag[T] {
	return &RequestBag[T]{
		req,
	}
}

func (r *RequestBag[T]) GetFields() string {
	return r.req.GetFields()
}

func (r *RequestBag[T]) GetSort() string {
	return r.req.GetSort()
}

func (r *RequestBag[T]) GetPage() int32 {
	return r.req.GetPage()
}

func (r *RequestBag[T]) GetPageSize() int32 {
	return r.req.GetPageSize()
}

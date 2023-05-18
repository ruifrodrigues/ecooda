package hateoas

import (
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
)

type ChallengeItemRequest struct {
	Req *challengev1.GetChallengeRequest
}

func NewChallengeRequest(req *challengev1.GetChallengeRequest) *ChallengeItemRequest {
	return &ChallengeItemRequest{
		Req: req,
	}
}

func (cib *ChallengeItemRequest) GetFields() string {
	return cib.Req.GetFields()
}

func (cib *ChallengeItemRequest) GetIncludes() string {
	return cib.Req.GetIncludes()
}

func (cib *ChallengeItemRequest) GetRelationships() string {
	return cib.Req.GetRelationships()
}

func (cib *ChallengeItemRequest) GetPage() int32 {
	return 0
}

func (cib *ChallengeItemRequest) GetPageSize() int32 {
	return 0
}

type ChallengeCollectionRequest struct {
	Req *challengev1.GetChallengesRequest
}

func NewChallengesRequest(req *challengev1.GetChallengesRequest) *ChallengeCollectionRequest {
	return &ChallengeCollectionRequest{
		Req: req,
	}
}

func (ccb *ChallengeCollectionRequest) GetFields() string {
	return ccb.Req.GetFields()
}

func (ccb *ChallengeCollectionRequest) GetIncludes() string {
	return ccb.Req.GetIncludes()
}

func (ccb *ChallengeCollectionRequest) GetRelationships() string {
	return ccb.Req.GetRelationships()
}

func (ccb *ChallengeCollectionRequest) GetPage() int32 {

	return ccb.Req.GetPage()
}

func (ccb *ChallengeCollectionRequest) GetPageSize() int32 {
	return ccb.Req.GetPageSize()
}

type Bag interface {
	*ChallengeItemRequest | *ChallengeCollectionRequest
	GetFields() string
	GetIncludes() string
	GetRelationships() string
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

func (r *RequestBag[T]) GetIncludes() string {
	return r.req.GetIncludes()
}

func (r *RequestBag[T]) GetRelationships() string {
	return r.req.GetRelationships()
}

func (r *RequestBag[T]) GetPage() int32 {
	return r.req.GetPage()
}

func (r *RequestBag[T]) GetPageSize() int32 {
	return r.req.GetPageSize()
}

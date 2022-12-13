package endpoints

import (
	"context"
	"errors"

	"github.com/MasonPhan2110/MicroserviceEx/pkg/watermark"
	"github.com/go-kit/kit/endpoint"

	"github.com/MasonPhan2110/MicroserviceEx/internal"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint     endpoint.Endpoint
}

func NewEndpointSet(svc watermark.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		AddDocumentEndpoint:   MakeAddDocumentEndpoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		WatermarkEndpoint:     MakeWatermarkEndpoint(svc),
	}
}

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}
		return GetResponse{docs, ""}, nil
	}
}

func MakeStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StatusRequest)
		status, err := svc.Status(ctx, req.TicketId)
		if err != nil {
			return StatusResponse{Status: status, Err: err.Error()}, nil
		}
		return StatusResponse{Status: status, Err: ""}, nil
	}
}

func MakeAddDocumentEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDocumentRequest)
		ticketId, err := svc.AddDocument(ctx, req.Document)
		if err != nil {
			return AddDocumentResponse{TicketId: ticketId, Err: err.Error()}, nil
		}
		return AddDocumentResponse{TicketId: ticketId, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WatermarkRequest)
		code, err := svc.Watermark(ctx, req.TicketId, req.Mark)
		if err != nil {
			return WatermarkResponse{Code: code, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}

	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}

	return getResp.Documents, nil
}

func (s *Set) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddDocumentEndpoint(ctx)
}

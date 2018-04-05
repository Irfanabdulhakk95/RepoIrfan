package endpoint

import (
	"context"

	svc "pabriktahu/inventori/server"

	kit "github.com/go-kit/kit/endpoint"
)

type InventoriEndpoint struct {
	AddInventoriEndpoint                 kit.Endpoint
	ReadInventoriByNamaInventoriEndpoint kit.Endpoint
	ReadInventoriEndpoint                kit.Endpoint
	UpdateInventoriEndpoint              kit.Endpoint
}

func NewInventoriEndpoint(service svc.InventoriService) InventoriEndpoint {
	addInventoriEp := makeAddInventoriEndpoint(service)
	readInventoriByNamaInventoriEp := makeReadInventoriByNamaInventoriEndpoint(service)
	readInventoriEp := makeReadInventoriEndpoint(service)
	updateInventoriEp := makeUpdateInventoriEndpoint(service)
	return InventoriEndpoint{AddInventoriEndpoint: addInventoriEp,
		ReadInventoriByNamaInventoriEndpoint: readInventoriByNamaInventoriEp,
		ReadInventoriEndpoint:                readInventoriEp,
		UpdateInventoriEndpoint:              updateInventoriEp,
	}
}

func makeAddInventoriEndpoint(service svc.InventoriService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Inventori)
		err := service.AddInventoriService(ctx, req)
		return nil, err
	}
}

func makeReadInventoriByNamaInventoriEndpoint(service svc.InventoriService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Inventori)
		result, err := service.ReadInventoriByNamaInventoriService(ctx, req.NamaInventori)
		return result, err
	}
}

func makeReadInventoriEndpoint(service svc.InventoriService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadInventoriService(ctx)
		return result, err
	}
}

func makeUpdateInventoriEndpoint(service svc.InventoriService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Inventori)
		err := service.UpdateInventoriService(ctx, req)
		return nil, err
	}
}

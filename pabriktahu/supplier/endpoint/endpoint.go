package endpoint

import (
	"context"

	svc "pabriktahu/supplier/server"

	kit "github.com/go-kit/kit/endpoint"
)

type SupplierEndpoint struct {
	AddSupplierEndpoint              kit.Endpoint
	ReadSupplierByKeteranganEndpoint kit.Endpoint
	ReadSupplierEndpoint             kit.Endpoint
	UpdateSupplierEndpoint           kit.Endpoint
}

func NewSupplierEndpoint(service svc.SupplierService) SupplierEndpoint {
	addSupplierEp := makeAddSupplierEndpoint(service)
	readSupplierByKeteranganEp := makeReadSupplierByKeteranganEndpoint(service)
	readSupplierEp := makeReadSupplierEndpoint(service)
	updateSupplierEp := makeUpdateSupplierEndpoint(service)
	return SupplierEndpoint{AddSupplierEndpoint: addSupplierEp,
		ReadSupplierByKeteranganEndpoint: readSupplierByKeteranganEp,
		ReadSupplierEndpoint:             readSupplierEp,
		UpdateSupplierEndpoint:           updateSupplierEp,
	}
}

func makeAddSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		err := service.AddSupplierService(ctx, req)
		return nil, err
	}
}

func makeReadSupplierByKeteranganEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		result, err := service.ReadSupplierByKeteranganService(ctx, req.Keterangan)
		return result, err
	}
}

func makeReadSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadSupplierService(ctx)
		return result, err
	}
}

func makeUpdateSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		err := service.UpdateSupplierService(ctx, req)
		return nil, err
	}
}

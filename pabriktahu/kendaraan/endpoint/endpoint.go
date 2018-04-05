package endpoint

import (
	"context"

	svc "pabriktahu/kendaraan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KendaraanEndpoint struct {
	AddKendaraanEndpoint             kit.Endpoint
	ReadKendaraanByNomorPlatEndpoint kit.Endpoint
	ReadKendaraanEndpoint            kit.Endpoint
	UpdateKendaraanEndpoint          kit.
	ReadKendaraanByNoPlatEndpoint		 kit.Endpoint
}

func NewKendaraanEndpoint(service svc.KendaraanService) KendaraanEndpoint {
	addKendaraanEp := makeAddKendaraanEndpoint(service)
	readKendaraanByNomorPlatEp := makeReadKendaraanByNomorPlatEndpoint(service)
	readKendaraanEp := makeReadKendaraanEndpoint(service)
	updateKendaraanEp := makeUpdateKendaraanEndpoint(service)
	return KendaraanEndpoint{AddKendaraanEndpoint: addKendaraanEp,
		ReadKendaraanByNomorPlatEndpoint: readKendaraanByNomorPlatEp,
		ReadKendaraanEndpoint:            readKendaraanEp,
		UpdateKendaraanEndpoint:          updateKendaraanEp,
		ReadKendaraanByNoPlatEndpoint	  readKendaraanByNoPlatEp
	}
}

func makeAddKendaraanEndpoint(service svc.KendaraanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kendaraan)
		err := service.AddKendaraanService(ctx, req)
		return nil, err
	}
}

func makeReadKendaraanByNomorPlatEndpoint(service svc.KendaraanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kendaraan)
		result, err := service.ReadKendaraanByNomorPlatService(ctx, req.NomorPlat)
		return result, err
	}
}

func makeReadKendaraanEndpoint(service svc.KendaraanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKendaraanService(ctx)
		return result, err
	}
}
func makeReadKendaraanByNoPlatEndpoint(service svc.KendaraanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKendaraanByNoPlatService(ctx)
		return result, err
	}
}
func makeUpdateKendaraanEndpoint(service svc.KendaraanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kendaraan)
		err := service.UpdateKendaraanService(ctx, req)
		return nil, err
	}
}

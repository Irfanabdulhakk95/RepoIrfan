package endpoint

import (
	"context"

	svc "pabriktahu/karyawan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KaryawanEndpoint struct {
	AddKaryawanEndpoint        kit.Endpoint
	ReadKaryawanByNamaEndpoint kit.Endpoint
	ReadKaryawanEndpoint       kit.Endpoint
	UpdateKaryawanEndpoint     kit.Endpoint
}

func NewKaryawanEndpoint(service svc.KaryawanService) KaryawanEndpoint {
	addKaryawanEp := makeAddKaryawanEndpoint(service)
	readKaryawanByNamaEp := makeReadKaryawanByNamaEndpoint(service)
	readKaryawanEp := makeReadKaryawanEndpoint(service)
	updateKaryawanEp := makeUpdateKaryawanEndpoint(service)
	return KaryawanEndpoint{AddKaryawanEndpoint: addKaryawanEp,
		ReadKaryawanByNamaEndpoint: readKaryawanByNamaEp,
		ReadKaryawanEndpoint:       readKaryawanEp,
		UpdateKaryawanEndpoint:     updateKaryawanEp,
	}
}

func makeAddKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		err := service.AddKaryawanService(ctx, req)
		return nil, err
	}
}

func makeReadKaryawanByNamaEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		result, err := service.ReadKaryawanByNamaService(ctx, req.NamaKaryawan)
		return result, err
	}
}

func makeReadKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKaryawanService(ctx)
		return result, err
	}
}

func makeUpdateKaryawanEndpoint(service svc.KaryawanService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Karyawan)
		err := service.UpdateKaryawanService(ctx, req)
		return nil, err
	}
}

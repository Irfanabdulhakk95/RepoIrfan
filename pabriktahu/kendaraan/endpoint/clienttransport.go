package endpoint

import (
	"context"
	"time"

	svc "pabriktahu/kendaraan/server"

	pb "pabriktahu/kendaraan/grpc"

	util "pabriktahu/util/grpc"
	disc "pabriktahu/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.KendaraanService"
)

func NewGRPCKendaraanClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KendaraanService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKendaraanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKendaraanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKendaraanEp = retry
	}

	var readKendaraanByNomorPlatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKendaraanByNomorPlatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKendaraanByNomorPlatEp = retry
	}

	var readKendaraanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKendaraanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKendaraanEp = retry
	}
	var readKendaraanByNoPlatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKendaraanByNoPlatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKendaraanByNoPlatEp = retry
	}
	var updateKendaraanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKendaraan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKendaraanEp = retry
	}
	return KendaraanEndpoint{AddKendaraanEndpoint: addKendaraanEp, ReadKendaraanByNomorPlatEndpoint: readKendaraanByNomorPlatEp,
		ReadKendaraanEndpoint: readKendaraanEp, ReadKendaraanByNoPlatEndpoint: readKendaraanByNoPlatEp, UpdateKendaraanEndpoint: updateKendaraanEp}, nil
}

func encodeAddKendaraanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kendaraan)
	return &pb.AddKendaraanReq{
		KodeKendaraan:  req.KodeKendaraan,
		JenisKendaraan: req.JenisKendaraan,
		NomorPlat:      req.NomorPlat,
		Status:         req.Status,
		CreateBy:       req.CreateBy,
	}, nil
}

func encodeReadKendaraanByNomorPlatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kendaraan)
	return &pb.ReadKendaraanByNomorPlatReq{NomorPlat: req.NomorPlat}, nil
}

func encodeReadKendaraanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKendaraanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Kendaraan)
	return &pb.UpdateKendaraanReq{
		KodeKendaraan:  req.KodeKendaraan,
		JenisKendaraan: req.JenisKendaraan,
		NomorPlat:      req.NomorPlat,
		Status:         req.Status,
		UpdateBy:       req.UpdateBy,
	}, nil
}

func decodeKendaraanResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKendaraanByNomorPlatRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKendaraanByNomorPlatResp)
	return svc.Kendaraan{
		KodeKendaraan:  resp.KodeKendaraan,
		JenisKendaraan: resp.JenisKendaraan,
		NomorPlat:      resp.NomorPlat,
		Status:         resp.Status,
		CreateBy:       resp.CreateBy,
	}, nil
}

func decodeReadKendaraanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKendaraanResp)
	var rsp svc.Kendaraans

	for _, v := range resp.AllKendaraan {
		itm := svc.Kendaraan{
			KodeKendaraan:  v.KodeKendaraan,
			JenisKendaraan: v.JenisKendaraan,
			NomorPlat:      v.NomorPlat,
			Status:         v.Status,
			CreateBy:       v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKendaraanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKendaraan",
		encodeAddKendaraanRequest,
		decodeKendaraanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKendaraan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKendaraan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKendaraanByNomorPlatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKendaraanByNomorPlat",
		encodeReadKendaraanByNomorPlatRequest,
		decodeReadKendaraanByNomorPlatRespones,
		pb.ReadKendaraanByNomorPlatResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKendaraanByNomorPlat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKendaraanByNomorPlat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKendaraanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKendaraan",
		encodeReadKendaraanRequest,
		decodeReadKendaraanResponse,
		pb.ReadKendaraanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKendaraan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKendaraan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKendaraan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKendaraan",
		encodeUpdateKendaraanRequest,
		decodeKendaraanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKendaraan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKendaraan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

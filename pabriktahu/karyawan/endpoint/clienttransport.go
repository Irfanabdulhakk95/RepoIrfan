package endpoint

import (
	"context"
	"time"

	svc "pabriktahu/karyawan/server"

	pb "pabriktahu/karyawan/grpc"

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
	grpcName = "grpc.KaryawanService"
)

func NewGRPCKaryawanClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KaryawanService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKaryawanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKaryawanEp = retry
	}

	var readKaryawanByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanByNamaEp = retry
	}

	var readKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKaryawanEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKaryawanEp = retry
	}

	var updateKaryawanEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKaryawan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKaryawanEp = retry
	}
	return KaryawanEndpoint{AddKaryawanEndpoint: addKaryawanEp, ReadKaryawanByNamaEndpoint: readKaryawanByNamaEp,
		ReadKaryawanEndpoint: readKaryawanEp, UpdateKaryawanEndpoint: updateKaryawanEp}, nil
}

func encodeAddKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.AddKaryawanReq{
		NamaKaryawan: req.NamaKaryawan,
		IdAgama:      req.IdAgama,
		Alamat:       req.Alamat,
		JenisKelamin: req.Jeniskelamin,
		Status:       req.Status,
		CreateBy:     req.CreateBy,
	}, nil
}

func encodeReadKaryawanByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.ReadKaryawanByNamaReq{NamaKaryawan: req.NamaKaryawan}, nil
}

func encodeReadKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Karyawan)
	return &pb.UpdateKaryawanReq{
		NamaKaryawan: req.NamaKaryawan,
		IdAgama:      req.IdAgama,
		Alamat:       req.Alamat,
		JenisKelamin: req.Jeniskelamin,
		Status:       req.Status,
		UpdateBy:     req.UpdateBy,
	}, nil
}

func decodeKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKaryawanByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanByNamaResp)
	return svc.Karyawan{
		NamaKaryawan: resp.NamaKaryawan,
		IdAgama:      resp.IdAgama,
		Alamat:       resp.Alamat,
		Jeniskelamin: resp.JenisKelamin,
		Status:       resp.Status,
		CreateBy:     resp.CreateBy,
	}, nil
}

func decodeReadKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKaryawanResp)
	var rsp svc.Karyawans

	for _, v := range resp.AllKaryawan {
		itm := svc.Karyawan{
			NamaKaryawan: v.NamaKaryawan,
			IdAgama:      v.IdAgama,
			Alamat:       v.Alamat,
			Jeniskelamin: v.JenisKelamin,
			Status:       v.Status,
			CreateBy:     v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKaryawanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKaryawan",
		encodeAddKaryawanRequest,
		decodeKaryawanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawanByNama",
		encodeReadKaryawanByNamaRequest,
		decodeReadKaryawanByNamaRespones,
		pb.ReadKaryawanByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawanByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawanByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKaryawanEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKaryawan",
		encodeReadKaryawanRequest,
		decodeReadKaryawanResponse,
		pb.ReadKaryawanResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKaryawan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKaryawan",
		encodeUpdateKaryawanRequest,
		decodeKaryawanResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKaryawan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKaryawan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

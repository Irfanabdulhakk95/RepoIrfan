package endpoint

import (
	"context"
	"time"

	svc "pabriktahu/supplier/server"

	pb "pabriktahu/supplier/grpc"

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
	grpcName = "grpc.SupplierService"
)

func NewGRPCSupplierClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.SupplierService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addSupplierEp = retry
	}

	var readSupplierByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadSupplierByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readSupplierByKeteranganEp = retry
	}

	var readSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readSupplierEp = retry
	}

	var updateSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateSupplier, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateSupplierEp = retry
	}
	return SupplierEndpoint{
		AddSupplierEndpoint:              addSupplierEp,
		ReadSupplierByKeteranganEndpoint: readSupplierByKeteranganEp,
		ReadSupplierEndpoint:             readSupplierEp,
		UpdateSupplierEndpoint:           updateSupplierEp}, nil
}

func encodeAddSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.AddSupplierReq{
		KodeSupplier: req.KodeSupplier,
		NamaSupplier: req.NamaSupplier,
		Keterangan:   req.Keterangan,
		Status:       req.Status,
		CreateBy:     req.CreateBy,
	}, nil
}

func encodeReadSupplierByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.ReadSupplierByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.UpdateSupplierReq{
		KodeSupplier: req.KodeSupplier,
		NamaSupplier: req.NamaSupplier,
		Keterangan:   req.Keterangan,
		Status:       req.Status,
		UpdateBy:     req.UpdateBy,
	}, nil
}

func decodeSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadSupplierByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadSupplierByKeteranganResp)
	var rsp svc.Suppliers

	for _, v := range resp.AllKeterangan {
		itm := svc.Supplier{
			KodeSupplier: v.KodeSupplier,
			NamaSupplier: v.NamaSupplier,
			Keterangan:   v.Keterangan,
			Status:       v.Status,
			CreateBy:     v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadSupplierResp)
	var rsp svc.Suppliers

	for _, v := range resp.AllSupplier {
		itm := svc.Supplier{
			KodeSupplier: v.KodeSupplier,
			NamaSupplier: v.NamaSupplier,
			Keterangan:   v.Keterangan,
			Status:       v.Status,
			CreateBy:     v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddSupplier",
		encodeAddSupplierRequest,
		decodeSupplierResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadSupplierByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadSupplierByKeterangan",
		encodeReadSupplierByKeteranganRequest,
		decodeReadSupplierByKeteranganResponse,
		pb.ReadSupplierByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadSupplierByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadSupplierByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadSupplier",
		encodeReadSupplierRequest,
		decodeReadSupplierResponse,
		pb.ReadSupplierResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateSupplier(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateSupplier",
		encodeUpdateSupplierRequest,
		decodeSupplierResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

package endpoint

import (
	"context"
	"time"

	svc "pabriktahu/inventori/server"

	pb "pabriktahu/inventori/grpc"

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
	grpcName = "grpc.InventoriService"
)

func NewGRPCInventoriClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.InventoriService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addInventoriEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddInventoriEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addInventoriEp = retry
	}

	var readInventoriByNamaInventoriEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadInventoriByNamaInventoriEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readInventoriByNamaInventoriEp = retry
	}

	var readInventoriEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadInventoriEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readInventoriEp = retry
	}

	var updateInventoriEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateInventori, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateInventoriEp = retry
	}
	return InventoriEndpoint{AddInventoriEndpoint: addInventoriEp, ReadInventoriByNamaInventoriEndpoint: readInventoriByNamaInventoriEp,
		ReadInventoriEndpoint: readInventoriEp, UpdateInventoriEndpoint: updateInventoriEp}, nil
}

func encodeAddInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Inventori)
	return &pb.AddInventoriReq{
		KodeInventori: req.KodeInventori,
		NamaInventori: req.NamaInventori,
		Jumlah:        req.Jumlah,
		Status:        req.Status,
		CreateBy:      req.CreateBy,
	}, nil
}

func encodeReadInventoriByNamaInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Inventori)
	return &pb.ReadInventoriByNamaInventoriReq{NamaInventori: req.NamaInventori}, nil
}

func encodeReadInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Inventori)
	return &pb.UpdateInventoriReq{
		KodeInventori: req.KodeInventori,
		NamaInventori: req.NamaInventori,
		Jumlah:        req.Jumlah,
		Status:        req.Status,
		UpdateBy:      req.UpdateBy,
	}, nil
}

func decodeInventoriResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadInventoriByNamaInventoriRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadInventoriByNamaInventoriResp)
	return svc.Inventori{
		KodeInventori: resp.KodeInventori,
		NamaInventori: resp.NamaInventori,
		Jumlah:        resp.Jumlah,
		Status:        resp.Status,
		CreateBy:      resp.CreateBy,
	}, nil
}

func decodeReadInventoriResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadInventoriResp)
	var rsp svc.Inventoris

	for _, v := range resp.AllInventori {
		itm := svc.Inventori{
			KodeInventori: v.KodeInventori,
			NamaInventori: v.NamaInventori,
			Jumlah:        v.Jumlah,
			Status:        v.Status,
			CreateBy:      v.CreateBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddInventoriEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddInventori",
		encodeAddInventoriRequest,
		decodeInventoriResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddInventori")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddInventori",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadInventoriByNamaInventoriEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadInventoriByNamaInventori",
		encodeReadInventoriByNamaInventoriRequest,
		decodeReadInventoriByNamaInventoriRespones,
		pb.ReadInventoriByNamaInventoriResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadInventoriByNamaInventori")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadInventoriByNamaInventori",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadInventoriEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadInventori",
		encodeReadInventoriRequest,
		decodeReadInventoriResponse,
		pb.ReadInventoriResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadInventori")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadInventori",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateInventori(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateInventori",
		encodeUpdateInventoriRequest,
		decodeInventoriResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateInventori")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateInventori",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

package endpoint

import (
	"context"
	"time"

	svc "git.bluebird.id/pabrikTahu/user/server"

	pb "git.bluebird.id/pabrikTahu/user/grpc"

	util "git.bluebird.id/pabrikTahu/util/grpc"
	disc "git.bluebird.id/pabrikTahu/util/microservice"

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
	grpcName = "grpc.UserService"
)

func NewGRPCUserClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.UserService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddUserEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addUserEp = retry
	}

	var readUserByKetEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadUserByKetEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readUserByKetEp = retry
	}

	var readUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadUserEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readUserEp = retry
	}

	var updateUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateUser, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateUserEp = retry
	}
	return UserEndpoint{
		AddUserEndpoint:       addUserEp,
		ReadUserByKetEndpoint: readUserByKetEp,
		ReadUserEndpoint:      readUserEp,
		UpdateUserEndpoint:    updateUserEp}, nil
}

func encodeAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.AddUserReq{
		Username:   req.Username,
		Password:   req.Password,
		IDKaryawan: req.IDKaryawan,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		Keterangan: req.Keterangan,
	}, nil
}

func encodeReadUserByKetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.ReadUserByKetReq{Keterangan: req.Keterangan}, nil
}

func encodeReadUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.UpdateUserReq{
		Username:   req.Username,
		Password:   req.Password,
		IDKaryawan: req.IDKaryawan,
		Status:     req.Status,
		UpdatedBy:  req.UpdatedBy,
		Keterangan: req.Keterangan,
	}, nil
}

func decodeUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadUserByKetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadUserByKetResp)
	var rsp svc.Users

	for _, v := range resp.AllKet {
		itm := svc.User{
			Username:   v.Username,
			Password:   v.Password,
			IDKaryawan: v.IDKaryawan,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			Keterangan: v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadUserResp)
	var rsp svc.Users

	for _, v := range resp.AllUser {
		itm := svc.User{
			Username:   v.Username,
			Password:   v.Password,
			IDKaryawan: v.IDKaryawan,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			Keterangan: v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddUserEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddUser",
		encodeAddUserRequest,
		decodeUserResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadUserByKetEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadUserByKet",
		encodeReadUserByKetRequest,
		decodeReadUserByKetResponse,
		pb.ReadUserByKetResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadUserByKet")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadUserByKet",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadUserEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadUser",
		encodeReadUserRequest,
		decodeReadUserResponse,
		pb.ReadUserResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateUser(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateUser",
		encodeUpdateUserRequest,
		decodeUserResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

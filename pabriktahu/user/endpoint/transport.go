package endpoint

import (
	"context"

	scv "git.bluebird.id/pabrikTahu/user/server"

	pb "git.bluebird.id/pabrikTahu/user/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcUserServer struct {
	addUser       grpctransport.Handler
	readUserByKet grpctransport.Handler
	readUser      grpctransport.Handler
	updateUser    grpctransport.Handler
}

func NewGRPCUserServer(endpoints UserEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.UserServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcUserServer{
		addUser: grpctransport.NewServer(endpoints.AddUserEndpoint,
			decodeAddUserRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddUser", logger)))...),

		readUserByKet: grpctransport.NewServer(endpoints.ReadUserByKetEndpoint,
			decodeReadUserByKetRequest,
			ReadUserByKetResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadUserByKet", logger)))...),

		readUser: grpctransport.NewServer(endpoints.ReadUserEndpoint,
			decodeReadUserRequest,
			encodeReadUserResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadUser", logger)))...),
		updateUser: grpctransport.NewServer(endpoints.UpdateUserEndpoint,
			decodeUpdateUserRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUser", logger)))...),
	}
}

func decodeAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {

	req := request.(*pb.AddUserReq)
	return scv.User{Username: req.GetUsername(), Password: req.GetPassword(), IDKaryawan: req.GetIDKaryawan(),
		Status: req.GetStatus(), CreatedBy: req.GetCreatedBy(), Keterangan: req.GetKeterangan()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcUserServer) AddUser(ctx oldcontext.Context, user *pb.AddUserReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addUser.ServeGRPC(ctx, user)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadUserByKetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadUserByKetReq)
	return scv.User{Keterangan: req.Keterangan}, nil
}

func decodeReadUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func ReadUserByKetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Users)

	rsp := &pb.ReadUserByKetResp{}

	for _, v := range resp {
		itm := &pb.ReadUserByKet{
			Username:   v.Username,
			Password:   v.Password,
			IDKaryawan: v.IDKaryawan,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			Keterangan: v.Keterangan,
		}
		rsp.AllKet = append(rsp.AllKet, itm)
	}
	return rsp, nil
}

func encodeReadUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Users)

	rsp := &pb.ReadUserResp{}

	for _, v := range resp {
		itm := &pb.ReadUserByKet{
			Username:   v.Username,
			Password:   v.Password,
			IDKaryawan: v.IDKaryawan,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
			Keterangan: v.Keterangan,
		}
		rsp.AllUser = append(rsp.AllUser, itm)
	}
	return rsp, nil
}

func (s *grpcUserServer) ReadUserByKet(ctx oldcontext.Context, keterangan *pb.ReadUserByKetReq) (*pb.ReadUserByKetResp, error) {
	_, resp, err := s.readUserByKet.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadUserByKetResp), nil
}

func (s *grpcUserServer) ReadUser(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadUserResp, error) {
	_, resp, err := s.readUser.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadUserResp), nil
}

func decodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserReq)
	return scv.User{Username: req.Username, Password: req.Password, IDKaryawan: req.IDKaryawan,
		Status: req.Status, UpdatedBy: req.UpdatedBy}, nil
}

func (s *grpcUserServer) UpdateUser(ctx oldcontext.Context, user *pb.UpdateUserReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, user)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

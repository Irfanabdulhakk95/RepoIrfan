package endpoint

import (
	"context"

	scv "pabriktahu/kendaraan/server"

	pb "pabriktahu/kendaraan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKendaraanServer struct {
	addKendaraan             grpctransport.Handler
	readKendaraanByNomorPlat grpctransport.Handler
	readKendaraan            grpctransport.Handler
	updateKendaraan          grpctransport.Handler
	readKendaraanByNoPlat    grpctransport.Handler
}

func NewGRPCKendaraanServer(endpoints KendaraanEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KendaraanServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKendaraanServer{
		addKendaraan: grpctransport.NewServer(endpoints.AddKendaraanEndpoint,
			decodeAddKendaraanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKendaraan", logger)))...),
		readKendaraanByNomorPlat: grpctransport.NewServer(endpoints.ReadKendaraanByNomorPlatEndpoint,
			decodeReadKendaraanByNomorPlatRequest,
			encodeReadKendaraanByNomorPlatResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKendaraanByNomorPlat", logger)))...),
		readKendaraanByNoPlat: grpctransport.NewServer(endpoints.ReadKendaraanByNoPlatEndpoint,
			decodeReadKendaraanByNoPlatRequest,
			encodeReadKendaraanByNoPlatResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKendaraanByNoPlat", logger)))...),
		readKendaraan: grpctransport.NewServer(endpoints.ReadKendaraanEndpoint,
			decodeReadKendaraanRequest,
			encodeReadKendaraanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKendaraan", logger)))...),
		updateKendaraan: grpctransport.NewServer(endpoints.UpdateKendaraanEndpoint,
			decodeUpdateKendaraanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKendaraan", logger)))...),
	}
}

func decodeAddKendaraanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKendaraanReq)
	return scv.Kendaraan{KodeKendaraan: req.GetKodeKendaraan(), JenisKendaraan: req.GetJenisKendaraan(), NomorPlat: req.GetNomorPlat(),
		Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcKendaraanServer) AddKendaraan(ctx oldcontext.Context, kendaraan *pb.AddKendaraanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKendaraan.ServeGRPC(ctx, kendaraan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadKendaraanByNomorPlatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKendaraanByNomorPlatReq)
	return scv.Kendaraan{NomorPlat: req.NomorPlat}, nil
}

func encodeReadKendaraanByNomorPlatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kendaraan)
	return &pb.ReadKendaraanByNomorPlatResp{KodeKendaraan: resp.KodeKendaraan, JenisKendaraan: resp.JenisKendaraan, NomorPlat: resp.NomorPlat,
		Status: resp.Status, CreateBy: resp.CreateBy}, nil
}

func decodeReadKendaraanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadKendaraanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kendaraans)

	rsp := &pb.ReadKendaraanResp{}

	for _, v := range resp {
		itm := &pb.ReadKendaraanByNomorPlatResp{
			KodeKendaraan:  v.KodeKendaraan,
			JenisKendaraan: v.JenisKendaraan,
			NomorPlat:      v.NomorPlat,
			Status:         v.Status,
			CreateBy:       v.CreateBy,
		}
		rsp.AllKendaraan = append(rsp.AllKendaraan, itm)
	}
	return rsp, nil
}

func decodeReadKendaraanByNoPlatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}
func encodeReadKendaraanByNoPlatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kendaraans)

	rsp := &pb.ReadKendaraanResp{}

	for _, v := range resp {
		itm := &pb.ReadKendaraanByNoPlatResp{
			KodeKendaraan:  v.KodeKendaraan,
			JenisKendaraan: v.JenisKendaraan,
			NomorPlat:      v.NomorPlat,
			Status:         v.Status,
			CreateBy:       v.CreateBy,
		}
		rsp.AllKendaraan = append(rsp.AllNoPlat, itm)
	}
	return rsp, nil
}
func (s *grpcKendaraanServer) ReadKendaraanByNomorPlat(ctx oldcontext.Context, nomorplat *pb.ReadKendaraanByNomorPlatReq) (*pb.ReadKendaraanByNomorPlatResp, error) {
	_, resp, err := s.readKendaraanByNomorPlat.ServeGRPC(ctx, nomorplat)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKendaraanByNomorPlatResp), nil
}

func (s *grpcKendaraanServer) ReadKendaraan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKendaraanResp, error) {
	_, resp, err := s.readKendaraan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKendaraanResp), nil
}

func (s *grpcKendaraanServer) ReadKendaraanByNoPlat(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKendaraanByNoPlatResp, error) {
	_, resp, err := s.readKendaraanByNoPlat.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKendaraanRespByNoPlat), nil
}

func decodeUpdateKendaraanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKendaraanReq)
	return scv.Kendaraan{KodeKendaraan: req.KodeKendaraan, JenisKendaraan: req.JenisKendaraan, NomorPlat: req.NomorPlat,
		Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func (s *grpcKendaraanServer) UpdateKendaraan(ctx oldcontext.Context, cus *pb.UpdateKendaraanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKendaraan.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

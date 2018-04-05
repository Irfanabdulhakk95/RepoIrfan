package endpoint

import (
	"context"

	scv "pabriktahu/karyawan/server"

	pb "pabriktahu/karyawan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKaryawanServer struct {
	addKaryawan        grpctransport.Handler
	readKaryawanByNama grpctransport.Handler
	readKaryawan       grpctransport.Handler
	updateKaryawan     grpctransport.Handler
}

func NewGRPCKaryawanServer(endpoints KaryawanEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KaryawanServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKaryawanServer{
		addKaryawan: grpctransport.NewServer(endpoints.AddKaryawanEndpoint,
			decodeAddKaryawanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKaryawan", logger)))...),
		readKaryawanByNama: grpctransport.NewServer(endpoints.ReadKaryawanByNamaEndpoint,
			decodeReadKaryawanByNamaRequest,
			encodeReadKaryawanByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawanByNama", logger)))...),
		readKaryawan: grpctransport.NewServer(endpoints.ReadKaryawanEndpoint,
			decodeReadKaryawanRequest,
			encodeReadKaryawanResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKaryawan", logger)))...),
		updateKaryawan: grpctransport.NewServer(endpoints.UpdateKaryawanEndpoint,
			decodeUpdateKaryawanRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKaryawan", logger)))...),
	}
}

func decodeAddKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKaryawanReq)
	return scv.Karyawan{NamaKaryawan: req.GetNamaKaryawan(), IdAgama: req.GetIdAgama(), Alamat: req.GetAlamat(), Jeniskelamin: req.GetJenisKelamin(),
		Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcKaryawanServer) AddKaryawan(ctx oldcontext.Context, karyawan *pb.AddKaryawanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKaryawan.ServeGRPC(ctx, karyawan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadKaryawanByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKaryawanByNamaReq)
	return scv.Karyawan{NamaKaryawan: req.NamaKaryawan}, nil
}

func decodeReadKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadKaryawanByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawan)
	return &pb.ReadKaryawanByNamaResp{NamaKaryawan: resp.NamaKaryawan, IdAgama: resp.IdAgama, Alamat: resp.Alamat, JenisKelamin: resp.Jeniskelamin,
		Status: resp.Status, CreateBy: resp.CreateBy}, nil
}

func encodeReadKaryawanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Karyawans)

	rsp := &pb.ReadKaryawanResp{}

	for _, v := range resp {
		itm := &pb.ReadKaryawanByNamaResp{
			NamaKaryawan: v.NamaKaryawan,
			IdAgama:      v.IdAgama,
			Alamat:       v.Alamat,
			JenisKelamin: v.Jeniskelamin,
			Status:       v.Status,
			CreateBy:     v.CreateBy,
		}
		rsp.AllKaryawan = append(rsp.AllKaryawan, itm)
	}
	return rsp, nil
}

func (s *grpcKaryawanServer) ReadKaryawanByNama(ctx oldcontext.Context, namakaryawan *pb.ReadKaryawanByNamaReq) (*pb.ReadKaryawanByNamaResp, error) {
	_, resp, err := s.readKaryawanByNama.ServeGRPC(ctx, namakaryawan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanByNamaResp), nil
}

func (s *grpcKaryawanServer) ReadKaryawan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKaryawanResp, error) {
	_, resp, err := s.readKaryawan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKaryawanResp), nil
}

func decodeUpdateKaryawanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKaryawanReq)
	return scv.Karyawan{NamaKaryawan: req.NamaKaryawan, IdAgama: req.IdAgama, Alamat: req.Alamat,
		Jeniskelamin: req.JenisKelamin, Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func (s *grpcKaryawanServer) UpdateKaryawan(ctx oldcontext.Context, cus *pb.UpdateKaryawanReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKaryawan.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

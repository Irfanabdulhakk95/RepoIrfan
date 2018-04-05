package endpoint

import (
	"context"

	scv "pabriktahu/supplier/server"

	pb "pabriktahu/supplier/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcSupplierServer struct {
	addSupplier              grpctransport.Handler
	readSupplierByKeterangan grpctransport.Handler
	readSupplier             grpctransport.Handler
	updateSupplier           grpctransport.Handler
}

func NewGRPCSupplierServer(endpoints SupplierEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.SupplierServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcSupplierServer{
		addSupplier: grpctransport.NewServer(endpoints.AddSupplierEndpoint,
			decodeAddSupplierRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddSupplier", logger)))...),
		readSupplierByKeterangan: grpctransport.NewServer(endpoints.ReadSupplierByKeteranganEndpoint,
			decodeReadSupplierByKeteranganRequest,
			encodeReadSupplierByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadSupplierByKeterangan", logger)))...),
		readSupplier: grpctransport.NewServer(endpoints.ReadSupplierEndpoint,
			decodeReadSupplierRequest,
			encodeReadSupplierResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadSupplier", logger)))...),
		updateSupplier: grpctransport.NewServer(endpoints.UpdateSupplierEndpoint,
			decodeUpdateSupplierRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateSupplier", logger)))...),
	}
}

func decodeAddSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddSupplierReq)
	return scv.Supplier{KodeSupplier: req.GetKodeSupplier(), NamaSupplier: req.GetNamaSupplier(), Keterangan: req.GetKeterangan(),
		Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcSupplierServer) AddSupplier(ctx oldcontext.Context, supplier *pb.AddSupplierReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addSupplier.ServeGRPC(ctx, supplier)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadSupplierByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadSupplierByKeteranganReq)
	return scv.Supplier{Keterangan: req.Keterangan}, nil
}

func decodeReadSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadSupplierByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Suppliers)
	rsp := &pb.ReadSupplierByKeteranganResp{}
	for _, v := range resp {
		itm := &pb.AddSupplierReq{
			KodeSupplier: v.KodeSupplier,
			NamaSupplier: v.NamaSupplier,
			Keterangan:   v.Keterangan,
			Status:       v.Status,
			CreateBy:     v.CreateBy,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func encodeReadSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Suppliers)

	rsp := &pb.ReadSupplierResp{}

	for _, v := range resp {
		itm := &pb.AddSupplierReq{
			KodeSupplier: v.KodeSupplier,
			NamaSupplier: v.NamaSupplier,
			Keterangan:   v.Keterangan,
			Status:       v.Status,
			CreateBy:     v.CreateBy,
		}
		rsp.AllSupplier = append(rsp.AllSupplier, itm)
	}
	return rsp, nil
}

func (s *grpcSupplierServer) ReadSupplierByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadSupplierByKeteranganReq) (*pb.ReadSupplierByKeteranganResp, error) {
	_, resp, err := s.readSupplierByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadSupplierByKeteranganResp), nil
}

func (s *grpcSupplierServer) ReadSupplier(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadSupplierResp, error) {
	_, resp, err := s.readSupplier.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadSupplierResp), nil
}

func decodeUpdateSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateSupplierReq)
	return scv.Supplier{KodeSupplier: req.KodeSupplier, NamaSupplier: req.NamaSupplier, Keterangan: req.Keterangan,
		Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func (s *grpcSupplierServer) UpdateSupplier(ctx oldcontext.Context, cus *pb.UpdateSupplierReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateSupplier.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

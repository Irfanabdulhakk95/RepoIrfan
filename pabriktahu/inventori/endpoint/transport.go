package endpoint

import (
	"context"

	scv "pabriktahu/inventori/server"

	pb "pabriktahu/inventori/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcInventoriServer struct {
	addInventori                 grpctransport.Handler
	readInventoriByNamaInventori grpctransport.Handler
	readInventori                grpctransport.Handler
	updateInventori              grpctransport.Handler
}

func NewGRPCInventoriServer(endpoints InventoriEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.InventoriServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcInventoriServer{
		addInventori: grpctransport.NewServer(endpoints.AddInventoriEndpoint,
			decodeAddInventoriRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddInventori", logger)))...),
		readInventoriByNamaInventori: grpctransport.NewServer(endpoints.ReadInventoriByNamaInventoriEndpoint,
			decodeReadInventoriByNamaInventoriRequest,
			encodeReadInventoriByNamaInventoriResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadInventoriByNamaInventori", logger)))...),
		readInventori: grpctransport.NewServer(endpoints.ReadInventoriEndpoint,
			decodeReadInventoriRequest,
			encodeReadInventoriResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadInventori", logger)))...),
		updateInventori: grpctransport.NewServer(endpoints.UpdateInventoriEndpoint,
			decodeUpdateInventoriRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateInventori", logger)))...),
	}
}

func decodeAddInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddInventoriReq)
	return scv.Inventori{KodeInventori: req.GetKodeInventori(), NamaInventori: req.GetNamaInventori(), Jumlah: req.GetJumlah(),
		Status: req.GetStatus(), CreateBy: req.GetCreateBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcInventoriServer) AddInventori(ctx oldcontext.Context, inventori *pb.AddInventoriReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addInventori.ServeGRPC(ctx, inventori)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadInventoriByNamaInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadInventoriByNamaInventoriReq)
	return scv.Inventori{NamaInventori: req.NamaInventori}, nil
}

func decodeReadInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadInventoriByNamaInventoriResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Inventori)
	return &pb.ReadInventoriByNamaInventoriResp{KodeInventori: resp.KodeInventori, NamaInventori: resp.NamaInventori, Jumlah: resp.Jumlah,
		Status: resp.Status, CreateBy: resp.CreateBy}, nil
}

func encodeReadInventoriResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Inventoris)

	rsp := &pb.ReadInventoriResp{}

	for _, v := range resp {
		itm := &pb.ReadInventoriByNamaInventoriResp{
			KodeInventori: v.KodeInventori,
			NamaInventori: v.NamaInventori,
			Jumlah:        v.Jumlah,
			Status:        v.Status,
			CreateBy:      v.CreateBy,
		}
		rsp.AllInventori = append(rsp.AllInventori, itm)
	}
	return rsp, nil
}

func (s *grpcInventoriServer) ReadInventoriByNamaInventori(ctx oldcontext.Context, namainventori *pb.ReadInventoriByNamaInventoriReq) (*pb.ReadInventoriByNamaInventoriResp, error) {
	_, resp, err := s.readInventoriByNamaInventori.ServeGRPC(ctx, namainventori)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadInventoriByNamaInventoriResp), nil
}

func (s *grpcInventoriServer) ReadInventori(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadInventoriResp, error) {
	_, resp, err := s.readInventori.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadInventoriResp), nil
}

func decodeUpdateInventoriRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateInventoriReq)
	return scv.Inventori{KodeInventori: req.KodeInventori, NamaInventori: req.NamaInventori, Jumlah: req.Jumlah,
		Status: req.Status, UpdateBy: req.UpdateBy}, nil
}

func (s *grpcInventoriServer) UpdateInventori(ctx oldcontext.Context, cus *pb.UpdateInventoriReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateInventori.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}

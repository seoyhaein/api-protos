package service

import (
	"context"
	pb "github.com/seoyhaein/api-protos/gen/go/datablock/ichthys/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DataBlockHandler interface {
	GetDataBlock(ctx context.Context, updateAt *timestamppb.Timestamp) (*pb.DataBlock, error)
}

type dataBlockServiceServerImpl struct {
	pb.UnimplementedDataBlockServiceServer
	handler DataBlockHandler
}

func NewDataBlockServiceServer(h DataBlockHandler) pb.DataBlockServiceServer {
	return &dataBlockServiceServerImpl{handler: h}
}

func (s *dataBlockServiceServerImpl) GetDataBlock(ctx context.Context, in *pb.FetchDataBlockRequest) (*pb.FetchDataBlockResponse, error) {
	// TODO 위힘 함. 추가적인 설명 달아 놓을 것
	data, err := s.handler.GetDataBlock(ctx, in.IfModifiedSince)

	if err != nil {
		return nil, err
	}

	// 2) 변경 없음(no_update) 케이스
	if data == nil {
		return &pb.FetchDataBlockResponse{
			Payload: &pb.FetchDataBlockResponse_NoUpdate{
				NoUpdate: &emptypb.Empty{},
			},
		}, nil
	}

	// 3) 변경된 DataBlock(data_block) 케이스
	return &pb.FetchDataBlockResponse{
		Payload: &pb.FetchDataBlockResponse_DataBlock{
			DataBlock: data,
		},
	}, nil
}

func RegisterDataBlockServiceServer(s *grpc.Server, h DataBlockHandler) {
	pb.RegisterDataBlockServiceServer(s, NewDataBlockServiceServer(h))
}

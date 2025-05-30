package service

import (
	"context"
	pb "github.com/seoyhaein/api-protos/gen/go/datablock/ichthys/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 좀더 살펴보자.
// 문서화 해놓자. 자꾸 잊어버림.

// 1) 최소한의 핸들러 인터페이스 정의
type DataBlockHandler interface {
	// 비즈니스 로직에 맞춘 시그니처
	GetDataBlock(ctx context.Context, updateAt *timestamppb.Timestamp) (*pb.DataBlock, error)
}

// 2) gRPC 서버 구조체는 핸들러만 보유
type dataBlockServiceServerImpl struct {
	pb.UnimplementedDataBlockServiceServer
	handler DataBlockHandler
}

// 3) 생성자: 핸들러 인터페이스를 받음
func NewDataBlockServiceServer(h DataBlockHandler) pb.DataBlockServiceServer {
	return &dataBlockServiceServerImpl{handler: h}
}

// 4) gRPC 메서드는 핸들러에 위임
func (s *dataBlockServiceServerImpl) GetDataBlock(ctx context.Context, in *pb.FetchDataBlockRequest) (*pb.FetchDataBlockResponse, error) {
	data, err := s.handler.GetDataBlock(ctx, in.IfModifiedSince)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return &pb.GetDataBlockResponse{NoUpdate: true}, nil
	}
	return &pb.GetDataBlockResponse{Data: data, NoUpdate: false}, nil
}

// 5) 서버 등록 헬퍼
func RegisterDataBlockService(s *grpc.Server, h DataBlockHandler) {
	pb.RegisterDataBlockServiceServer(s, NewDataBlockServiceServer(h))
}

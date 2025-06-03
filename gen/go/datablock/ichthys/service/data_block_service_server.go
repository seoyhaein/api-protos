package service

import (
	"context"
	"fmt"
	pb "github.com/seoyhaein/api-protos/gen/go/datablock/ichthys"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"sort"
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

// SaveProtoToFile proto.Message → 바이너리 파일로 저장
func SaveProtoToFile(filePath string, message proto.Message, perm os.FileMode) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}
	if err := os.WriteFile(filePath, data, perm); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

// LoadFileBlock 바이너리 파일 → *pb.FileBlock 역직렬화
func LoadFileBlock(filePath string) (*pb.FileBlock, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	fb := &pb.FileBlock{}
	if err := proto.Unmarshal(data, fb); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pb.FileBlock: %w", err)
	}
	return fb, nil
}

// LoadDataBlock 바이너리 파일 → *pb.DataBlock 역직렬화
func LoadDataBlock(filePath string) (*pb.DataBlock, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	db := &pb.DataBlock{}
	if err := proto.Unmarshal(data, db); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pb.DataBlock: %w", err)
	}
	return db, nil
}

// SaveFileBlockToTextFile FileBlock → 텍스트(proto 텍스트 포맷)로 저장
func SaveFileBlockToTextFile(filePath string, data *pb.FileBlock) error {
	txtBytes, err := prototext.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal to text format: %w", err)
	}
	return os.WriteFile(filePath, txtBytes, os.ModePerm)
}

// MergeFileBlocks 여러 바이너리 FileBlock 파일들을 읽어 하나의 DataBlock으로 병합→저장
func MergeFileBlocks(inputFiles []string, outputFile string) error {
	var blocks []*pb.FileBlock

	for _, fn := range inputFiles {
		fb, err := LoadFileBlock(fn)
		if err != nil {
			return fmt.Errorf("failed to load FileBlock %s: %w", fn, err)
		}
		blocks = append(blocks, fb)
	}

	dataBlock := &pb.DataBlock{
		UpdatedAt: timestamppb.Now(),
		Blocks:    blocks,
	}
	if err := SaveProtoToFile(outputFile, dataBlock, os.ModePerm); err != nil {
		return fmt.Errorf("failed to save DataBlock: %w", err)
	}
	fmt.Printf("Merged %d FileBlock → %s\n", len(inputFiles), outputFile)
	return nil
}

// MergeFileBlocksFromData 이미 로드된 []*FileBlock → *DataBlock 반환
func MergeFileBlocksFromData(inputBlocks []*pb.FileBlock) (*pb.DataBlock, error) {
	if len(inputBlocks) == 0 {
		return nil, fmt.Errorf("no input blocks provided")
	}
	db := &pb.DataBlock{
		UpdatedAt: timestamppb.Now(),
		Blocks:    inputBlocks,
	}
	fmt.Printf("Merged %d FileBlock 객체 into one DataBlock\n", len(inputBlocks))
	return db, nil
}

// SaveDataBlockToTextFile DataBlock → 텍스트 포맷(proto text)으로 저장
func SaveDataBlockToTextFile(filePath string, data *pb.DataBlock) error {
	txtBytes, err := prototext.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal DataBlock to text format: %w", err)
	}
	if err := os.WriteFile(filePath, txtBytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write %s: %w", filePath, err)
	}
	fmt.Printf("Saved DataBlock → %s\n", filePath)
	return nil
}

// ConvertMapToFileBlock rules 패키지에서 생성된 (map[int]map[string]string) → *pb.FileBlock 반환
func ConvertMapToFileBlock(rows map[int]map[string]string, headers []string, blockID string) *pb.FileBlock {
	fb := &pb.FileBlock{
		BlockId:       blockID,
		ColumnHeaders: headers,
		Rows:          make([]*pb.Row, 0, len(rows)),
	}

	// 인덱스를 정렬해서 순차 처리 (RowNumber: 0-based 혹은 1-based 조정 가능)
	var idxs []int
	for i := range rows {
		idxs = append(idxs, i)
	}
	sort.Ints(idxs)

	for _, i := range idxs {
		cols := rows[i]
		r := &pb.Row{
			RowNumber: int32(i),
			Cells:     make(map[string]string, len(cols)),
		}
		for colKey, val := range cols {
			r.Cells[colKey] = val
		}
		fb.Rows = append(fb.Rows, r)
	}
	return fb
}

// GenerateRows 테스트용 로우 데이터 → []*pb.Row
func GenerateRows(data [][]string, headers []string) []*pb.Row {
	out := make([]*pb.Row, 0, len(data))
	for i, cells := range data {
		r := &pb.Row{
			RowNumber: int32(i + 1),
			Cells:     make(map[string]string, len(headers)),
		}
		for j, h := range headers {
			if j < len(cells) {
				r.Cells[h] = cells[j]
			}
		}
		out = append(out, r)
	}
	return out
}

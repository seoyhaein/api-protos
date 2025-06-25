package service

import (
	pb "github.com/seoyhaein/api-protos/gen/go/syncfolders/ichthys"
	"google.golang.org/grpc"
)

type SyncFoldersHandler interface {
}

type syncFoldersServiceServerImpl struct {
	pb.UnimplementedSyncFoldersServiceServer
	handler SyncFoldersHandler
}

func NewSyncFoldersServiceServer(h SyncFoldersHandler) pb.SyncFoldersServiceServer {
	return &syncFoldersServiceServerImpl{handler: h}
}

func RegisterSyncFoldersServiceServer(service *grpc.Server, h SyncFoldersHandler) {
	pb.RegisterSyncFoldersServiceServer(service, NewSyncFoldersServiceServer(h))
}

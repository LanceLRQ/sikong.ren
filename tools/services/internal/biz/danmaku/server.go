package danmaku

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	danmaku "launcher/internal/biz/danmaku/proto"
	"log"
	"net"
)

type DaemonServer struct {
	danmaku.UnimplementedDanmakuDaemonServer
}

func (s *DaemonServer) StartWatcher(ctx context.Context, req *danmaku.StartWatcherRequest) (*danmaku.StartWatcherResponse, error) {
	fmt.Printf("%s | %d", req.SessionId, req.RoomId)
	return &danmaku.StartWatcherResponse{
		Result: true,
		Message: "Done",
	}, nil
}

func RunServer() {
	listen, err := net.Listen("tcp", "0.0.0.0:8977")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	danmaku.RegisterDanmakuDaemonServer(s, &DaemonServer{})

	log.Printf("server listening at %v", "0.0.0.0:8977")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
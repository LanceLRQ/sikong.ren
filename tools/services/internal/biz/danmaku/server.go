package danmaku

import (
	"context"
	"google.golang.org/grpc"
	danmaku "launcher/internal/biz/danmaku/proto"
	"log"
	"net"
)

type DaemonServer struct {
	danmaku.UnimplementedDanmakuDaemonServer
}

// 启动监视器
func (s *DaemonServer) StartWatcher(ctx context.Context, req *danmaku.WatcherRequest) (*danmaku.WatcherResponse, error) {
	worker, err := NewLiveWatcher(int(req.RoomId), req.SessionId)
	if err != nil {
		return &danmaku.WatcherResponse{
			Result: false,
			Message: err.Error(),
		}, nil
	}
	log.Printf("[rpc] Connecting room: %d (%s)", req.RoomId, req.SessionId)
	// 启动监视器
	worker.Start()
	return &danmaku.WatcherResponse{
		Result: true,
		Message: "done",
	}, nil
}

// 关闭监视器
func (s *DaemonServer) StopWatcher(ctx context.Context, req *danmaku.WatcherRequest) (*danmaku.WatcherResponse, error) {
	rel, ok := WatcherPool.Load(int(req.RoomId))
	if !ok {
		return &danmaku.WatcherResponse {
			Result:  false,
			Message: "watcher is not exists",
		}, nil
	}
	worker := rel.(*LiveWatcher)
	if !worker.working {
		return &danmaku.WatcherResponse {
			Result:  false,
			Message: "watcher is not working",
		}, nil
	}
	log.Printf("[rpc] Closing room: %d", req.RoomId)
	// 停止
	worker.Stop()
	return &danmaku.WatcherResponse{
		Result:  true,
		Message: "done",
	}, nil
}

// 运行RPC服务器
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
package danmaku

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	danmaku "launcher/internal/biz/danmaku/proto"
	"launcher/internal/server"
	"log"
	"net"
)

type DaemonServer struct {
	danmaku.UnimplementedDanmakuDaemonServer
}

func checkAndGetWatcher (req *danmaku.WatcherRequest) (*LiveWatcher, error) {
	rel, ok := WatcherPool.Load(int(req.RoomId))
	if !ok {
		return nil, fmt.Errorf("watcher is not exists")
	}
	worker := rel.(*LiveWatcher)
	if !worker.working {
		return nil, fmt.Errorf("watcher is not working")
	}
	return worker, nil
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
	worker, err := checkAndGetWatcher(req)
	if err != nil {
		return &danmaku.WatcherResponse {
			Result:  false,
			Message: err.Error(),
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

func (s *DaemonServer) KeepAlive(ctx context.Context, req *danmaku.WatcherRequest) (*danmaku.WatcherResponse, error) {
	worker, err := checkAndGetWatcher(req)
	if err != nil {
		return &danmaku.WatcherResponse {
			Result:  false,
			Message: err.Error(),
		}, nil
	}
	worker.KeepAlive()
	return &danmaku.WatcherResponse{
		Result:  true,
		Message: "done",
	}, nil
}

func (s *DaemonServer) IsAlive(ctx context.Context, req *danmaku.WatcherRequest) (*danmaku.WatcherResponse, error) {
	rel, ok := WatcherPool.Load(int(req.RoomId))
	if !ok {
		return &danmaku.WatcherResponse {
			Result:  false,
			Message: "watcher is not exists",
		}, nil
	}
	worker := rel.(*LiveWatcher)
	expireAt, err := worker.ExpireAt()
	if err != nil {
		return &danmaku.WatcherResponse{
			Result:  worker.working,
			Message: "",
		}, nil
	}
	return &danmaku.WatcherResponse{
		Result:  worker.working,
		Message: fmt.Sprintf("expire at: %s", expireAt.Format("2006-01-02 15:04:05")),
	}, nil
}

// 运行RPC服务器
func RunServer(configFile string) error {
	// 载入配置
	err := server.LoadConfiguration(configFile)
	if err != nil { return err }
	// 启动监控器
	go WatcherMonitor()
	// 启动RPC服务
	listen, err := net.Listen("tcp", "0.0.0.0:8977")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	danmaku.RegisterDanmakuDaemonServer(s, &DaemonServer{})
	log.Printf("[rpc] Server listening at %v", "0.0.0.0:8977")
	if err = s.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
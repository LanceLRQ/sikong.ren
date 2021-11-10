package danmaku

import (
	"context"
	"google.golang.org/grpc"
	danmaku "launcher/internal/biz/danmaku/proto"
	"log"
	"time"
)

func RunClient () {
	conn, err := grpc.Dial("127.0.0.1:8977", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := danmaku.NewDanmakuDaemonClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StartWatcher(ctx, &danmaku.StartWatcherRequest{ SessionId: "abc", RoomId: 123})
	if err != nil {
		log.Fatalf("could not run: %v", err)
	}
	log.Printf("Result: %s", r.GetMessage())
}
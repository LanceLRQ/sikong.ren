package danmaku

import (
	"fmt"
	"launcher/internal/utils/gobilibili"
	"log"
	"sync"
	"time"
)

type LiveWatcher struct {
	roomId int
	sessionId string
	client *gobilibili.BiliBiliClient
	working bool
}

const MaxWatcherAliveTime = 600 * time.Second
const WatcherAliveMonitorSleepTime = 60 * time.Second

var WatcherPool sync.Map
var WatcherExpireAt sync.Map

func WatcherMonitor () {
	log.Println("[monitor] danmaku worker is ready")
	for {
		time.Sleep(WatcherAliveMonitorSleepTime)
		WatcherExpireAt.Range(func(key, value interface{}) bool {
			expireAt := value.(time.Time)
			if expireAt.Before(time.Now()) { // 如果过期，清理
				roomId := key.(int)
				t, ok := WatcherPool.Load(roomId)
				if ok {
					worker := t.(*LiveWatcher)
					if worker.working {
						log.Printf("[monitor] danmaku worker (room: %s) is inactive, force stop!")
						worker.Stop()
					}
				}
			}
			return true
		})
	}
}

func NewLiveWatcher(roomId int, sessionId string) (*LiveWatcher, error) {
	// 先看看池子里有没有watcher
	existsWatcher, ok := WatcherPool.Load(roomId)
	if ok {
		eWorker := existsWatcher.(*LiveWatcher)
		if eWorker.working {
			return eWorker, fmt.Errorf("bilibili live room (%d) is watching", roomId)
		}
	}
	// 没有就创建一个
	watcher := &LiveWatcher{
		roomId: roomId,
		sessionId: sessionId,
		client: gobilibili.NewBiliBiliClient(),
	}
	// 写入池子
	WatcherPool.Store(roomId, watcher)
	WatcherExpireAt.Store(roomId, time.Now().Add(MaxWatcherAliveTime))
	return watcher, nil
}

// 启动监视器
func (w *LiveWatcher) Start() {
	w.working = true
	w.registerHandler()
	go w.connect()
}

// 关闭监视器
func (w *LiveWatcher) Stop() {
	_ = w.client.Close()
	w.working = false
	WatcherPool.Delete(w.roomId)
	WatcherExpireAt.Delete(w.roomId)
}

// 获取Expire时间
func (w *LiveWatcher) ExpireAt() (*time.Time, error) {
	rel, ok := WatcherExpireAt.Load(w.roomId)
	expireAt := rel.(time.Time)
	if ok {
		return &expireAt, nil
	}
	return nil, fmt.Errorf("worker not exists")
}

// 心跳
func (w *LiveWatcher) KeepAlive() {
	WatcherExpireAt.Store(w.roomId, time.Now().Add(MaxWatcherAliveTime))
}

// 注册事件监听
func (w *LiveWatcher) registerHandler () {
	w.client.RegHandleFunc(gobilibili.CmdDanmuMsg, func(c *gobilibili.Context) bool {
		info := c.GetDanmuInfo()
		log.Printf("[%d]%d 说: %s\r\n", c.RoomID, info.UID, info.Text)
		//server.GameStatus.CurrentDanmaku = append([]gobilibili.DanmuInfo{info}, server.GameStatus.CurrentDanmaku...)
		//server.GameStatus.GlobalDanmaku = append(server.GameStatus.GlobalDanmaku, info)
		//if server.DrawingWebSocketServer != nil {
		//	server.DrawingWebSocketServer.Broadcast(nil, neffos.Message{
		//		Namespace: "drawing",
		//		Event:     "danmaku",
		//		Body:      []byte(utils.ObjectToJSONString(info, false)),
		//	})
		//}
		return false
	})
}

func (w *LiveWatcher) connect() {
	// 传入房间号
	err := w.client.ConnectServer(w.roomId)
	if err != nil {
		w.working = false
		log.Printf("[danmaku] Error: %s\n", err.Error())
	}
}
package server

import (
	"fmt"
	"github.com/kataras/neffos"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"launcher/internal/structs"
	"os"
	"path/filepath"
)

var Config structs.ServerConfiguration
// 全局websocket
var DrawingWebSocketServer *neffos.Server = nil

// 载入配置文件
func LoadConfiguration(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %s", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal configuration: %s", err.Error())
	}
	// parse store
	storePath, err := filepath.Abs(Config.Server.Storages)
	if err != nil {
		Config.Server.Storages = storePath
	}
	storeDir, err := os.Stat(storePath)
	if err != nil {
		if !os.IsExist(err) {
			err = os.Mkdir(storePath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("create store dir error: %s", err.Error())
			}
		} else {
			return fmt.Errorf("get store dir error: %s", err.Error())
		}
	} else {
		if !storeDir.IsDir() {
			return fmt.Errorf("store dir must be a directory: %s", storePath)
		}
	}
	return nil
}


//// 全局可用的游戏状态
//var GameStatus = structs.GameStatus {
//	CurrentId:      0, // <= 0 表示游戏没有开始
//	CurrentRiddle:  nil,
//	CurrentDanmaku: []gobilibili.DanmuInfo{},
//	GlobalDanmaku:  []gobilibili.DanmuInfo{},
//	DrawingHistory: []structs.DrawingOperation{},
//}

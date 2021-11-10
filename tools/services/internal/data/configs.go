package data

type ServerConfiguration struct {
	// Redis 配置
	Redis struct {
		// 监听IP
		Host string `yaml:"host" json:"host"`
		// 监听端口
		Port int `yaml:"port" json:"port"`
		// 数据库密码
		Password string `yaml:"password" json:"password"`
	} `yaml:"redis" json:"redis"`

	// MySQL 配置
	MySQL struct {
		// 监听IP
		Host string `yaml:"host" json:"host"`
		// 监听端口
		Port int `yaml:"port" json:"port"`
		// 数据库用户名
		User string `yaml:"user" json:"user"`
		// 数据库密码
		Password string `yaml:"password" json:"password"`
	} `yaml:"mysql" json:"mysql"`

	// 服务器配置  （优先级比-l低）
	Server struct {
		// 监听IP
		Listen string `yaml:"listen" json:"listen"`
		// 监听端口
		Port int `yaml:"port" json:"port,int"`
		// 存储目录
		Storages string `yaml:"storages" json:"storages"`
		// Redis 数据库位置
		RedisDB int `yaml:"redis_db" json:"redis_db"`
		// MySQL 数据库名称
		MySQLDBName string `yaml:"mysql_db_name" json:"mysql_db_name"`
		// Danmaku Server gRPC url
		DanmakuRPC string `yaml:"danmaku_rpc" json:"danmaku_rpc"`
	} `yaml:"server" json:"server"`

	Websocket struct {
		// 监听IP
		Listen string `yaml:"listen" json:"listen"`
		// 监听端口
		Port int `yaml:"port" json:"port,int"`
	}

	Danmaku struct {
		// 监听IP
		RPCListen string `yaml:"rpc_listen" json:"rpc_listen"`
		// 监听端口
		RPCPort int `yaml:"rpc_port" json:"rpc_port,int"`
	}

	// (废弃
	Bilibili struct {
		// 数据库IP
		RoomId int `yaml:"room_id" json:"room_id"`
	} `yaml:"bilibili" json:"bilibili"`

	// 调试模式
	DebugMode bool `yaml:"debug" json:"debug"`
}

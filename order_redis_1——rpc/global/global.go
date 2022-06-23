package global

import (
	"OrderTakeaway/config"
	"go.mongodb.org/mongo-driver/mongo"
)


//第五步 连接mysql并使用全局变量记录返回句柄提供其他文件调用

//全局变量
var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	MongoDB *mongo.Client






)
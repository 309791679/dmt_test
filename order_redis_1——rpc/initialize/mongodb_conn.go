package initialize

import (
	"OrderTakeaway/global"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

func IntiMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://root:123456x@192.168.38.128:27017/?directConnection=true") // # connect=direct这个参数设置是副本集我只连接其中一个主节点
	// Connect to MongoDB 连接数据库
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection 测试连接
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	global.MongoDB = c
}

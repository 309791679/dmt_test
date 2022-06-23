package main

import (
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "mm/grpc_"
	"time"

	//_ "github.com/gorilla/websocket"
	//"github.com/yeqown/go-qrcode"

	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"net/http"
)

var dtmServer = "consul://192.168.38.128:8500/dtmservice"
var orderRpcBusiServer = "consul://192.168.38.128:8500/user.rpc"

func main() {

	fmt.Print("连接成功\n")
	gin.DisableConsoleColor()
	//var UserRpc   proto.UsercenterClient
	//UserRpc = proto.NewUsercenterClient(NewClient(orderRpcBusiServer))
	//ctx := context.Background()
	router := gin.Default()
	router.POST("/ccc/", func(c *gin.Context) {


		//// 使用gin 调用rpc(不使用go-zero)
		//fmt.Println("准备执行",time.Now().String())
		//_, _ = UserRpc.QueryUser(ctx, &proto.UniversalUserRequest{})
		//fmt.Println("数据到达",time.Now().String())




		//// 使用gin 调用rpc(使用go-zero)
		//fmt.Println("准备执行",time.Now().String())
		//_, _ = UserRpc.QueryUser(ctx, &proto.UniversalUserRequest{})
		//fmt.Println("数据到达",time.Now().String())


		gid := dtmgrpc.MustGenGid(dtmServer)
		saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
			Add(orderRpcBusiServer+"/usercenter/QueryUser", orderRpcBusiServer+"/usercenter/CreateUser",
				&proto.UniversalUserRequest{Pattern: 2})

		saga.WaitResult =true
		saga.TimeoutToFail =90
		saga.RequestTimeout =90
		saga.RetryInterval =-90
		fmt.Println("调用接口准备提交")
		err_ := saga.Submit()

		fmt.Println("执行完毕",time.Now().String())
		fmt.Println(err_)
		c.String(http.StatusOK, "错误-")
	})

	s := &http.Server{
		Addr:           ":7788",
		Handler:        router,
		ReadTimeout:    6000 * time.Second,
		WriteTimeout:   6000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go GrpcConn(zrpc.RpcClientConf{Target: orderRpcBusiServer})
	s.ListenAndServe()
	//router.Run(":7788")

}
func NewClient(Content string) grpc.ClientConnInterface {
	Conn, err := grpc.Dial(Content, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(any(err))
	}

	return Conn
}
func GrpcConn(conn zrpc.RpcClientConf) {

	client := zrpc.MustNewClient(conn)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	fmt.Println(client.Conn().Target())

}
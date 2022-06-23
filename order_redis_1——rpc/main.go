package main

import (
	"OrderTakeaway/global"
	"OrderTakeaway/handler"
	"OrderTakeaway/utils"
	"OrderTakeaway/proto"
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"

	"OrderTakeaway/initialize"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func main() {



	initialize.InitConfig()

	initialize.IntiMongoDB()


	

	ip := flag.String("ip", "192.168.0.112", "IP地址")
	port := flag.Int("port", 0, "端口号")
	flag.Parse()

	if *port == 0 {
		//50051
		*port, _ = utils.GetFeePort()
	}

	server := grpc.NewServer()

	grpc_.RegisterUsercenterServer(
		server,
		&handler.UsercenterServer{},
	)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		panic("failed to lisen:" + err.Error())
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	//注册服务
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	client, err := api.NewClient(cfg)

	//生成对应的检查对象

	//配置健康检查的具体事项
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.0.112:%d", *port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	servuceID := fmt.Sprintf("%s", uuid.NewV4())
	registration := api.AgentServiceRegistration{
		Name:    global.ServerConfig.Name,
		ID:      servuceID,
		Port:    *port,
		Tags:    []string{"imooc", "bobby", "user", "srv"},
		Address: "192.168.0.112",
		Check:   check,
	}

	client.Agent().ServiceRegister(&registration)

	go func() {
		//正常情况下 这里会阻塞下方代码不会执行
		//解决方法使用协程加闭包函数
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()


	//接受系统终止信号，销毁前退出consul
	//非终止信号堵塞
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//_=c.Shutdown() 此时调用
	if err = client.Agent().ServiceDeregister(servuceID); err != nil {
		zap.S().Infof("注销失败")
	}
	zap.S().Infof("注销成功")
}


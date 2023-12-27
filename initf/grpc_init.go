package initf

import (
	"RyzeSCA/global"
	"RyzeSCA/grpc/server/ryzesca_server"

	"RyzeSCA/grpc/server/ryzesca"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func GrpcServerInit() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		global.Logger.Fatalf("http server 启动失败 请检测端口是否被占用%d", viper.GetInt("server.port"))
		fmt.Println("http server 启动失败 请检测端口是否被占用", viper.GetInt("server.port"))
		return
	}
	var recvSize = 1024 * 1024 * 1024 * 1024 * 1024
	var sendSize = 1024 * 1024 * 1024 * 1024 * 1024
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(recvSize),
		grpc.MaxSendMsgSize(sendSize),
	}
	grpcServer := grpc.NewServer(options...)

	ryzesca.RegisterRyzescaServer(grpcServer, ryzesca_server.RyzescaServer{})

	global.Logger.Info("grpc server start")
	if err := grpcServer.Serve(lis); err != nil {
		global.Logger.Fatalf("启动grpc server失败")
		fmt.Println("启动grpc server失败")
	}
}

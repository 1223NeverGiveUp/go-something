
//获取创建的客户端以及服务端
package grpcManager

import (
	"buildGrpc/grpcManager/client"
	"buildGrpc/grpcManager/server"
)

type GrpcManager struct {
	GrpcClient *client.GRPCClient
	//GrpcServer server.GrpcServer
	servers map[string]server.GrpcServer
}

func (grpcManager * GrpcManager)InitServer(localip,localport string){
	if grpcManager.servers == nil{
		grpcManager.servers = make(map[string]server.GrpcServer,0)
	}
	var addr = localip + ":" + localport
	//查找map表中是否已经注册服务，没有则添加新的服务
	if _,ok := grpcManager.servers[addr];!ok{
		var server = server.GrpcServer{Addr:addr}
		server.Init()
		grpcManager.servers[addr] = server
	}
}
func (grpcManager * GrpcManager)InitClient(){
	grpcManager.GrpcClient = &client.GRPCClient{}
	grpcManager.GrpcClient.Init()
}

func (grpcManager * GrpcManager)GetClient()(client * client.GRPCClient){
	return grpcManager.GrpcClient
}

func (grpcManager * GrpcManager) RunServers(){
	for _,v := range grpcManager.servers{
		v.RunGrpcServer()
	}
}


package server

import (
	"net"
	"google.golang.org/grpc"
	netconfig "buildGrpc/grpcManager/proto/NetConfig"
	"math"
	"log"
	"buildGrpc/grpcManager/server/NetConfigServer"
	"fmt"
)
type GrpcServer struct {
	Addr string
	pNetConfigTranServer * NetConfigServer.NetConfigServer
}

//初始化服务端
func (grpcServer * GrpcServer)Init(){
	NetConfig := NetConfigServer.NewNetConfigServer()
	grpcServer.pNetConfigTranServer = &NetConfig
}
func (grpcServer * GrpcServer)RunGrpcServer(){
	lis,err := net.Listen("tcp",grpcServer.Addr)
	if err != nil{
		log.Println("server print error:%s",err)
	}
	optionMaxRecvMsgSize := grpc.MaxRecvMsgSize(math.MaxInt32)
	s:=grpc.NewServer(optionMaxRecvMsgSize)
	//将服务注册到grpc服务中
	//netConfigTranServer := NetConfigServer.NewNetConfigServer()
	netconfig.RegisterNetConfigTranServer(s,grpcServer.pNetConfigTranServer)
	fmt.Println("启动监听[",grpcServer.Addr,"]...")
	s.Serve(lis)
}

//不需要
func RunBuildGrpcServer(localip string,localport string){
	lis,err := net.Listen("tcp",localip+":"+localport)
	if err != nil{
		log.Println("server print error:%s",err)
	}
	optionMaxRecvMsgSize := grpc.MaxRecvMsgSize(math.MaxInt32)
	s:=grpc.NewServer(optionMaxRecvMsgSize)
	//将服务注册到grpc服务中
	netConfigTranServer := NetConfigServer.NewNetConfigServer()
	netconfig.RegisterNetConfigTranServer(s,&netConfigTranServer)
	fmt.Println("启动监听[",localip,localport,"]...")
	s.Serve(lis)
}
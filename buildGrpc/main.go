package main

import (
	"buildGrpc/grpcManager"
)

func main() {
	//server.RunBuildGrpcServer("127.0.0.1","60065")
	var gm = grpcManager.GrpcManager{}
	gm.InitServer("127.0.0.1","60065")
	gm.RunServers()
}

package main

import (
	"buildGrpc/grpcManager"
	"time"
	"fmt"
)

func main(){
	//var NetConfigClient client.GRPCClient
	//NetConfigClient.Init()
	//for {
	//	now := time.Now()
	//	NetConfigClient.NetConfigClientImpl.PushNetConfig("127.0.0.1", "60065", 0, "ma.xml", nil)
	//	fmt.Println("调用延迟:",time.Now().Sub(now))
	//	fmt.Println("wait 5s...")
	//	time.Sleep(time.Second * 5)
	//
	//}
	var gm = grpcManager.GrpcManager{}
	gm.InitClient()
	var NetConfigClient = gm.GetClient()
	for {
		now := time.Now()
		NetConfigClient.NetConfigClientImpl.PushNetConfig("127.0.0.1", "60065", 0, "ma.xml", nil)
		fmt.Println("调用延迟:",time.Now().Sub(now))
		fmt.Println("wait 5s...")
		time.Sleep(time.Second * 5)
	}
}

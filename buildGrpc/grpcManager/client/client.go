package client

import (
	"buildGrpc/grpcManager/client/NetConfigClient"
)


//只有实现
type GRPCClient struct {
	NetConfigClientImpl NetConfigClient.NetConfigClientImpl
}
func (grpcClient *GRPCClient)Init(){
	//初始化实现
	grpcClient.NetConfigClientImpl = NetConfigClient.NewNetConfigClientImpl()
}
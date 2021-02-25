package NetConfigClient

import(
	netConfig "buildGrpc/grpcManager/proto/NetConfig"
	"context"
	"google.golang.org/grpc"
	"fmt"
	"errors"
	"buildGrpc/grpcManager/client/grpcConn"
)
type netConfigClient struct {
	netConfigTranClient netConfig.NetConfigTranClient
	conn *grpc.ClientConn
}

func (configClient *netConfigClient)PutConfig(addr string,cfg *netConfig.CfgMessage)(err error){
	fmt.Println("接口层:call PutConfig接口")
	//获取连接
	conn, _ := grpcConn.GetGRPCConn(addr, 10)
	if conn == nil{
		return errors.New("获取grpc连接失败!")
	}
	configClient.netConfigTranClient =  netConfig.NewNetConfigTranClient(conn)
	configClient.netConfigTranClient.PushNetConfig(context.Background(),cfg)
	return
}

func (configClient *netConfigClient)PullConfig(addr string,cfg *netConfig.CfgMessage)(err error){
	fmt.Println("client 调用PullConfig接口")
	//获取连接
	conn, _ := grpcConn.GetGRPCConn(addr, 10)
	if conn == nil{
		return errors.New("获取grpc连接失败!")
	}
	configClient.netConfigTranClient =  netConfig.NewNetConfigTranClient(conn)
	configClient.netConfigTranClient.PullNetConfig(context.Background(),cfg)
	return
}

package NetConfigServer

import (
	netConfig "buildGrpc/grpcManager/proto/NetConfig"
	"golang.org/x/net/context"
	"fmt"
)

//实现NetConfig proto定义的服务器端接口
type NetConfigServer struct {
	pNetConfigServerData * netConfigServerData
}
func NewNetConfigServer()(NetConfigServer){
	data := newNetConfigServerData()
	return NetConfigServer{pNetConfigServerData:&data}
}
func (config *NetConfigServer)PushNetConfig(c context.Context, cfg *netConfig.CfgMessage) (status *netConfig.NetConfigStatus, err error){
	fmt.Println("1.接口层收到消息.")
	config.pNetConfigServerData.PushNetConfig(cfg)
	return &netConfig.NetConfigStatus{},nil
}

func (config *NetConfigServer)PullNetConfig(c context.Context, cfg *netConfig.CfgMessage) (retcfg *netConfig.CfgMessage, err error){
	fmt.Println("PullNetConfig call")
	fmt.Printf("recv data:%#v\n",*cfg)
	config.pNetConfigServerData.PullNetConfig(cfg)
	return &netConfig.CfgMessage{},nil
}
func (config *NetConfigServer)PushPartNetConfig(ctx context.Context, partCfg *netConfig.CfgMessage) (retcfg *netConfig.NetConfigStatus,err error){
	fmt.Println("PushPartNetConfig call")
	config.pNetConfigServerData.PushPartNetConfig(partCfg)
	return &netConfig.NetConfigStatus{},nil
}

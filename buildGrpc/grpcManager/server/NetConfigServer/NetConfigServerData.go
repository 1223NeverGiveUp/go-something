//数据处理
package NetConfigServer

import (
	netConfig "buildGrpc/grpcManager/proto/NetConfig"
	"fmt"
)
type netConfigServerData struct {
	pNetConfigServerImpl * netConfigServerImpl
}

func newNetConfigServerData()(netConfigServerData){
	return netConfigServerData{pNetConfigServerImpl:&netConfigServerImpl{}}
}
func (config *netConfigServerData)PushNetConfig(cfg *netConfig.CfgMessage) (status *netConfig.NetConfigStatus, err error){
	fmt.Println("2.数据层处理消息.")
	config.pNetConfigServerImpl.PushNetConfig(cfg.Type,cfg.Filename,cfg.Config)
	return
}

func (config *netConfigServerData)PullNetConfig(cfg *netConfig.CfgMessage) (retcfg *netConfig.CfgMessage, err error){
	return
}
func (netConfig *netConfigServerData)PushPartNetConfig(*netConfig.CfgMessage) (retcfg *netConfig.NetConfigStatus,err error){
	return
}

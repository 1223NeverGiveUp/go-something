package NetConfigClient

import (
	"fmt"
	"buildGrpc/grpcManager/proto/NetConfig"
)

type netConfigClientData struct {
	Addr string
	pNetConfigClient *netConfigClient
}

func NewNetConfigServerData()(netConfigClientData){
	return netConfigClientData{pNetConfigClient:&netConfigClient{}}
}
func (clientData *netConfigClientData)PushNetConfig(addr string,fileType int32,filename string,data []byte) (code int32,errMsg string, err error){
	fmt.Println("数据层:netconfig handle data")
	clientData.Addr = addr
	err = clientData.pNetConfigClient.PutConfig(addr,&NetConfig.CfgMessage{Filename:filename})
	if err!= nil{
		fmt.Println("调用PushNetConfig error:",err)
	}

	//数据封装
	//接口调用
	//错误处理
	return
}

func (clientData *netConfigClientData)PullNetConfig(addr string,fileType int32,filename string,data []byte) (code int32,errMsg string, err error){
	return
}
func (clientData *netConfigClientData)PushPartNetConfig(addr string,fileType int32,filename string,data []byte) (code int32,errMsg string,err error){
	return
}

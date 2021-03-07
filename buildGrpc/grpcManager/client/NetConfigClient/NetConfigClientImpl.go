package NetConfigClient

import "fmt"

type NetConfigClientImpl struct {
	pNetConfigClientData * netConfigClientData
}

func NewNetConfigClientImpl()(NetConfigClientImpl){
	var pData = NewNetConfigServerData()
	return NetConfigClientImpl{pNetConfigClientData:&pData}
}
func (impl *NetConfigClientImpl)PushNetConfig(ip,port string,fileType int32,filename string,data []byte) (code int32,errMsg string, err error){
	fmt.Println("业务层:netConfig handle impl")
	impl.pNetConfigClientData.PushNetConfig(ip+":"+port,fileType,filename,data)
	return
}

func (impl *NetConfigClientImpl)PullNetConfig(ip,port string,fileType int32,filename string,websiteId string) (fileName string,data []byte, err error){
	return
}
func (impl *NetConfigClientImpl)PushPartNetConfig(ip,port string,fileType int32,filename string,data []byte,websiteId string) (code int32,errMsg string,err error){
	return
}
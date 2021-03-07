//具体的业务处理
package NetConfigServer

import "fmt"

type netConfigServerImpl struct {

}

func (config *netConfigServerImpl)PushNetConfig(fileType int32,filename string,data []byte) (code int32,errMsg string, err error){
	fmt.Println("3.业务层:实现具体业务.")
	return
}

func (config *netConfigServerImpl)PullNetConfig(fileType int32,filename string,websiteId string) (fileName string,data []byte, err error){
	return
}
func (netConfig *netConfigServerImpl)PushPartNetConfig(fileType int32,filename string,data []byte,websiteId string) (code int32,errMsg string,err error){
	return
}

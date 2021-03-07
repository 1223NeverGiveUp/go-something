package grpcConn

import (
	"google.golang.org/grpc"
	"context"
	"time"
	"sync"
	"google.golang.org/grpc/connectivity"
	"log"
	"fmt"
)

var mapLock sync.Mutex
var grpcConn map[string]*grpc.ClientConn

func init(){
	//初始化grpc连接存储map
	grpcConn = make(map[string]*grpc.ClientConn,0)
}

func GetGRPCConn(addr string,timeout int)(conn *grpc.ClientConn,err error){
	mapLock.Lock()
	defer mapLock.Unlock()
	var ok bool
	//怎么检测addr格式是否符合ip:port,不需要检测,如果不符合规则，直接获取空的grpc连接对象
	if conn,ok = grpcConn[addr];!ok{
		//设置grpc连接的超时时间
		ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		conn,err = grpc.DialContext(ctx,addr,grpc.WithInsecure(),)
		if err != nil{
			return nil,err
		}
		//加入到map中
		grpcConn[addr] = conn
		return
	}else{
		fmt.Println("conn state:",conn.GetState().String())
		//检查状态
		switch conn.GetState() {
		case connectivity.Ready:
			fallthrough
		case connectivity.Connecting:
			fallthrough
		case connectivity.Idle:
		case connectivity.Shutdown:
			fallthrough
		case connectivity.TransientFailure:
			conn.Close()
		default:
			log.Println("Invalid-State")
		}
	}
	return
}
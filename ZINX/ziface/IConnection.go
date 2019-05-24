package ziface

import "net"

/*
	抽象链接层,定义连接曾接口
*/
type IConnection interface {
	//开启连接
	Start()
	//停止连接
	Stop()
	//获取ID
	GetConnID() uint32
	//获取conn原声tcp socket 陶介子
	GetTCPConnection() *net.TCPConn
	//获取远程客户端的IP地址
	GetRemoteAddr() net.Addr
	//发送数据给合客户端
	Send(data []byte, cnt int) error
}

//业务处理方法
//type HandleFunc func(*net.TCPConn,[]byte,int)error
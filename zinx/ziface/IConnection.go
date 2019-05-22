package ziface

import "net"

/*
抽象曾
*/
type IConnection interface {
	//启动连接
	Start()
	//停止连接
	Stop()
	//获取连接
	GetConnID() uint32
	//获取conn的原声socket陶介子
	GetTCPconnection() *net.Conn
	//获取远程客户端IP地址
	GetRemoteAddr() *net.Addr
	//发送数据给对方客户端
	Send(data []byte) error
}

//业务处理方法
type HandleFunc func(*net.TCPConn, []byte, int) error

package net

import (
	"net"
	"review/zinx/ziface"
)

// 具体的TCP连接模块
type Connection struct {
	//当前连接的原声陶介子
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前连接状态
	isClosed bool
	//当前连接锁绑定的方法
	handleAPI ziface.HandleFunc
}

/*
  初始化方法
 */

func NewConnection(conn *net.TCPConn, connID uint32, callback ziface.HandleFunc) ziface.IConnection{
	c :=&Connection{
		Conn:conn,
		ConnID:connID,
		isClosed:false,
		handleAPI:callback,
	}

	return c
}

//启动连接
func (c *Connection)Start(){

}
//停止连接
func (c *Connection)Stop(){

}
//获取连接
func (c *Connection)GetConnID() uint32{
	return 0
}
//获取conn的原声socket陶介子
func (c *Connection)GetTCPconnection() *net.Conn{
	return nil
}
//获取远程客户端IP地址
func (c *Connection)GetRemoteAddr() *net.Addr{
	return nil
}
//发送数据给对方客户端
func (c *Connection)Send(data []byte) error{
	return nil
}
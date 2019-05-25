package net

import "review/zinx/ziface"

type Request struct {
	//具体链接模块
	conn ziface.IConnection
	//客户端发送的数据
	Msg ziface.IMessage
}

//创建具体Request对象函数
func NewRequest(conn ziface.IConnection, msg ziface.IMessage) ziface.IRequest {
	req := &Request{
		conn: conn,
		Msg:  msg,
	}
	return req
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetMsg() ziface.IMessage {
	return r.Msg
}

package znet

import "zinx/ZINX/ziface"

//定义请求类型结构体
type Request struct {
	//当前链接
	conn ziface.IConnection
	//传递的数据
	data []byte
	//数据的长度
	len int
}

//创建具体的请求结构体
func NewRequest(conn ziface.IConnection, data []byte, len int) ziface.IRequest {
	r := &Request{
		conn: conn,
		data: data,
		len:  len,
	}
	return r
}

//获取当前链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//获取当前数据
func (r *Request) GetData() []byte {
	return r.data
}

//获当前数据长度
func (r *Request) GetLen() int {
	return r.len
}

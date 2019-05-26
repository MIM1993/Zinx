package main

import (
	"review/zinx/net"
	"review/zinx/ziface"
	"fmt"
)

//创建路由控制器
type PingRouter struct {
	net.BaseRouter
}

//业务一的方法
func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router PingRouter Handle project one")
	err := request.GetConnection().Send(200, []byte("project one ..."))
	if err != nil {
		fmt.Println("call back ping err:", err)
		return
	}
}

//业务二路由
type pongRouter struct {
	net.BaseRouter
}

//业务二的方法
func (r *pongRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router pongRouter Handle project two")
	err := request.GetConnection().Send(201, []byte("project two ..."))
	if err != nil {
		fmt.Println("call back ping err:", err)
		return
	}
}

//创建链接之后运行的钩子函数
func DoConnBegin(conn ziface.IConnection) {
	fmt.Println("====>DoConnBegin ....")
	//链接成功后发送给客户端
	if err := conn.Send(277, []byte("Hello welcome to ZinxServer...")); err != nil {
		fmt.Println("DoConnBegin call err:", err)
		return
	}
}

//销毁链接之前用运行的Hook函数
func DoConnLost(conn ziface.IConnection) {
	fmt.Println("====>DoConnLost ....")
	fmt.Println("Conn Id =", conn.GetConnID(), "is Lost!")
}

func main() {
	//创建一个zinx server对象
	s := net.NewServer("zinx v0.6")

	//注册一个创建连接之后的方法业务
	s.AddOnConnStart(DoConnBegin)
	//注册一个销毁链接之前的方法业务
	s.AddOnConnStop(DoConnLost)

	//添加router
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &pongRouter{})
	//让server对象 启动服务
	s.Serve()

	return
}

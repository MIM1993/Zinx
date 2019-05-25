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

func main() {
	//创建一个zinx server对象
	s := net.NewServer("zinx v0.6")

	//添加router
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &pongRouter{})
	//让server对象 启动服务
	s.Serve()

	return
}

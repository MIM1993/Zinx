package net

import (
	"review/zinx/ziface"
	"net"
	"fmt"
	"review/zinx/config"
)

type Server struct {
	//服务器协议
	IPVersion string
	//IP
	IP string
	//端口号
	Port int
	//服务器名
	Name string
	//路由属性
	//Router ziface.IRouter

	//录用map集合属性
	MsgHandler ziface.IMsgHandler
}

//初始化服务器方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		IPVersion:  "tcp4",
		IP:         config.GlobalObject.Host,
		Port:       config.GlobalObject.Port,
		Name:       config.GlobalObject.Name,
		MsgHandler: NewMsgHandler(),
	}
	return s
}


//启动服务器
func (s *Server) Start() {
	//创建陶介子
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr err:", err)
		return
	}
	//监听服务器地址
	Listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IPVersion, "err", err)
		return
	}

	//生成id
	var cid uint32
	cid = 0
	//3 阻塞等待客户端发送请求
	go func() {
		for {
			//阻塞等待客户端发送请求
			conn, err := Listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err :", err)
				return
			}
			//与客户端进行连接
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()

}

//停止服务器
func (s *Server) Stop() {

}

//运行服务器
func (s *Server) Serve() {
	//开启服务;
	s.Start()

	//TODO  做一些其他的扩展
	//阻塞//告诉CPU不再需要处理的，节省cpu资源
	select {} //main函数不退出  //main函数 阻塞在这
}

//添加路由函数
func (s *Server) AddRouter(MsgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(MsgId, router)
}

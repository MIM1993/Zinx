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

	//链接管理模块
	connMgr ziface.IConnManager

	//该server创建之后自动调用的Hook函数
	OnConnStart func(conn ziface.IConnection)
	//该server销毁前自动调用的Hook函数
	OnConnStop func(conn ziface.IConnection)
}

//初始化服务器方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		IPVersion:  "tcp4",
		IP:         config.GlobalObject.Host,
		Port:       config.GlobalObject.Port,
		Name:       config.GlobalObject.Name,
		MsgHandler: NewMsgHandler(),
		connMgr:    NewConnManager(),
	}
	return s
}

//启动服务器
func (s *Server) Start() {
	fmt.Printf("[start] Server Linstenner at IP :%s, Port :%d, is starting..\n", s.IP, s.Port)

	//开启Worker线程池
	s.MsgHandler.StartWorkerPool()

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
			//判断当前server链接数量是否已经最大值
			fmt.Println(config.GlobalObject.MaxConn)
			fmt.Println(s.connMgr.Len())
			if s.connMgr.Len() >= config.GlobalObject.MaxConn {
				fmt.Println("---> Too many Connection MAxConn = ", config.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//生成对应这个链接的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			//放入NewConnection()中
			//s.connMgr.Add(dealConn)
			cid++

			//开启groutine 开始进行链接后的业务
			go dealConn.Start()
		}
	}()

}

//停止服务器
func (s *Server) Stop() {
	//服务器停止  应该清空当前全部的链接
	s.connMgr.ClearConn()
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

//获取链接管理模块
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.connMgr
}

//注册函数--->创建链接之后需要调用的HOOK方法
func (s *Server) AddOnConnStart(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

//注册函数--->销毁链接之前需要调用的HOOK方法
func (s *Server) AddOnConnStop(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

//调用创建链接后运行的方法
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---->Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

//调用销毁链接前运行的方法
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---->Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}

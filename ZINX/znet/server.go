package znet

import (
	"fmt"
	"net"
	"log"
	"zinx/ZINX/ziface"
	"zinx/ZINX/config"
)

//iServer接口实现，定义一个Server服务类
type Server struct {
	//服务器名称
	Name string
	//tcp4、tcp6、tcp  协议名
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口号
	Port int
	//路由属性
	Router ziface.IRouter
}

//创建一个服务器句柄
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      config.Conf.Name,
		IPVersion: "tcp4",
		IP:        config.Conf.Host,
		Port:      config.Conf.Port,
		Router:    nil,
	}
	return s
}

//开启网络服务
func (s *Server) Start() {
	//打印开始信息
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	//开启go程进行服务端监听
	go func() {
		//1、获取一个TCP地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Fatal("resolve tcp addr err:", err)
		}

		//2、监听服务器地址
		Listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Fatal("listen", s.IPVersion, "err", err)
		}

		//已经监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")

		//设置链接ID
		var cid uint32
		cid = 0

		//3、启动server网络连接业务   循还
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := Listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			//我们这里暂时做一个最大512字节的回显服务
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()

		}
	}()
}

//关闭网络服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

//服务器运行操作
func (s *Server) Serve() {
	//开启服务
	s.Start()

	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

package net

import (
	"review/zinx/ziface"
	"net"
	"fmt"
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
}

//初始化服务器方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Name:      name,
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

	//3 阻塞等待客户端发送请求
	go func() {
		for {
			//阻塞等待客户端发送请求
			conn, err := Listener.Accept()
			if err != nil {
				fmt.Println("Accept err :", err)
				return
			}
			//与客户端进行连接
			go func() {
				buf := make([]byte, 512)
				for {
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err:", err)
						break
					}

					fmt.Printf("recv client buf %s, cnt = %d\n", buf, cnt)

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back err :", err)
						continue
					}
				}
			}()
		}
	}()

}

//停止服务器
func (s *Server) Stop() {

}

//运行服务器
func (s *Server) Serve() {
	//开启服务
	s.Start()

	//阻塞
	select {}
}

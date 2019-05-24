package znet

import (
	"net"
	"zinx/ZINX/ziface"
	"fmt"
	"zinx/ZINX/config"
)

//具体的TCP连接模块
type Connection struct {
	//当前连接的原生陶介子
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前在连接状态
	isclosed bool
	//路由属性
	Router ziface.IRouter
}

//生成TCP连接模块
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) ziface.IConnection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isclosed: false,
		Router:   router,
	}
	return c
}

//读取业务
func (c *Connection) StartReader() {
	//从客户端读取数据
	fmt.Println("Reader Groutine is running ...")
	defer c.Stop()
	//循环读取
	buf := make([]byte, config.Conf.MaxPackageSize)
	for {
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error:", err)
			break
		}

		//将数据和链接进行绑定
		req := NewRequest(c, buf, cnt)

		//调用router模板函数
		go func() {
			//c.Router.PreHandle(req)
			c.Router.Handle(req)
			//c.Router.PostHandle(req)
		}()
	}
}

//开启连接
func (c *Connection) Start() {
	fmt.Println("conn start ...ID = ", c.ConnID, )
	//读业务
	go c.StartReader()
}

//停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn stop ID = ", c.ConnID)

	if c.isclosed == true {
		return
	}

	c.isclosed = true

	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Conn Close err:", err)
		return
	}

}

//获取ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取conn原声tcp socket 陶介子
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取远程客户端的IP地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据给合客户端
func (c *Connection) Send(data []byte, cnt int) error {
	if _, err := c.Conn.Write(data[:cnt]); err != nil {
		fmt.Println("Send msg err:", err)
		return err
	}
	return nil
}

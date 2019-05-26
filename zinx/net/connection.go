package net

import (
	"net"
	"review/zinx/ziface"
	"fmt"
	"errors"
	"io"
	"review/zinx/config"
)

// 具体的TCP连接模块
type Connection struct {
	//当前链接是属于那个server创建的
	server ziface.IServer
	//当前连接的原声陶介子
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前连接状态
	isClosed bool
	//当前连接锁绑定的方法
	//handleAPI ziface.HandleFunc
	//路由属性
	MsgHandler ziface.IMsgHandler

	//添加一个reader和writer及通信的channel
	msgChan chan []byte

	//创建一个channel,用于reader通知writer,conn链接已经关闭,需要退出,释放系统资源
	writerExitChan chan bool
}

/*
  初始化方法
 */

//初始化生成链接模块
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, MsgHandler ziface.IMsgHandler) ziface.IConnection {
	c := &Connection{
		server:   server,
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI: callback,
		MsgHandler:     MsgHandler,
		msgChan:        make(chan []byte),
		writerExitChan: make(chan bool),
	}

	//当创建一个链接时,添加到链接管理器中
	c.server.GetConnMgr().Add(c)

	return c
}

//处理conn读取数据的goroutine
func (c *Connection) StartReader() {
	//从对端读取数据
	fmt.Println("Reader Goroutine is running...")
	//关闭时打印
	defer fmt.Println("connID=", c.ConnID, "Reader is exit ,remote addr is = ", c.GetRemoteAddr().String())
	//调用停止函数
	defer c.Stop()

	//循环读取数据
	for {
		//创建封包对象
		dp := NewDataPack()

		//读取客户端发送过来的数据头部
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			fmt.Println("read msg head err:", err)
			break //头部读取错误
		}

		//根据头部 获取数据的长度，进行第二次读取
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unPack err", err)
			break
		}

		//根据头部 获取数据的长度，进行第二次读取
		var data []byte
		if msg.GetMsgLen() > 0 {
			//有内容
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("read msg data err", err)
				break
			}
		}
		//调用函数,赋值
		msg.SetData(data)

		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req := NewRequest(c, msg)

		//调用用户传递进来的自定义的业务模板
		//将req数据交给worker工作池来处理
		if config.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(req)
		} else {
			go c.MsgHandler.DoMsgHandler(req)
		}

		//清空buf
		//buf = []byte{}
	}
}

//写消息的groutine ,专门负责给客户端发送消息
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Groutine is Start...]")

	defer fmt.Println("Writer Groutine is Stop")
	//IO多路复用
	for {
		select {
		case data := <-c.msgChan:
			//用数据发送过来
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data err :", err)
				return
			}

		case <-c.writerExitChan:
			return
		}
	}

}

//启动连接
func (c *Connection) Start() {
	fmt.Println("conn start .... ID = ", c.ConnID)

	//读业务
	go c.StartReader()

	//写业务
	go c.StartWriter()

	//调用创建链接后的Hook函数
	c.server.CallOnConnStart(c)
}

//停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn stop ... ID = ", c.ConnID)

	//调用销毁链接前的Hook函数
	c.server.CallOnConnStop(c)

	//回收工作
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	//当StartReader执行了Stop()时,用于通知StartWriter对端Conn已关闭
	c.writerExitChan <- true

	//关闭源生套接字
	_ = c.Conn.Close()

	//从链接管理模块中删除当前链接
	c.server.GetConnMgr().Remove(c.ConnID)

	//资源回收,释放资源
	close(c.msgChan)
	close(c.writerExitChan)

}

//获取连接
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取conn的原声socket陶介子
func (c *Connection) GetTCPconnection() *net.TCPConn {
	return c.Conn
}

//获取远程客户端IP地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据给对方客户端
func (c *Connection) Send(msgId uint32, msgData []byte) error {
	//判断链接状态
	if c.isClosed == true {
		return errors.New("Connection closed ...send Msg")
	}

	//封装Msg
	dp := NewDataPack()                  //创建封包对象
	msg := NewMsgPackage(msgId, msgData) //创建msg消息对象

	//开始封装
	binaryMsg, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("Pack error msgId=", msgId, "err:", err)
		return err
	}

	//将binaryMsg二进制数据流发送到对端
	//if _, err := c.Conn.Write(binaryMsg); err != nil {
	//	fmt.Println("send buf error:", err)
	//	return err
	//}

	//将需要发送的数据发送到channel 让writer写给客户端
	c.msgChan <- binaryMsg

	return nil
}

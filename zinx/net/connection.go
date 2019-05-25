package net

import (
	"net"
	"review/zinx/ziface"
	"fmt"
	"errors"
	"io"
)

// 具体的TCP连接模块
type Connection struct {
	//当前连接的原声陶介子
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前连接状态
	isClosed bool
	//当前连接锁绑定的方法
	handleAPI ziface.HandleFunc
	//路由属性
	MsgHandler ziface.IMsgHandler
}

/*
  初始化方法
 */

func NewConnection(conn *net.TCPConn, connID uint32, MsgHandler ziface.IMsgHandler) ziface.IConnection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI: callback,
		MsgHandler: MsgHandler,
	}
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
		go c.MsgHandler.DoMsgHandler(req)

		//清空buf
		//buf = []byte{}
	}
}

//启动连接
func (c *Connection) Start() {
	fmt.Println("conn start .... ID = ", c.ConnID)

	//读业务
	go c.StartReader()

	//TODO:写业务
}

//停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn stop ... ID = ", c.ConnID)

	//回收工作
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	_ = c.Conn.Close()
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
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("send buf error:", err)
		return err
	}
	return nil
}

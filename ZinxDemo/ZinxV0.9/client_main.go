package main

import (
	"net"
	"fmt"
	"time"
	net2 "review/zinx/net"
	"io"
)

//模拟客户端
func main() {
	fmt.Println("client start ...")
	time.Sleep(1 * time.Second)
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client link err:", err)
		return
	}
	//起循环
	for {
		//写
		dp := net2.NewDataPack()

		binaryMsg, err := dp.Pack(net2.NewMsgPackage(1, []byte("Zinx 0.6 project one...")))
		if err != nil {
			fmt.Println("Pack err", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write err", err)
			return
		}

		//服务器就会给我们返回一个 消息ID 1 的 pingping TLV格式的二进制数据
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("client unpack msghead err", err)
			return
		}

		//根据头的长度进行第二次读取
		msgHead, err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("UnPack err", err)
			return
		}
		if msgHead.GetMsgLen() > 0 {
			//读取包体
			msg := msgHead.(*net2.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err", err)
				return
			}

			fmt.Println("--->Revc Server Msg: id = ", msg.Id, " dataLen = ", msg.Datalen, " data = ", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}

}

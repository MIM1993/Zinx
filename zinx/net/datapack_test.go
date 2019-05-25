package net

import (
	"testing"
	"fmt"
	"net"
	"io"
)

func TestDataPack(t *testing.T) {

	fmt.Println("Test dataPack ......")
	/*
		模拟一个server
		收到二进制流  进行解包操作
	*/
	Listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Server Listen err :", err)
		return
	}
	//defer Listener.Close()
	//Accept
	go func() {
		for {
			conn, err := Listener.Accept()
			//fmt.Printf("%T",conn)
			if err != nil {
				fmt.Println("server accept err:", err)
				return
			}

			//读取包体
			go func(conn *net.Conn) {
				//读取客户端的请求
				//----拆包过程---
				// |datalen|id|data|
				dp := NewDataPack() //创建对象
				//循环读取
				for {
					//读出包头
					headData := make([]byte, dp.GetHeadLen())
					//直到headData填充满，才会返回，否则阻塞
					//fmt.Printf("%T",*conn)
					_, err := io.ReadFull(*conn, headData) //conn继承了reader接口
					if err != nil {
						fmt.Println("read head err:", err)
						break
					}
					//headData ==  > |datalen|id|  （8字节的长度）
					//将headData ---> Message结构体中 填充 len  id两个字段
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("UnPack err:", err)
						return
					}

					//判断数据长度是否大于0
					if msgHead.GetMsgLen() > 0 {
						//数据区含有数据,进行二次读取
						//将msgHead断言为Message类型
						msg := msgHead.(*Message)
						//给msg.Data开辟空间  长度就是数据的长度  data|
						msg.Data = make([]byte, msg.GetMsgLen())

						//根据datalen的长度进行第二次read
						_, err := io.ReadFull(*conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack err", err)
							return
						}
						fmt.Println("--->Recv MsgId=", msg.Id, "MsgDataLen=", msg.Datalen, "MsgData=", string(msg.Data))
					}
				}
			}(&conn)

		}
	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dail err:", err)
		return
	}

	//封包
	dp := NewDataPack()

	msg1 := &Message{
		Id:      1,
		Datalen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client send data1 err:", err)
		return
	}

	msg2 := &Message{
		Id:      2,
		Datalen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client send data2 err:", err)
		return
	}

	//将两个包连在一起
	sendData1 = append(sendData1, sendData2...) //打散

	//发送数据
	conn.Write(sendData1)

	select {}

}

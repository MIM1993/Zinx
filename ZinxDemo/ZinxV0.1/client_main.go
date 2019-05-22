package main

import (
	"net"
	"fmt"
	"time"
)

//模拟客户端

func main() {
	fmt.Println("client start ...")
	time.Sleep(1 * time.Second)
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client link err:", err)
		return
	}
	//起循环
	for {
		//写
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("write conn err :", err)
			return
		}

		//读
		buff := make([]byte, 512)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("read conn err :", err)
			return
		}

		fmt.Printf("server call back: %s,cnt=%d\n", buff, n)


		time.Sleep(1 * time.Second)
	}

}

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start .....")
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("client link err :", err)
		return
	}

	buff := make([]byte, 512)
	for {
		_, err := conn.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("conn write err :", err)
			return
		}

		cnt, err := conn.Read(buff)
		if err != nil {
			fmt.Println("conn read err:", err)
			return
		}

		fmt.Printf("server call back :%s  cnt:%d\n", buff[:cnt], cnt)

		time.Sleep(1 * time.Second)
	}
}

package dataTransfer

import (
	"fmt"
	"net"
	"socksFive"
)

func DataHandle(destination *socksFive.Destination, conn net.Conn)  {
	remote,err := net.Dial("tcp", destination.Addr+":"+destination.Port)
	//fmt.Println(destination.Addr,destination.Port)
	if err != nil {
		fmt.Println("connect destination error")
		return
	}
	go sendData(conn, remote)
	go sendData(remote, conn)
}

func sendData(conn net.Conn, remote net.Conn)  {
	defer conn.Close()
	defer remote.Close()
	buffer := make([]byte,256)

	for {
		//fmt.Println(666)
		//fmt.Println(buffer)
		num,err := conn.Read(buffer)
		fmt.Println(123)
		if err != nil {
			fmt.Println(err,1)
			return
		}
		fmt.Println(buffer,111)

		if num == 0 {
			return
		}
		_, err =remote.Write(buffer[:num])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
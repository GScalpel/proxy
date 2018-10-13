package dataTransfer

import (
	"fmt"
	"net"
)


func SendData(conn net.Conn, remote net.Conn)  {
	defer conn.Close()
	defer remote.Close()
	buffer := make([]byte,1400)

	for {
		num,err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err,1)
			return
		}

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
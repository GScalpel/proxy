package main

import (
	"dataTransfer"
	"flag"
	"fmt"
	"socksFive"
	"startTask"
)

func main()  {
	host := flag.String("host","127.0.0.1","need input host")
	port := flag.String("port","1080","need input port")
	sock := startTask.ListenAndServer(*host, *port)
	defer sock.Close()
	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Connection Error")
			continue
		}

		go func() {
			destination := socksFive.HandleSocks5(conn)
			if destination == nil {
				fmt.Println("HandleSocks5 happens error")
				return
			}
			dataTransfer.DataHandle(destination, conn)
		}()
	}
}

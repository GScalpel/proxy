package main

import (
	"flag"
	"linkHandle"
)

type basicSetting struct {
	opinion string
	host string
	port string
}


func main()  {
	sel := flag.String("select", "local", "select  to open server or local service(`server or local`)")
	host := flag.String("host","127.0.0.1","need input host")
	port := flag.String("port","1080","need input port")
	sHost := flag.String("sHost","8.8.8.8", "index host")
	sPort := flag.String("sPort", "8888", "index port")
	flag.Parse()

	base := basicSetting{*sel, *host, *port}


	sock := linkHandle.ListenAndServer(base.host, base.port)
	defer sock.Close()
	switch base.opinion {
	case "local":
		serverMsg := new(linkHandle.ServerMessage)
		serverMsg.ServerHost, serverMsg.ServerPort = *sHost, *sPort
		linkHandle.Local(sock, serverMsg)
	case "server":
		linkHandle.Server(sock)
	}
}


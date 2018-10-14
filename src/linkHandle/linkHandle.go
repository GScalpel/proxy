package linkHandle

import (
	"dataTransfer"
	"log"
	"net"
	"socksFive"
)

type ServerMessage struct {
	ServerHost string
	ServerPort string
}

func ListenAndServer(host string,port string) net.Listener {
	sock, err := net.Listen("tcp", host + ":" +port)
	if err != nil {
		log.Fatalf("index start error whether host is mistake or port has been used")
	}
	return sock
}

func Local(sock net.Listener, serverMsg *ServerMessage) {
	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
			continue
		}

		go func() {
			destination, err := socksFive.HandleSocks5(conn)
			if destination == nil || err != nil {
				conn.Close()
				log.Println("HandleSocks5 happens error: ", err)
				return
			}

			ConnHandleL(destination, serverMsg, conn)
		}()
	}
}

func Server(sock net.Listener) {
	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Println("Accept Error: ",err)
			continue
		}

		go func() {
			buf := make([]byte, 256)
			num,err := conn.Read(buf)
			if err != nil {
				log.Println("Destination Message Read error: ", err)
				conn.Close()
				return
			}
			dst := string(buf[:num])
			ConnHandleS(dst, conn)
		}()
	}
}

func ConnHandleL(destination *socksFive.Destination, serverMsg *ServerMessage, conn net.Conn)  {
	remote,err := net.Dial("tcp", serverMsg.ServerHost+":"+serverMsg.ServerPort)
	if err != nil {
		return
	}
	buf := []byte(destination.Addr + ":" + destination.Port)
	remote.Write(buf)
	go dataTransfer.SendData(conn, remote)
	go dataTransfer.SendData(remote, conn)
}

func ConnHandleS(dst string,  conn net.Conn)  {
	remote,err := net.Dial("tcp", dst)
	if err != nil {
		return
	}
	go dataTransfer.SendData(conn, remote)
	go dataTransfer.SendData(remote, conn)
}

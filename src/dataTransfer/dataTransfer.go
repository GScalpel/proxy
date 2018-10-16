package dataTransfer

import (
	"encoding/gob"
	"log"
	"net"
	"os"
)

type PWDBook map[string][256]byte

var Book = PWDBook{}


func init()  {
	file ,err := os.OpenFile("/home/scalpel/go/src/proxyServer/pwdBook.gob", os.O_RDONLY, 0666)
	if err != nil {
		log.Println("Loading pwdBook error")
		os.Exit(1)
	}
	dec := gob.NewDecoder(file)
	dec.Decode(&Book)
}

func SendDataEncrypt(conn net.Conn, remote net.Conn,status string)  {
	defer conn.Close()
	defer remote.Close()

	buffer := make([]byte,1024)
	newBuf := make([]byte,1024)

	//randNum := rand.Intn(3)+1
	//bookNum := status + strconv.Itoa(randNum)
	//newBuf[0] = byte(randNum)
	book := Book[status]

	for {
		num,err := conn.Read(buffer)
		if err != nil || num ==0 {
			return
		}

		for i := 0; i < num; i++ {
			newBuf[i] = book[buffer[i]]
		}

		_, err =remote.Write(newBuf[:num])
		if err != nil {
			return
		}
	}
}

func SendDataDecrypt(remote net.Conn, conn net.Conn, status string)  {
	defer remote.Close()
	defer conn.Close()

	buffer := make([]byte,1024)
	newBuf := make([]byte,1024)
	book := Book[status]

	for {
		num, err := remote.Read(buffer)
		if err != nil || num ==0 {
			return
		}

		//book := Book[status + strconv.Itoa(int(buffer[0]))]
		for i := 0; i < num; i++ {
			newBuf[i] = book[buffer[i]]
		}

		_, err = conn.Write(newBuf[:num])
		if err != nil {
			return
		}
	}

}
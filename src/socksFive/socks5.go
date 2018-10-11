package socksFive

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

const (
	ipv4 = 1
	ipv6 = 4
	domain = 3
	ipv4Len = 4
	ipv6Len = 16
	portLen = 2
)
type (
	Destination struct {
		Addr string
		Port string
	}
)

var	(
	repAuth = []byte{5, 0}
	repInfo = []byte{
		5, 0, 0, 1,
		0, 0, 0, 0,
		80, 80,
	}
)

// handle socks5 and analyze Destination message
func HandleSocks5(conn net.Conn) *Destination {
	buffer := make([]byte, 257)

	// Accept request for certification
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error")
		return nil
	}

	// Respond for certification
	_, err = conn.Write(repAuth)
	if err != nil {
		fmt.Println("Write error")
		return nil
	}

	// request message
	//bufType := make([]byte, 4)
	bufType := buffer[:4]
	_, err = conn.Read(bufType)
	if err != nil {
		fmt.Println("Read error")
		return nil
	}

	if bufType[1] != 1 {
		fmt.Println("commend error : need connect commend")
		return nil
	}

	destination := new(Destination)
	switch bufType[3] {
	case ipv4:
		{
			//bufHost := make([]byte, ipv4Len)
			bufHost := buffer[:ipv4Len]
			_, err = conn.Read(bufHost)
			if err != nil {
				fmt.Println("Read error")
				return nil
			}
			host := net.IP(bufHost).String()
			destination.Addr = host
		}

	case ipv6:
		{
			//bufHost := make([]byte, ipv6Len)
			bufHost := buffer[:ipv6Len]
			_, err = conn.Read(bufHost)
			if err != nil {
				fmt.Println("Read error")
				return nil
			}
			host := net.IP(bufHost).String()
			destination.Addr = host
		}

	case domain:
		{
			//hostLen := make([]byte, 1)
			hostLen := buffer[:1]
			_, err := conn.Read(hostLen)
			if err != nil {
				fmt.Println("Read error")
				return nil
			}

			//bufHost := make([]byte, num)
			bufHost := buffer[:hostLen[0]]
			_, err = conn.Read(bufHost)
			if err != nil {
				fmt.Println("Read error")
				return nil
			}

			host := string(bufHost)
			check := domainCheck(host)
			if check != true {
				fmt.Println("Not a domain")
				return nil
			}

			destination.Addr = host
		}
	}

	//bufPort := make([]byte, portLen)
	bufPort := buffer[:portLen]
	_, err = conn.Read(bufPort)
	if err != nil {
		fmt.Println("Read error port")
		return nil
	}

	port := strconv.Itoa(int(bufPort[0])*256 + int(bufPort[1]))
	destination.Port = port

	return destination
}

// check domain is legal
func domainCheck(host string) bool {
	check, err := regexp.MatchString(".[a-z]+$", host)
	fmt.Println(check, host)
	if err != nil {
		fmt.Println("regexp error")
		return false
	}

	return check
}
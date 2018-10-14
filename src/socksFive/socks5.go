package socksFive

import (
	"errors"
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
func HandleSocks5(conn net.Conn) (*Destination, error) {
	buffer := make([]byte, 257)

	// Accept request for certification
	_, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	// Respond for certification
	_, err = conn.Write(repAuth)
	if err != nil {
		return nil, err
	}

	// request message
	//bufType := make([]byte, 4)
	bufType := buffer[:4]
	_, err = conn.Read(bufType)
	if err != nil {
		return nil, err
	}

	if bufType[1] != 1 {
		err = errors.New("NOT CONNECT REQUEST")
		return nil, err
	}

	destination := new(Destination)
	switch bufType[3] {
	case ipv4:
		{
			//bufHost := make([]byte, ipv4Len)
			bufHost := buffer[:ipv4Len]
			_, err = conn.Read(bufHost)
			if err != nil {
				return nil, err
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
				return nil, err
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
				return nil, err
			}

			//bufHost := make([]byte, num)
			bufHost := buffer[:hostLen[0]]
			_, err = conn.Read(bufHost)
			if err != nil {
				return nil,err
			}

			host := string(bufHost)
			check := domainCheck(host)
			if check != true {
				return nil,err
			}

			destination.Addr = host
		}
	}

	//bufPort := make([]byte, portLen)
	bufPort := buffer[:portLen]
	_, err = conn.Read(bufPort)
	if err != nil {
		return nil,err
	}
	conn.Write(repInfo)
	port := strconv.Itoa(int(bufPort[0])*256 + int(bufPort[1]))
	destination.Port = port

	return destination, nil
}

// check domain is legal
func domainCheck(host string) bool {
	check, err := regexp.MatchString(".[a-z]+$", host)
	if err != nil {
		return false
	}

	return check
}
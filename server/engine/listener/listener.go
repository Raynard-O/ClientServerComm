package listener

import (
	"log"
	"net"
	"os"
)

type ListenInterface interface {
	GetListener() *net.TCPListener
}

type Listen struct {
	Network *net.TCPListener
}

func CreateListener(port int) *Listen {
	addr := &net.TCPAddr{
		IP:   net.IP{127, 0, 0, 1},
		Port: port,
	}
	list, err := net.ListenTCP(addr.Network(), addr)

	if err != nil {
		log.Printf("Error Listening on Address : %v", addr)
		os.Exit(1)
	}
	log.Printf("Listening on TCP Addr : %v", addr)
	return &Listen{Network: list}
}

func (l *Listen) GetListener() *net.TCPListener {
	return l.Network
}

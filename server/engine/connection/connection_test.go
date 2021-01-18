package connection

import (
	listener2 "ClientServerComm/server/engine/listener"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net"
	"testing"
)

type Conn struct {
	conn        *net.TCPConn
	connections *Connection
}

func TestNewConnection(t *testing.T) {

	l := CreateConnection(t)

	c, _ := l.AcceptTCP()
	conn := NewConnection(c)
	require.NotNil(t, conn)
}

func CreateConnection(t *testing.T) *net.TCPListener {
	l, _ := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IP{127, 0, 0, 1},
		Port: 9090,
	})
	defer l.Close()
	return l
}

func TestConnection_Read(t *testing.T) {

	listener := listener2.CreateListener(8080)
	l, err := listener.Network.AcceptTCP()
	defer l.Close()
	assert.NoError(t, err)

	for {
		conn := NewConnection(l)
		defer conn.Close()
		msg, err := conn.Read()
		log.Print(msg)
		assert.NoError(t, err)
		assert.NotNil(t, msg)
		return
	}
}

func TestConnection_Write(t *testing.T) {

	listener := listener2.CreateListener(8080)
	l, err := listener.Network.AcceptTCP()
	defer l.Close()
	assert.NoError(t, err)

	for {
		conn := NewConnection(l)
		defer conn.Close()
		err := conn.Write("Server Test\n")
		assert.NoError(t, err)
		return
	}
}
func TestConnection_Close(t *testing.T) {
	listener := listener2.CreateListener(8080)
	l, err := listener.Network.AcceptTCP()
	defer l.Close()
	assert.NoError(t, err)

	for {
		conn := NewConnection(l)
		assert.NoError(t, conn.Close())
		return
	}
}

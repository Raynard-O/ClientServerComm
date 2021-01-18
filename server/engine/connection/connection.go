package connection

import (
	"bufio"
	"net"
)

type InterfaceConnection interface {
	GetConn() *net.TCPConn
	Read() ([]byte, error)
	Write(msg string) error
	Close() error
}

type Connection struct {
	conn *net.TCPConn
	r    *bufio.ReadWriter
}

func NewConnection(c *net.TCPConn) *Connection {
	return &Connection{conn: c}
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.conn
}
func (c *Connection) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	_, err := c.conn.Read(buf)
	return buf, err
}

func (c *Connection) Write(msg string) error {
	_, err := c.conn.Write([]byte(msg))
	return err
}

func (c *Connection) Close() error {
	return c.conn.Close()
}

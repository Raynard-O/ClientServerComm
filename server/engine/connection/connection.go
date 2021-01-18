package connection

import (
	"bufio"
	"net"
)


type ConnectionInterface interface {
	Read() ([]byte, error )
	Write(msg string) error
	Close()
}

type Connection struct {
	conn *net.TCPConn
	r *bufio.ReadWriter
}

func NewConnection(c *net.TCPConn) *Connection {
	return &Connection{conn: c}
}


func (c *Connection) Read() ([]byte, error ){
	buf := make([]byte, 1024)
	_, err := c.conn.Read(buf)
	return buf,err
}

func (c *Connection) Write(msg string) error {
	_, err  := c.conn.Write([]byte(msg))
	return err
}

func (c *Connection) Close()  {
	c.conn.Close()
}
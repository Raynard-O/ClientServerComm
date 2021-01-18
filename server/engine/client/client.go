package client

import (
	"ClientServerComm/server/engine/connection"
	"net"
	"sync"
)

type InterfaceClient interface {
	CreateClient(conn connection.InterfaceConnection) map[uint64]*client
	//AllService() reflect.Value
	GetAllClients() *map[uint64]*client
}

type client struct {
	id   uint64
	conn *net.TCPConn
}

type Service struct {
	m       sync.Mutex
	clients map[uint64]*client
}

func Init() *Service {
	return &Service{
		clients: make(map[uint64]*client),
	}
}

func (s *Service) CreateClient(conn connection.InterfaceConnection) map[uint64]*client {
	c := conn.GetConn()
	s.m.Lock()
	defer s.m.Unlock()
	// get new ID
	id := uint64(1) //replace with a generate id algorithm
	// add client to clients
	cli := new(client)
	cli.conn = c
	cli.id = id
	s.clients[id] = cli
	m := make(map[uint64]*client)
	m[id] = cli
	return m
}

func (s *Service) GetAllClients() *map[uint64]*client {
	return &s.clients
}

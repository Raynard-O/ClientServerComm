package client

import (
	"ClientServerComm/server/engine/connection"
	listener2 "ClientServerComm/server/engine/listener"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

func TestInit(t *testing.T) {
	i := Init()

	assert.NotNil(t, i)
}

func TestService_CreateClient(t *testing.T) {
	i := Init()
	listener := listener2.CreateListener(8080)

	var wg sync.WaitGroup

	for {
		l, err := listener.Network.AcceptTCP()
		defer l.Close()
		assert.NoError(t, err)

		conn := connection.NewConnection(l)
		defer conn.Close()
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			ret := i.CreateClient(conn)
			assert.NotNil(t, ret)
			wg.Done()
		}(&wg)
		wg.Wait()
		return
	}
}

func TestService_GetAllClients(t *testing.T) {
	i := Init()
	listener := listener2.CreateListener(8080)
	l, err := listener.Network.AcceptTCP()
	defer l.Close()
	assert.NoError(t, err)
	for {
		conn := connection.NewConnection(l)

		i.CreateClient(conn)
		c := i.GetAllClients()
		msg := fmt.Sprintf("Client Created!!!\n Client Details : %v", c)
		log.Printf(msg)
		assert.NotNil(t, c)
		conn.Close()
		return
	}
}

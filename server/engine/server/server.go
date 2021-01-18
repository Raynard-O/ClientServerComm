package server

import (
	"ClientServerComm/server/engine/client"
	"ClientServerComm/server/engine/connection"
	"ClientServerComm/server/engine/listener"
	"ClientServerComm/server/message"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

func Server(l listener.ListenInterface, cli client.InterfaceClient) {
	listen := l.GetListener()
	var wg sync.WaitGroup
	defer listen.Close()
	//go func() {

		for {
			c, err := listen.AcceptTCP()

			if err != nil {
				log.Printf("Error Connecting on Address : %v", listen.Addr())
				return
			}
			log.Printf("New Connection on TCP Addr : %v", listen.Addr())
			conn := connection.NewConnection(c)
			wg.Add(1)

			cli.CreateClient(conn)


			//go Serve(conn, cli)
			// get user client struct
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				Serve(conn, cli)

			}(&wg)
		}
	//}()

}

func Serve(c connection.InterfaceConnection, cli client.InterfaceClient) {
	defer c.Close()
	fmt.Println(cli.GetAllClients())
	for {
		r, err := c.Read()
		if err != nil {
			log.Printf("Error Reading From Connection")
			return
		}
		log.Printf("Incoming Message")
		msg := strings.TrimSpace(string(r))
		fmt.Printf("Message From Client: %v", msg)

		comp := strings.Compare(msg, "stop")
		fmt.Println(comp)
		if comp == 0 {
			fmt.Printf("Closing Communication with Client")
			break
		}
		switch msg {
		case message.Identify:
			// send user id to client
		}
		t := time.Now().String()
		log.Printf("Sending Message to Client")
		err = c.Write(t)
		if err != nil {
			log.Printf("Error Writing to Connection")
			return
		}
	}

}

package server

import (
	"ClientServerComm/server/engine/connection"
	"ClientServerComm/server/engine/listener"
	"fmt"
	"log"
	"strings"
	"time"
)

func Server(l listener.ListenInterface)  {
	listen := l.GetListener()
	defer listen.Close()
	for {
		c , err := listen.AcceptTCP()

		if err != nil {
			log.Printf("Error Connecting on Address : %v", listen.Addr())
			return
		}
		log.Printf("New Connection on TCP Addr : %v", listen.Addr())
		conn := connection.NewConnection(c)

		go Serve(conn)

	}
}

func Serve(c connection.ConnectionInterface)  {
	defer c.Close()
	for {
		message, err := c.Read()
		if err != nil {
			log.Printf("Error Reading From Connection")
			return
		}
		log.Printf("Incoming Message")
		msg := strings.TrimSpace(string(message))
		fmt.Printf("Message From Client: %v", msg)
		fmt.Println(msg)
		comp := strings.Compare(msg, "STOP")
		if comp == 1 {
			fmt.Printf("Closing Communication with Client")
			break
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